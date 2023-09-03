package entities

import (
	"errors"
	"os"
	"path/filepath"

	"github.com/nimbo1999/environment-setup/pkg/path"
)

type ProfileConfigurationKey = string

var (
	SshConfigKey ProfileConfigurationKey = "ssh"
)

var (
	ErrProfileNotFound      = errors.New("this profile does not exist")
	ErrProfileAlreadyExists = errors.New("this profile already exist")
)

type Profile struct {
	Name           string
	Configurations map[ProfileConfigurationKey]Configuration
}

func GetProfile(name string) (*Profile, error) {
	if !existsProfile(name) {
		return nil, ErrProfileNotFound
	}
	return &Profile{
		Name: name,
	}, nil
}

func (profile *Profile) GetDirectory() string {
	return getProfilePath(profile.Name)
}

func (profile *Profile) GetConfiguration(configKey ProfileConfigurationKey) *Configuration {
	config, exists := profile.Configurations[configKey]
	if !exists {
		return nil
	}
	return &config
}

func existsProfile(profileName string) bool {
	profilePath := getProfilePath(profileName)
	stat, err := os.Stat(profilePath)
	if err != nil {
		return os.IsExist(err)
	}
	return stat.IsDir()
}

func getProfilePath(profileName string) string {
	configPath := path.GetProfilesPath()
	return filepath.Join(configPath, profileName)
}
