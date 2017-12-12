// uipgo is a rewrite of UIP package from python to download wallpapers from
// chosen sites.

package main

import (
	"io/ioutil"
	"log"
	"os"
	"os/user"
	"path/filepath"

	// third party imports
	"github.com/nkprince007/uipgo/lib"
	"github.com/urfave/cli"
)

// The client ID provided is free for development use for upto 50 requests/hr.
const unsplashClientID = "74f6347705c15665e0d3d4b241fce1e9c2ef26761aeddfe0724dcd00d2823af5"

// Websites is the list of precollected websites to download wallpapers from.
var Websites = map[string][]string{
	"unsplash": {
		"https://api.unsplash.com/photos?client_id=" + unsplashClientID,
	},
	"desktoppr": {
		"https://api.desktoppr.co/1/wallpapers",
	},
}

func getVersion() string {
	// notest

	data, err := ioutil.ReadFile("VERSION")
	lib.Check(err)

	return string(data)
}

func main() {
	// notest

	var directory string

	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	app := cli.NewApp()
	app.Name = "uipgo"
	app.Version = getVersion()
	app.Usage = "a tool to download wallpapers"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "directory",
			Value:       filepath.Join(usr.HomeDir, ".uipgo"),
			Usage:       "directory to store wallpapers in",
			Destination: &directory,
		},
	}

	app.Action = func(c *cli.Context) error {
		lib.GetAndStoreImages(Websites, c)
		return nil
	}

	app.Run(os.Args)
}
