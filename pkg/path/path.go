package path

import (
	"os"
	"path/filepath"
)

type Folder = string

const (
	ConfigPath      Folder = "config"
	CachePath       Folder = "cache"
	HomePath        Folder = "home"
	ApplicationName Folder = "Matheus"
	ProfilesPath    Folder = "profiles"
)

var FolderMap map[Folder]func() (string, error) = map[Folder]func() (string, error){
	ConfigPath: os.UserConfigDir,
	CachePath:  os.UserCacheDir,
	HomePath:   os.UserHomeDir,
}

func GetApplicationPaths(folder Folder) string {
	function, hasProperty := FolderMap[folder]
	if !hasProperty {
		return ""
	}
	dir, err := function()
	if err != nil {
		return ""
	}
	return CreateDirectory(dir)
}

func CreateDirectory(dir string) string {
	return filepath.Join(dir, ApplicationName)
}

func GetProfilesPath() string {
	return filepath.Join(GetApplicationPaths(ConfigPath), ProfilesPath)
}
