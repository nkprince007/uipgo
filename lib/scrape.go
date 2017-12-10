package lib

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
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
		log.Println(e)
	}
}

// GetUnsplashImages gets a list of images from provided Unsplash API endpoint.
func GetUnsplashImages(rawurl string) []Image {
	_, err := url.ParseRequestURI(rawurl)
	Check(err)

	client := &http.Client{}
	req, err := http.NewRequest("GET", rawurl, nil)
	Check(err)

	req.Header.Add("User-Agent", "uipgo")
	resp, err := client.Do(req)
	Check(err)

	defer resp.Body.Close()

	ret := [NOOFIMAGES]UnsplashImage{}
	err = json.NewDecoder(resp.Body).Decode(&ret)
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

	list, ok := sites["unsplash"]
	if ok {
		for _, site := range list {
			images = append(images, GetUnsplashImages(site)...)
		}
	}

	var wg sync.WaitGroup
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
