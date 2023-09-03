package services

import (
	sysErrors "errors"
	"fmt"
	"io/fs"
	"os"

	"github.com/nimbo1999/environment-setup/internal/entities"
	"github.com/nimbo1999/environment-setup/pkg/errors"
	"github.com/nimbo1999/environment-setup/pkg/path"
)

var ErrInvalidProfile = sysErrors.New("invalid profile at the list")

type ProfileService struct{}

func NewProfileService() *ProfileService {
	return &ProfileService{}
}

func (service *ProfileService) Create(entityName string) error {
	if _, err := entities.GetProfile(entityName); err != entities.ErrProfileNotFound {
		return entities.ErrProfileAlreadyExists
	}

	profile := entities.Profile{
		Name: entityName,
	}

	if err := os.MkdirAll(profile.GetDirectory(), 0766); err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func (service *ProfileService) Update(profile entities.Profile) error {
	return errors.ErrMethodNotImplementedYet
}

func (service *ProfileService) List() ([]fs.DirEntry, error) {
	profileDirectory := path.GetProfilesPath()
	entries, err := os.ReadDir(profileDirectory)
	if err != nil {
		return nil, err
	}
	return entries, nil
}

func (service *ProfileService) Remove(profile entities.Profile) error {
	dir := profile.GetDirectory()
	return os.RemoveAll(dir)
}
