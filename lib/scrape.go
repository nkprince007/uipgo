package lib

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"sync"

	"github.com/urfave/cli"
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

// GetAndStoreImages downloads and stores images from given websites.
func GetAndStoreImages(sites map[string][]string, c *cli.Context) {
	images := []Image{}
	var wg sync.WaitGroup

	list, ok := sites["unsplash"]
	if ok {
		for _, site := range list {
			wg.Add(1)
			images = append(images, GetUnsplashImages(site, &wg)...)
		}
	}

	list, ok = sites["desktoppr"]
	if ok {
		for _, site := range list {
			wg.Add(1)
			images = append(images, GetDesktopprImages(site, &wg)...)
		}
	}
	wg.Wait()

	for _, image := range images {
		wg.Add(1)
		go DownloadFile(c.String("directory"), image.Name(), image.URL(), &wg)
	}
	wg.Wait()

	return
}

// DownloadFile downloads a file from the given url and stores it in filepath
func DownloadFile(
	dir string, filename string, rawurl string, wg *sync.WaitGroup) {

	if wg != nil {
		defer wg.Done()
	}

	_, err := url.ParseRequestURI(rawurl)
	Check(err)

	err = os.MkdirAll(dir, os.ModePerm)
	Check(err)

	out, err := os.Create(filename)
	Check(err)

	defer os.Rename(filename, filepath.Join(dir, filename))
	defer out.Close()

	resp, err := http.Get(rawurl)
	Check(err)

	defer resp.Body.Close()

	_, err = io.Copy(out, resp.Body)
	Check(err)

	fmt.Println("Image downloaded successfully: " + filename)
	return
}
