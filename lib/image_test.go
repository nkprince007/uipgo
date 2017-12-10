package lib

import (
	"testing"
)

func TestUnsplashImage(t *testing.T) {
	const TESTURL = "https://images.unsplash.com/photo-1512869915288-bc1f58324319?ixlib=rb-0.3.5&q=80"

	image := UnsplashImage{
		URLs: map[string]string{
			"regular": TESTURL,
		},
	}

	// testing for URL()
	if image.URL() != TESTURL {
		t.Errorf("UnsplashImage.URL() returns an unintended URL. (%s != %s)",
			image.URL(), TESTURL)
	}

	// testing for Name()
	if image.Name() != "unsplash-photo-1512869915288-bc1f58324319.jpg" {
		t.Errorf("UnsplashImage.Name() returns unintended results. (%s != %s)",
			image.Name(), TESTURL)
	}
}
