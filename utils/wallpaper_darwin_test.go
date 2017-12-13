// +build darwin

package utils

import "testing"

func TestMacOSWallpaper(t *testing.T) {
	wp, err := GetMacOSWallpaper()
	if err != nil {
		t.Error(err)
	}

	err = SetMacOSWallpaper(wp)
	if err != nil {
		t.Error(err)
	}

	wpTest, err := GetMacOSWallpaper()
	if err != nil {
		t.Error(err)
	}

	if wpTest != wp {
		t.Errorf("SetMacOSWallpaper failed!")
	}
}
