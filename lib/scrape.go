package lib

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
	"sync"
	"time"

	"github.com/gosuri/uiprogress"
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

// DownloadFile downloads a file from the given url and stores it in filepath
func DownloadFile(
	path string, rawurl string, bar *uiprogress.Bar, wg *sync.WaitGroup,
) {
	if wg != nil {
		defer wg.Done()
	}

	_, err := url.ParseRequestURI(rawurl)
	Check(err)

	err = os.MkdirAll(filepath.Dir(path), os.ModePerm)
	Check(err)

	out, err := os.Create(path)
	Check(err)

	defer out.Close()

	headResp, err := http.Head(rawurl)
	Check(err)

	defer headResp.Body.Close()

	size, err := strconv.Atoi(headResp.Header.Get("Content-Length"))
	Check(err)

	done := make(chan int64)
	go ShowDownloadProgress(done, bar, path, int64(size))

	resp, err := http.Get(rawurl)
	Check(err)

	defer resp.Body.Close()

	written, err := io.Copy(out, resp.Body)
	Check(err)

	done <- written

	// wait for refreshing bar to full before exiting
	bar.Set(40)
	time.Sleep(10 * time.Millisecond)

	return
}
