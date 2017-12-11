package lib

import (
	"testing"
)

func TestUnsplashImage(t *testing.T) {
	const TESTURL = "https://images.unsplash.com/photo-1512869915288-bc1f58324319?ixlib=rb-0.3.5&q=80"
	const TESTIMGNAME = "unsplash-photo-1512869915288-bc1f58324319.jpg"

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
	if image.Name() != TESTIMGNAME {
		t.Errorf("UnsplashImage.Name() returns unintended results. (%s != %s)",
			image.Name(), TESTIMGNAME)
	}
}

func TestDesktopprImage(t *testing.T) {
	const TESTURL = "http://a.desktopprassets.com/wallpapers/b384dd91199fcac4d1a2f3b73968c35f72d3ce73/utxrqmp.jpg"
	const TESTIMGNAME = "desktoppr-utxrqmp.jpg"

	image := DesktopprImage{
		Image: DesktopprResponseURL{TESTURL},
	}

	// testing for URL()
	if image.URL() != TESTURL {
		t.Errorf("DesktopprImage.URL() returns an unintended URL. (%s != %s)",
			image.URL(), TESTURL)
	}

	// testing for Name()
	if image.Name() != TESTIMGNAME {
		t.Errorf("DesktopprImage.Name() returns unintended results. (%s != %s)",
			image.Name(), TESTIMGNAME)
	}
}
