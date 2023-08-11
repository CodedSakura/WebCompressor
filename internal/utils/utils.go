package utils

import (
	"WebCompressor/internal/configuration"
	"errors"
	"io/fs"
	"os"
	"path"
	"path/filepath"
)

type Utils struct {
	configuration *configuration.Configuration
}

func New(configuration *configuration.Configuration) *Utils {
	return &Utils{configuration: configuration}
}

func (u *Utils) DirExists(relativePath string) bool {
	stat, err := os.Stat(u.GetAbsolutePath(relativePath))
	return err == nil && stat.IsDir()
}

func (u *Utils) FileExists(relativePath string) bool {
	stat, err := os.Stat(u.GetAbsolutePath(relativePath))
	return err == nil && !stat.IsDir()
}

func (u *Utils) ReadDir(relativePath string) ([]os.DirEntry, error) {
	if !u.DirExists(relativePath) {
		return nil, errors.New("directory does not exist")
	}
	return os.ReadDir(u.GetAbsolutePath(relativePath))
}

func (u *Utils) CalculateDir(relativePath string) (int64, int, error) {
	var size int64
	var fileCount int
	err := filepath.Walk(u.GetAbsolutePath(relativePath), func(_ string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		size += info.Size()
		if !info.IsDir() {
			fileCount++
		}
		return err
	})
	return size, fileCount, err
}

func (u *Utils) GetAbsolutePath(relativePath string) string {
	return path.Join(u.configuration.RootPath, relativePath)
}
