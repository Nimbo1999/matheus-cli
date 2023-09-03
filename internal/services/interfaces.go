package services

import (
	"io/fs"

	"github.com/nimbo1999/environment-setup/internal/entities"
)

type Service interface {
	Create(entityName string) error
	Update(profile entities.Profile) error
	List() ([]fs.DirEntry, error)
	Remove(profile entities.Profile) error
}
