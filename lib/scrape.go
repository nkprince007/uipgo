package lib

import (
	"encoding/json"
	"net/http"
	"net/url"
	"sync"
)

// NOOFIMAGES denotes the number of images to be downloaded per Website URL.
const NOOFIMAGES = 5

// Check logs any errors occured when an error is passed.
func Check(e error) {
	if e != nil {
		panic(e)
	}
}

func sendRequest(rawurl string) *http.Response {
	_, err := url.ParseRequestURI(rawurl)
	Check(err)

	client := &http.Client{}
	req, err := http.NewRequest("GET", rawurl, nil)
	Check(err)

	req.Header.Add("User-Agent", "uipgo")
	resp, err := client.Do(req)
	Check(err)

	return resp
}

// GetDesktopprImages gets a list of images from provided Desktoppr API
// endpoint.
func GetDesktopprImages(rawurl string, wg *sync.WaitGroup) []Image {
	if wg != nil {
		defer wg.Done()
	}

	resp := sendRequest(rawurl)
	defer resp.Body.Close()

	ret := DesktopprAPIResponse{}
	err := json.NewDecoder(resp.Body).Decode(&ret)
	Check(err)

	// type conversion for abiding to interface
	images := ret.Images
	retImage := make([]Image, len(images))
	for i := range images {
		retImage[i] = images[i]
	}

	return retImage[:NOOFIMAGES]
}

// GetUnsplashImages gets a list of images from provided Unsplash API endpoint.
func GetUnsplashImages(rawurl string, wg *sync.WaitGroup) []Image {
	if wg != nil {
		defer wg.Done()
	}

	resp := sendRequest(rawurl)
	defer resp.Body.Close()

	ret := [NOOFIMAGES]UnsplashImage{}
	err := json.NewDecoder(resp.Body).Decode(&ret)
	Check(err)

	// type conversion for abiding to interface
	retImage := make([]Image, len(ret))
	for i := range ret {
		retImage[i] = ret[i]
	}

	return retImage
}
