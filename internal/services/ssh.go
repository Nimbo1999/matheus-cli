package services

import (
	"errors"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/nimbo1999/environment-setup/pkg/path"
	cp "github.com/otiai10/copy"
)

type SSHService struct {
	FolderName string
}

var (
	ErrProfileDoesNotExists = errors.New("provided profile does not exists")
)

func NewSSHService(folderName string) *SSHService {
	return &SSHService{
		FolderName: folderName,
	}
}

func (service *SSHService) List() ([]fs.DirEntry, error) {
	sshPath := service.GetStoragedSSHPath()
	dir, err := os.Open(sshPath)
	if err != nil {
		return nil, err
	}
	defer dir.Close()
	return dir.ReadDir(0)
}

func (service *SSHService) Change(profile string) error {
	profilePath, err := service.GetProfilePath(profile)
	if err != nil {
		return err
	}

	// Removes the .ssh folder under the Home directory, that indicates the current
	// active profile.
	if err := service.Remove(); err != nil {
		return err
	}

	return service.Use(profilePath)
}

func (service *SSHService) Remove() error {
	sshPath, err := service.GetUserSSHPath()
	if err != nil {
		return err
	}
	return os.RemoveAll(sshPath)
}

func (service *SSHService) GetStoragedSSHPath() string {
	return filepath.Join(path.GetPath(path.ConfigPath), service.FolderName)
}

func (service *SSHService) GetProfilePath(profile string) (string, error) {
	sshPath := service.GetStoragedSSHPath()
	profilePath := filepath.Join(sshPath, profile)
	if _, err := os.Stat(profilePath); os.IsNotExist(err) {
		return "", ErrProfileDoesNotExists
	}
	return profilePath, nil
}

func (service *SSHService) GetUserSSHPath() (string, error) {
	homePath, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	sshPath := filepath.Join(homePath, ".ssh")
	return sshPath, nil
}

func (service *SSHService) Use(profilePath string) error {
	sshPath, err := service.GetUserSSHPath()
	if err != nil {
		return err
	}

	if err = cp.Copy(profilePath, sshPath); err != nil {
		return err
	}
	return nil
}
