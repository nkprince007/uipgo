package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"strings"
)

func main() {
	dir := flag.String("d", ".", "the directory to store images")
	flag.Parse()

	args := flag.Args()

	if len(args) < 1 {
		flag.Usage()

		fmt.Println("Please provide at least one image URL as argument")
		os.Exit(1)
	}

	for i := range args {
		parts := strings.Split(args[i], "/")
		filename := parts[len(parts)-1]
		filepath := path.Join(*dir, filename)
		err := DownloadFile(filepath, args[i])
		if err != nil {
			log.Fatal(err)
		}
	}

	fmt.Println()
}

// DownloadFile downloads a file from the given url and stores it in filepath
func DownloadFile(filepath string, url string) (err error) {
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	resp, err := http.Get(url)
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
