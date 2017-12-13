package lib

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
	"sync"
	"time"

	"github.com/gosuri/uiprogress"
	"github.com/urfave/cli"
)

// ShowDownloadProgress shows the download progress through the bar provided.
func ShowDownloadProgress(
	done chan int64, bar *uiprogress.Bar, path string, total int64,
) {
	// notest

	bar.PrependFunc(func(b *uiprogress.Bar) string {
		return fmt.Sprintf("%50s", filepath.Base(path))
	})

	for {
		select {
		case <-done:
			return
		default:
			file, err := os.Open(path)
			Check(err)

			fi, err := file.Stat()
			Check(err)

			size := fi.Size()

			if size == 0 {
				size = 1
			}

			percent := float64(size) / float64(total) * 40
			bar.Set(int(percent))
		}
		time.Sleep(50 * time.Millisecond)
	}
}

// GetAndStoreImages downloads and stores images from given websites.
func GetAndStoreImages(sites map[string][]string, c *cli.Context) {
	// notest

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

	// progress bars
	uiprogress.Start()

	for _, image := range images {
		wg.Add(1)
		path := filepath.Join(c.String("directory"), image.Name())
		bar := uiprogress.AddBar(40).AppendCompleted()
		go DownloadFile(path, image.URL(), bar, &wg)
	}
	wg.Wait()

	return
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
