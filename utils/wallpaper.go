// +build !darwin,!linux

package utils

import (
	"errors"
)

// GetWallpaper is a dummy function to help build successfully on all platforms
func GetWallpaper() (string, error) {
	// notest
	return "", errors.New("GetWallpaper is not implemented for your platform")
}

// SetWallpaper is a dummy function to help build successfully on all platforms
func SetWallpaper(path string) error {
	// notest
	return errors.New("SetWallpaper is not implemented for your platform")
}
