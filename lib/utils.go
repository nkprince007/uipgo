package lib

import (
	"fmt"
	"os"
	"path/filepath"
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
