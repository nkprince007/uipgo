package utils

import (
	"errors"
	"runtime"
)

// GetWallpaper retrieves the path to the current wallpaper.
func GetWallpaper() (string, error) {
	// notest

	if runtime.GOOS == "darwin" {
		return GetMacOSWallpaper()
	}

	return "", errors.New("GetWallpaper isn't implemented for your platform yet")
}

// SetWallpaper sets the wallpaper to the picture from provided path.
func SetWallpaper(path string) error {
	//notest

	if runtime.GOOS == "darwin" {
		return SetMacOSWallpaper(path)
	}

	return errors.New("SetWallpaper isn't implemented for your platform yet")
}
