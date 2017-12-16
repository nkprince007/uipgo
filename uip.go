// uipgo is a rewrite of UIP package from python to download wallpapers from
// chosen sites.

package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	// third party imports
	"github.com/nkprince007/uipgo/lib"
	"github.com/nkprince007/uipgo/utils"
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

func main() {
	// notest

	var directory string

	app := cli.NewApp()
	app.Name = "uipgo"
	app.Version = "0.0.2"
	app.Usage = "a tool to download wallpapers"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "directory",
			Value:       filepath.Join(lib.GetUserHomeDir(), ".uipgo"),
			Usage:       "directory to store wallpapers in",
			Destination: &directory,
		},
	}

	app.Commands = []cli.Command{
		cli.Command{
			Name:    "wallpaper",
			Aliases: []string{"wp"},
			Usage:   "to change your desktop wallpaper",
			Action: func(c *cli.Context) error {
				if c.NArg() == 1 {
					err := utils.SetWallpaper(c.Args()[0])
					if err != nil {
						log.Fatal(err)
					}
				} else if c.NArg() == 0 {
					wp, err := utils.GetWallpaper()
					if err != nil {
						log.Fatal(err)
					}
					fmt.Println(wp)
				}
				return nil
			},
		},
	}

	app.Action = func(ctx *cli.Context) error {
		conf := &lib.Settings{}
		lib.FetchConfig(ctx, conf)
		lib.GetAndStoreImages(Websites, conf.StoragePath)
		return nil
	}

	app.Run(os.Args)
}
