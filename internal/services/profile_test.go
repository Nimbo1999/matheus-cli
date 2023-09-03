package services

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/nimbo1999/environment-setup/internal/entities"
	"github.com/nimbo1999/environment-setup/pkg/path"
	"github.com/stretchr/testify/assert"
)

func TestProfileCreationAndRemoval(t *testing.T) {
	profileService := NewProfileService()
	err := profileService.Create("test")
	assert.NoError(t, err)

	err = profileService.Create("test")
	assert.Error(t, err, entities.ErrProfileAlreadyExists)

	testProfile, err := entities.GetProfile("test")
	assert.NoError(t, err)

	err = profileService.Remove(*testProfile)
	assert.NoError(t, err)

	profilePath := filepath.Join(path.GetProfilesPath(), "test")
	_, err = os.Stat(profilePath)
	assert.Error(t, err, os.ErrNotExist)
}

func TestListProfile(t *testing.T) {
	profileService := NewProfileService()
	err := profileService.Create("test")
	assert.NoError(t, err)

	err = profileService.Create("test2")
	assert.NoError(t, err)

	testProfile, err := entities.GetProfile("test")
	assert.NoError(t, err)
	test2Profile, err := entities.GetProfile("test2")
	assert.NoError(t, err)

	profileEntries, err := profileService.List()
	assert.NoError(t, err)

	profiles := []string{}

	for _, entry := range profileEntries {
		entryName := entry.Name()
		if entryName == testProfile.Name || entryName == test2Profile.Name {
			profiles = append(profiles, entryName)
		}
	}

	assert.Len(t, profiles, 2)
	err = profileService.Remove(*testProfile)
	assert.NoError(t, err)
	err = profileService.Remove(*test2Profile)
	assert.NoError(t, err)
}
