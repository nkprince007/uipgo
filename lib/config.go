package lib

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/user"
	"path/filepath"

	"github.com/urfave/cli"
)

// ConfigPath contains the path for uipgo configuration.
var ConfigPath = filepath.Join(GetUserHomeDir(), ".uipgo.json")

// Settings contains the configuration for uipgo.
type Settings struct {
	StoragePath string `json:"storage_path"`
}

// Fetch retrieves the configuration for uipgo.
func (s *Settings) Fetch() (*Settings, error) {
	bytes, err := ioutil.ReadFile(ConfigPath)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(bytes, s)
	return s, err
}

// Store sets the configuration for uipgo.
func (s *Settings) Store() (*Settings, error) {
	bytes, err := json.Marshal(*s)
	if err != nil {
		return nil, err
	}
	err = ioutil.WriteFile(ConfigPath, bytes, 0777)
	return s, err
}

// GetUserHomeDir returns the current user's home directory.
func GetUserHomeDir() string {
	// notest

	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	return usr.HomeDir
}

// FetchConfig retrieves intial configuration into the Settings object.
func FetchConfig(ctx *cli.Context, conf *Settings) {
	// notest

	var path string
	reader := bufio.NewReader(os.Stdin)

	if chk, _ := conf.Fetch(); chk == nil {
		fmt.Println("It looks like you've launched uipgo for the first time.")
		for {
			fmt.Print("Where would you like to store images [~/.uipgo] ?")
			text, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println("\nInvalid configuration. Please try again!")
				continue
			}
			switch text {
			case "", "\n":
				path = ctx.String("directory")
			default:
				path = text
			}
			break
		}
		conf.StoragePath = path
		conf.Store()
	}
}
