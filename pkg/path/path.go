package path

import (
	"fmt"
	"os"
	"path/filepath"
)

type Folder = string

const (
	ConfigPath      Folder = "config"
	CachePath       Folder = "cache"
	HomePath        Folder = "home"
	ApplicationName Folder = "Matheus"
)

var FolderMap map[Folder]func() (string, error) = map[Folder]func() (string, error){
	ConfigPath: os.UserConfigDir,
	CachePath:  os.UserCacheDir,
	HomePath:   os.UserHomeDir,
}

func GetPath(folder Folder) string {
	function, hasProperty := FolderMap[folder]
	if !hasProperty {
		panic(fmt.Errorf("there is no folder available for %s", folder))
	}
	dir, err := function()
	if err != nil {
		panic(err)
	}
	return CreateDirectory(dir)
}

func CreateDirectory(dir string) string {
	return filepath.Join(dir, ApplicationName)
}
