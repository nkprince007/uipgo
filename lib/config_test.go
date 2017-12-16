package lib

import (
	"os"
	"path/filepath"
	"testing"
)

func TestSettings(t *testing.T) {
	ConfigPath = filepath.Join(os.TempDir(), ".uipgo.json")

	conf := &Settings{}
	conf.StoragePath = filepath.Dir(ConfigPath)

	_, err := conf.Store()
	if err != nil {
		t.Errorf("Settings.Store() encountered an error: %s", err)
	}

	_, err = conf.Fetch()
	if err != nil {
		t.Errorf("Settings.Fetch() encountered an error: %s", err)
	}

	if conf.StoragePath != filepath.Dir(ConfigPath) {
		t.Error("Settings.Fetch() or Settings.Store() failed.")
	}
}
