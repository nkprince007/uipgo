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
	"strings"

	"github.com/urfave/cli"
)

// Image is a golang representation for image metadata.
type Image struct {
	url  string
	name string
}

// Check raises panic when an error is passed.
func Check(e error) {
	if e != nil {
		panic(e)
	}
}

// UnsplashURLs has a list of urls
type UnsplashURLs struct {
	Raw     string `json:"raw"`
	Full    string `json:"full"`
	Regular string `json:"regular"`
	Small   string `json:"small"`
	Thumb   string `json:"thumb"`
}

// UnsplashImage represents an image from https://unsplash.com
type UnsplashImage struct {
	URLs UnsplashURLs `json:"urls"`
}

func getUnsplashImages(rawurl string) []UnsplashImage {
	ret := &[]UnsplashImage{}

	_, err := url.ParseRequestURI(rawurl)
	Check(err)

	client := &http.Client{}
	req, err := http.NewRequest("GET", rawurl, nil)
	Check(err)

	req.Header.Add("User-Agent", "uipgo")
	resp, err := client.Do(req)
	Check(err)

	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(ret)
	Check(err)

	return *ret
}

// GetAndStoreImages downloads and stores images from given websites.
func GetAndStoreImages(sites map[string][]string, c *cli.Context) {
	images := []Image{}

	list, ok := sites["unsplash"]
	if ok {
		for _, site := range list {
			unsplashImages := getUnsplashImages(site)
			for _, image := range unsplashImages {
				parts := strings.Split(image.URLs.Regular, "/")
				index := strings.Index(parts[len(parts)-1], "?")
				identifier := parts[len(parts)-1][:index]
				name := fmt.Sprintf("unsplash-%s.jpg", identifier)
				images = append(images, Image{image.URLs.Regular, name})
			}
		}
	}

	for _, image := range images {
		err := DownloadFile(c.String("directory"), image.name, image.url)
		if err != nil {
			log.Println(err)
		} else {
			fmt.Println("Downloaded image: " + image.name)
		}
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
