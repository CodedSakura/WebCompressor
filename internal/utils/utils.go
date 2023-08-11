package utils

import (
	"WebCompressor/internal/configuration"
	"errors"
	"os"
	"path"
)

type Utils struct {
	configuration *configuration.Configuration
}

func New(configuration *configuration.Configuration) *Utils {
	return &Utils{configuration: configuration}
}

func (u Utils) DirExists(relativePath string) bool {
	stat, err := os.Stat(u.getAbsolutePath(relativePath))
	return err == nil && stat.IsDir()
}

func (u Utils) FileExists(relativePath string) bool {
	stat, err := os.Stat(u.getAbsolutePath(relativePath))
	return err == nil && !stat.IsDir()
}

func (u Utils) ReadDir(relativePath string) ([]os.DirEntry, error) {
	if !u.DirExists(relativePath) {
		return nil, errors.New("directory does not exist")
	}
	return os.ReadDir(u.getAbsolutePath(relativePath))
}

func (u Utils) getAbsolutePath(relativePath string) string {
	return path.Join(u.configuration.RootPath, relativePath)
}
