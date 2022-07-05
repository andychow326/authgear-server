package resource

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/authgear/authgear-server/pkg/util/readcloserthunk"
)

type Location struct {
	Fs   Fs
	Path string
}

// nolint: golint
type ResourceFile struct {
	Location        Location
	ReadCloserThunk readcloserthunk.ReadCloserThunk
}

type Match struct {
	LanguageTag string
}

type Descriptor interface {
	MatchResource(path string) (*Match, bool)
	FindResources(fs Fs) ([]Location, error)
	ViewResources(resources []ResourceFile, view View) (interface{}, error)
	UpdateResource(ctx context.Context, resourcesInAllFss []ResourceFile, resourceInTargetFs *ResourceFile, data []byte) (*ResourceFile, error)
}

// SimpleDescriptor does not support view.
type SimpleDescriptor struct {
	Path string
}

var _ Descriptor = SimpleDescriptor{}

func (d SimpleDescriptor) MatchResource(path string) (*Match, bool) {
	if path == d.Path {
		return &Match{}, true
	}
	return nil, false
}

func (d SimpleDescriptor) FindResources(fs Fs) ([]Location, error) {
	location := Location{
		Fs:   fs,
		Path: d.Path,
	}
	_, err := StatLocation(location)
	if os.IsNotExist(err) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return []Location{location}, nil
}

func (d SimpleDescriptor) ViewResources(resources []ResourceFile, rawView View) (interface{}, error) {
	switch rawView.(type) {
	case AppFileView:
		var appResources []ResourceFile
		for _, resrc := range resources {
			if resrc.Location.Fs.GetFsLevel() == FsLevelApp {
				s := resrc
				appResources = append(appResources, s)
			}
		}
		return d.viewResources(appResources)
	case EffectiveFileView:
		return d.viewResources(resources)
	case EffectiveResourceView:
		return d.viewResources(resources)
	case ValidateResourceView:
		return d.viewResources(resources)
	default:
		return nil, fmt.Errorf("unsupported view: %T", rawView)
	}
}

func (d SimpleDescriptor) viewResources(resources []ResourceFile) (interface{}, error) {
	if len(resources) == 0 {
		return nil, ErrResourceNotFound
	}
	last := resources[len(resources)-1]
	return readcloserthunk.Performance_Bytes(last.ReadCloserThunk)
}

func (d SimpleDescriptor) UpdateResource(_ context.Context, _ []ResourceFile, resource *ResourceFile, data []byte) (*ResourceFile, error) {
	return &ResourceFile{
		Location:        resource.Location,
		ReadCloserThunk: readcloserthunk.Reader(bytes.NewReader(data)),
	}, nil
}

type NewlineJoinedDescriptor struct {
	Path  string
	Parse func([]byte) (interface{}, error)
}

var _ Descriptor = NewlineJoinedDescriptor{}

func (d NewlineJoinedDescriptor) MatchResource(path string) (*Match, bool) {
	if path == d.Path {
		return &Match{}, true
	}
	return nil, false
}

func (d NewlineJoinedDescriptor) FindResources(fs Fs) ([]Location, error) {
	location := Location{
		Fs:   fs,
		Path: d.Path,
	}
	_, err := StatLocation(location)
	if os.IsNotExist(err) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return []Location{location}, nil
}

func (d NewlineJoinedDescriptor) ViewResources(resources []ResourceFile, rawView View) (interface{}, error) {
	switch rawView.(type) {
	case AppFileView:
		var appResources []ResourceFile
		for _, resrc := range resources {
			if resrc.Location.Fs.GetFsLevel() == FsLevelApp {
				s := resrc
				appResources = append(appResources, s)
			}
		}
		return d.viewResources(appResources)
	case EffectiveFileView:
		return d.viewResources(resources)
	case EffectiveResourceView:
		bytes, err := d.viewResources(resources)
		if err != nil {
			return nil, err
		}
		if d.Parse == nil {
			return bytes, nil
		}
		return d.Parse(bytes)
	case ValidateResourceView:
		bytes, err := d.viewResources(resources)
		if err != nil {
			return nil, err
		}
		if d.Parse == nil {
			return bytes, nil
		}
		return d.Parse(bytes)
	default:
		return nil, fmt.Errorf("unsupported view: %T", rawView)
	}
}

func (d NewlineJoinedDescriptor) viewResources(resources []ResourceFile) ([]byte, error) {
	if len(resources) == 0 {
		return nil, ErrResourceNotFound
	}

	var thunks []readcloserthunk.ReadCloserThunk
	for _, resrc := range resources {
		rct := resrc.ReadCloserThunk
		thunks = append(thunks, rct, readcloserthunk.Reader(strings.NewReader("\n")))
	}

	output := readcloserthunk.MultiReadCloserThunk(thunks...)
	return readcloserthunk.Performance_Bytes(output)
}

func (d NewlineJoinedDescriptor) UpdateResource(_ context.Context, _ []ResourceFile, resource *ResourceFile, data []byte) (*ResourceFile, error) {
	return &ResourceFile{
		Location:        resource.Location,
		ReadCloserThunk: readcloserthunk.Reader(bytes.NewReader(data)),
	}, nil
}
