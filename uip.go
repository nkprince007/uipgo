package main

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"path"
	"strings"

	"github.com/urfave/cli"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func getVersion() string {
	data, err := ioutil.ReadFile("VERSION")
	check(err)

	return string(data)
}

func main() {
	var directory string

	app := cli.NewApp()
	app.Name = "UIP"
	app.Version = getVersion()
	app.Usage = "a tool to download wallpapers"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "directory",
			Value:       ".",
			Usage:       "directory to store wallpapers in",
			Destination: &directory,
		},
	}

	app.Action = func(c *cli.Context) error {
		if c.NArg() < 1 {
			cli.ShowAppHelpAndExit(c, 1)
		}

		args := c.Args()
		for i := range args {
			parts := strings.Split(args[i], "/")
			filename := parts[len(parts)-1]
			filepath := path.Join(directory, filename)
			err := DownloadFile(filepath, args[i])
			if err != nil {
				log.Fatal(err)
			}
		}

		return nil
	}

	app.Run(os.Args)
}

// DownloadFile downloads a file from the given url and stores it in filepath
func DownloadFile(filepath string, rawurl string) (err error) {
	parsedurl, err := url.ParseRequestURI(rawurl)
	if err != nil {
		return err
	}

	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	resp, err := http.Get(parsedurl.EscapedPath())
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
