package services

import "io/fs"

type Service interface {
	Change(profile string) error
	List() ([]fs.DirEntry, error)
	Remove() error
}
