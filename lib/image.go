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
