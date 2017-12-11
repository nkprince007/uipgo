package lib

import (
	"fmt"
	"strings"
)

// Image is a golang representation for image metadata.
type Image interface {
	URL() string
	Name() string
}

// UnsplashImage represents an image from https://unsplash.com
type UnsplashImage struct {
	URLs map[string]string `json:"urls"`
}

// URL retrieves the url of UnsplashImage
func (i UnsplashImage) URL() string {
	return i.URLs["regular"]
}

// Name retrieves the name of UnsplashImage
func (i UnsplashImage) Name() string {
	parts := strings.Split(i.URL(), "/")
	index := strings.Index(parts[len(parts)-1], "?")
	identifier := parts[len(parts)-1][:index]
	name := fmt.Sprintf("unsplash-%s.jpg", identifier)
	return name
}

// DesktopprResponseURL contains the actual URL field for DesktopprImage
type DesktopprResponseURL struct {
	URL string `json:"url"`
}

// DesktopprImage represents an image on https://desktoppr.co
type DesktopprImage struct {
	Image DesktopprResponseURL `json:"image"`
}

// DesktopprAPIResponse represents a list of DesktopprImages from the API
// response on https://api.desktoppr.co
type DesktopprAPIResponse struct {
	Images []DesktopprImage `json:"response"`
}

// URL retrieves the url of DesktopprImage
func (i DesktopprImage) URL() string {
	return i.Image.URL
}

// Name retrieves the name of DesktopprImage
func (i DesktopprImage) Name() string {
	parts := strings.Split(i.URL(), "/")
	name := parts[len(parts)-1]
	name = fmt.Sprintf("desktoppr-%s", name)
	return name
}
