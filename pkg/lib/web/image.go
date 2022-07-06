package web

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"mime"
	"os"
	"path"
	"regexp"
	"sort"

	"github.com/authgear/authgear-server/pkg/api/apierrors"
	"github.com/authgear/authgear-server/pkg/util/intlresource"
	"github.com/authgear/authgear-server/pkg/util/libmagic"
	"github.com/authgear/authgear-server/pkg/util/readcloserthunk"
	"github.com/authgear/authgear-server/pkg/util/resource"
)

var imageResolveFsLevelPriority = []resource.FsLevel{
	resource.FsLevelApp, resource.FsLevelCustom, resource.FsLevelBuiltin,
}

type languageImage struct {
	LanguageTag     string
	RealLanguageTag string
	ReadCloserThunk readcloserthunk.ReadCloserThunk
}

func (i languageImage) GetLanguageTag() string {
	return i.LanguageTag
}

var preferredExtensions = map[string]string{
	"image/png":  ".png",
	"image/jpeg": ".jpeg",
	"image/gif":  ".gif",
}

var imageRegex = regexp.MustCompile(`^static/([a-zA-Z0-9-]+)/(.+)\.(png|jpe|jpeg|jpg|gif)$`)

type ImageDescriptor struct {
	Name string
}

var _ resource.Descriptor = ImageDescriptor{}

func (a ImageDescriptor) MatchResource(path string) (*resource.Match, bool) {
	matches := imageRegex.FindStringSubmatch(path)
	if len(matches) != 4 {
		return nil, false
	}
	languageTag := matches[1]
	name := matches[2]

	if name != a.Name {
		return nil, false
	}
	return &resource.Match{LanguageTag: languageTag}, true
}

func (a ImageDescriptor) FindResources(fs resource.Fs) ([]resource.Location, error) {
	staticDir, err := fs.Open("static")
	if os.IsNotExist(err) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	defer staticDir.Close()

	langTagDirs, err := staticDir.Readdirnames(0)
	if err != nil {
		return nil, err
	}

	var locations []resource.Location
	for _, langTag := range langTagDirs {
		stat, err := fs.Stat(path.Join("static", langTag))
		if err != nil {
			return nil, err
		}
		if !stat.IsDir() {
			continue
		}

		for mediaType := range preferredExtensions {
			exts, _ := mime.ExtensionsByType(mediaType)
			for _, ext := range exts {
				p := path.Join("static", langTag, a.Name+ext)
				location := resource.Location{
					Fs:   fs,
					Path: p,
				}
				_, err := resource.StatLocation(location)
				if os.IsNotExist(err) {
					continue
				} else if err != nil {
					return nil, err
				}
				locations = append(locations, location)
			}
		}
	}

	return locations, nil
}

func (a ImageDescriptor) ViewResources(resources []resource.ResourceFile, rawView resource.View) (interface{}, error) {
	switch view := rawView.(type) {
	case resource.AppFileView:
		return a.viewAppFile(resources, view)
	case resource.EffectiveFileView:
		return a.viewEffectiveFile(resources, view)
	case resource.EffectiveResourceView:
		return a.viewEffectiveResource(resources, view)
	case resource.ValidateResourceView:
		return a.viewValidateResource(resources, view)
	default:
		return nil, fmt.Errorf("unsupported view: %T", rawView)
	}
}

func (a ImageDescriptor) UpdateResource(_ context.Context, _ []resource.ResourceFile, resrc *resource.ResourceFile, data []byte) (*resource.ResourceFile, error) {
	if len(data) > 0 {
		typ := libmagic.MimeFromBytes(data)
		_, ok := preferredExtensions[typ]
		if !ok {
			return nil, UnsupportedImageFile.NewWithDetails("unsupported image file", apierrors.Details{
				"type": apierrors.APIErrorDetail.Value(typ),
			})
		}
	}

	return &resource.ResourceFile{
		Location:        resrc.Location,
		ReadCloserThunk: readcloserthunk.Reader(bytes.NewReader(data)),
	}, nil
}

func (a ImageDescriptor) viewValidateResource(resources []resource.ResourceFile, view resource.ValidateResourceView) (interface{}, error) {
	// Ensure there is at most one resource
	// For each Fs and for each locale, remember how many paths we have seen.
	seen := make(map[resource.Fs]map[string][]string)
	for _, resrc := range resources {
		languageTag := imageRegex.FindStringSubmatch(resrc.Location.Path)[1]
		m, ok := seen[resrc.Location.Fs]
		if !ok {
			m = make(map[string][]string)
			seen[resrc.Location.Fs] = m
		}
		paths := m[languageTag]
		paths = append(paths, resrc.Location.Path)
		m[languageTag] = paths
	}
	for _, m := range seen {
		for _, paths := range m {
			if len(paths) > 1 {
				sort.Strings(paths)
				return nil, fmt.Errorf("duplicate resource: %v", paths)
			}
		}
	}

	return nil, nil
}

func (a ImageDescriptor) viewEffectiveResource(resources []resource.ResourceFile, view resource.EffectiveResourceView) (interface{}, error) {
	preferredLanguageTags := view.PreferredLanguageTags()
	defaultLanguageTag := view.DefaultLanguageTag()

	var fallbackImage *languageImage
	images := make(map[resource.FsLevel]map[string]intlresource.LanguageItem)
	extractLanguageTag := func(location resource.Location) string {
		langTag := imageRegex.FindStringSubmatch(location.Path)[1]
		return langTag
	}
	add := func(langTag string, resrc resource.ResourceFile) error {
		fsLevel := resrc.Location.Fs.GetFsLevel()
		i := languageImage{
			LanguageTag:     langTag,
			RealLanguageTag: extractLanguageTag(resrc.Location),
			ReadCloserThunk: resrc.ReadCloserThunk,
		}
		if images[fsLevel] == nil {
			images[fsLevel] = make(map[string]intlresource.LanguageItem)
		}
		images[fsLevel][langTag] = i
		if fallbackImage == nil {
			fallbackImage = &i
		}
		return nil
	}

	err := intlresource.Prepare(resources, view, extractLanguageTag, add)
	if err != nil {
		return nil, err
	}

	var matched intlresource.LanguageItem
	for _, fsLevel := range imageResolveFsLevelPriority {
		var items []intlresource.LanguageItem
		imagesInFsLevel, ok := images[fsLevel]
		if !ok {
			continue
		}

		for _, i := range imagesInFsLevel {
			items = append(items, i)
		}

		matched, err = intlresource.Match(preferredLanguageTags, defaultLanguageTag, items)
		if err == nil {
			break
		} else if errors.Is(err, intlresource.ErrNoLanguageMatch) {
			continue
		} else {
			return nil, err
		}
	}

	if matched == nil {
		if fallbackImage != nil {
			// Use first item in case of no match, to ensure resolution always succeed
			matched = fallbackImage
		} else {
			// If no configured translation, fail the resolution process
			return nil, resource.ErrResourceNotFound
		}
	}

	tagger := matched.(languageImage)

	mimeType := readcloserthunk.HTTPDetectContentType(tagger.ReadCloserThunk)
	ext, ok := preferredExtensions[mimeType]
	if !ok {
		return nil, fmt.Errorf("invalid image format: %s", mimeType)
	}

	path := fmt.Sprintf("%s%s/%s%s", StaticAssetResourcePrefix, tagger.RealLanguageTag, a.Name, ext)
	return &StaticAsset{
		Path:            path,
		ReadCloserThunk: tagger.ReadCloserThunk,
	}, nil
}

func (a ImageDescriptor) viewAppFile(resources []resource.ResourceFile, view resource.AppFileView) (interface{}, error) {
	path := view.AppFilePath()
	var appResources []resource.ResourceFile
	for _, resrc := range resources {
		if resrc.Location.Fs.GetFsLevel() == resource.FsLevelApp {
			appResources = append(appResources, resrc)
		}
	}
	asset, err := a.viewByPath(appResources, path)
	if err != nil {
		return nil, err
	}
	return readcloserthunk.Performance_Bytes(asset.ReadCloserThunk)
}

func (a ImageDescriptor) viewEffectiveFile(resources []resource.ResourceFile, view resource.EffectiveFileView) (interface{}, error) {
	path := view.EffectiveFilePath()
	asset, err := a.viewByPath(resources, path)
	if err != nil {
		return nil, err
	}
	return readcloserthunk.Performance_Bytes(asset.ReadCloserThunk)
}

func (a ImageDescriptor) viewByPath(resources []resource.ResourceFile, path string) (*StaticAsset, error) {
	matches := imageRegex.FindStringSubmatch(path)
	if len(matches) < 4 {
		return nil, resource.ErrResourceNotFound
	}
	requestedLangTag := matches[1]
	requestedExtension := matches[3]

	var found bool
	var rct readcloserthunk.ReadCloserThunk
	for _, resrc := range resources {
		m := imageRegex.FindStringSubmatch(resrc.Location.Path)
		langTag := m[1]
		extension := m[3]
		if langTag == requestedLangTag && extension == requestedExtension {
			found = true
			rct = resrc.ReadCloserThunk
		}
	}

	if !found {
		return nil, resource.ErrResourceNotFound
	}

	mimeType := readcloserthunk.HTTPDetectContentType(rct)
	ext, ok := preferredExtensions[mimeType]
	if !ok {
		return nil, fmt.Errorf("invalid image format: %s", mimeType)
	}

	p := fmt.Sprintf("%s%s/%s%s", StaticAssetResourcePrefix, requestedLangTag, a.Name, ext)
	return &StaticAsset{
		Path:            p,
		ReadCloserThunk: rct,
	}, nil
}
