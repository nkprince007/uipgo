package utils

import (
	"os/exec"
	"path/filepath"
	"strings"
)

// GetWallpaper returns the path to the current wallpaper.
func GetWallpaper() (string, error) {
	// notest

	stdout, err := exec.Command(
		"gsettings",
		"get",
		"org.gnome.desktop.background",
		"picture-uri",
	).Output()

	if err != nil {
		return "", err
	}

	return strings.TrimSpace(string(stdout)), nil
}

// SetWallpaper sets the wallpaper to the picture from provided path.
func SetWallpaper(path string) error {
	// notest

	modPath := ""

	if strings.HasPrefix(path, "file:/") {
		absPath, err := filepath.Abs(path)
		if err != nil {
			return err
		}
		modPath = "file:///" + absPath
	} else {
		modPath = path
	}

	return exec.Command(
		"gsettings",
		"set",
		"org.gnome.desktop.background",
		"picture-uri",
		modPath,
	).Run()
}
