package web

import (
	"github.com/authgear/authgear-server/pkg/util/readcloserthunk"
	"github.com/authgear/authgear-server/pkg/util/resource"
)

const StaticAssetResourcePrefix = "static/"
const GeneratedStaticAssetResourcePrefix = "static/generated/"
const StaticAssetFontResourcePrefix = "static/fonts/"

type StaticAsset struct {
	Path            string
	ReadCloserThunk readcloserthunk.ReadCloserThunk
}

var GeneratedAsset = resource.RegisterResource(NewGeneratedAssetDescriptor())

var AuthgearLightThemeCSS = resource.RegisterResource(CSSDescriptor{
	Path: StaticAssetResourcePrefix + "authgear-light-theme.css",
})

var AuthgearDarkThemeCSS = resource.RegisterResource(CSSDescriptor{
	Path: StaticAssetResourcePrefix + "authgear-dark-theme.css",
})

var AuthgearCSS = resource.RegisterResource(CSSDescriptor{
	Path: StaticAssetResourcePrefix + "authgear.css",
})

var AppLogo = resource.RegisterResource(ImageDescriptor{Name: "app_logo"})
var AppLogoDark = resource.RegisterResource(ImageDescriptor{Name: "app_logo_dark"})
var Favicon = resource.RegisterResource(ImageDescriptor{Name: "favicon"})
var AvatarPlaceholder = resource.RegisterResource(ImageDescriptor{Name: "avatar_placeholder"})
