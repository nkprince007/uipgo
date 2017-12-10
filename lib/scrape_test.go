package lib

import (
	"net/url"
	"os"
	"testing"
)

const unsplashClientID = "74f6347705c15665e0d3d4b241fce1e9c2ef26761aeddfe0724dcd00d2823af5"

func TestGetUnsplashImages(t *testing.T) {
	endpoint := "https://api.unsplash.com/photos?client_id=" + unsplashClientID
	images := GetUnsplashImages(endpoint)

	// testing length of no. of images retrieved
	if len(images) != NOOFIMAGES {
		t.Errorf(
			"Returns more images than expected. (%d > %d)",
			len(images), NOOFIMAGES)
	}

	// testing validity of URLs retrieved
	for _, img := range images {
		_, err := url.ParseRequestURI(img.URL())
		if err != nil {
			t.Errorf("%s is not a valid URL.", img.URL())
		}
	}
}

func TestDownloadFile(t *testing.T) {
	const TESTFILE = "test.json"
	defer os.Remove(TESTFILE)

	DownloadFile(".", TESTFILE, "https://github.com/nkprince007/uipgo.json", nil)

	info, err := os.Stat(TESTFILE)
	if info.Size() <= 0 || err != nil {
		t.Error("DownloadFile stored an empty file.")
	}
}
