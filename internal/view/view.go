package view

import (
	"WebCompressor/internal/configuration"
	"WebCompressor/internal/repository"
)

type View struct {
	configuration configuration.Configuration
	repository    repository.Repository
}

func NewView(configuration configuration.Configuration, repository repository.Repository) View {
	return View{configuration: configuration, repository: repository}
}
