package lib

import (
	"errors"
	"net/url"
	"os"
	"sync"
	"testing"
)

const unsplashClientID = "74f6347705c15665e0d3d4b241fce1e9c2ef26761aeddfe0724dcd00d2823af5"

func TestGetUnsplashImages(t *testing.T) {
	endpoint := "https://api.unsplash.com/photos?client_id=" + unsplashClientID
	var wg sync.WaitGroup
	wg.Add(1)
	images := GetUnsplashImages(endpoint, &wg)
	wg.Wait()

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

func TestGetDesktopprImages(t *testing.T) {
	endpoint := "https://api.desktoppr.co/1/wallpapers"

	var wg sync.WaitGroup
	wg.Add(1)
	images := GetDesktopprImages(endpoint, &wg)
	wg.Wait()

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

func TestCheck(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The Check function did not panic.")
		}
	}()

	// the code to be tested
	Check(errors.New("darthvader has arrived"))
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
