package repository

import (
	"WebCompressor/internal/configuration"
	"errors"
	"os"
	"path"
	"strings"
)

type Repository struct {
	configuration *configuration.Configuration
}

func New(configuration *configuration.Configuration) *Repository {
	return &Repository{configuration: configuration}
}

func dirExists(pathName string) bool {
	stat, err := os.Stat(pathName)
	return err == nil && stat.IsDir()
}

func (r Repository) GetDirectoryContents(relativePath string) ([]string, []string, []string, error) {
	lookupPath := path.Join(r.configuration.RootPath, relativePath)

	if !dirExists(lookupPath) {
		return nil, nil, nil, errors.New("directory does not exist")
	}

	var pathParts []string
	for _, pathPart := range strings.Split(relativePath, "/") {
		if len(pathPart) > 0 {
			pathParts = append(pathParts, pathPart)
		}
	}

	// very unlikely to encounter error, as path exists
	dir, _ := os.ReadDir(lookupPath)

	var folders []string
	var files []string

	for _, entry := range dir {
		if entry.IsDir() {
			folders = append(folders, entry.Name())
		} else {
			files = append(files, entry.Name())
		}
	}

	return pathParts, folders, files, nil
}
