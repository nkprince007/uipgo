package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"

	"github.com/urfave/cli"
)

// Check raises panic when an error is passed.
func Check(e error) {
	if e != nil {
		panic(e)
	}
}

func getUnsplashImages(rawurl string) []Image {
	retImage := []Image{}

	_, err := url.ParseRequestURI(rawurl)
	Check(err)

	client := &http.Client{}
	req, err := http.NewRequest("GET", rawurl, nil)
	Check(err)

	req.Header.Add("User-Agent", "uipgo")
	resp, err := client.Do(req)
	Check(err)

	defer resp.Body.Close()

	ret := []UnsplashImage{}
	err = json.NewDecoder(resp.Body).Decode(&ret)
	Check(err)

	// type conversion for abiding to interface
	for i := range ret {
		retImage = append(retImage, Image(ret[i]))
	}

	return retImage
}

// GetAndStoreImages downloads and stores images from given websites.
func GetAndStoreImages(sites map[string][]string, c *cli.Context) {
	images := []Image{}

	list, ok := sites["unsplash"]
	if ok {
		for _, site := range list {
			images = append(images, getUnsplashImages(site)...)
		}
	}

	for _, image := range images {
		err := DownloadFile(c.String("directory"), image.Name(), image.URL())
		if err != nil {
			log.Println(err)
			continue
		}
		fmt.Println("Downloaded image: " + image.Name())
	}

	return
}

// DownloadFile downloads a file from the given url and stores it in filepath
func DownloadFile(dir string, filename string, rawurl string) error {
	_, err := url.ParseRequestURI(rawurl)
	if err != nil {
		return err
	}

	err = os.MkdirAll(dir, os.ModePerm)
	Check(err)

	out, err := os.Create(filename)
	defer os.Rename(filename, filepath.Join(dir, filename))

	if err != nil {
		return err
	}
	defer out.Close()

	resp, err := http.Get(rawurl)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}
	return nil
}
