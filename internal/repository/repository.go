package repository

import (
	"WebCompressor/internal/configuration"
	"os"
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
