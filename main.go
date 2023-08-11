package main

import (
	"WebCompressor/internal/compression"
	"WebCompressor/internal/configuration"
	"WebCompressor/internal/http"
	"WebCompressor/internal/repository"
	"WebCompressor/internal/utils"
	"WebCompressor/internal/view"
)

func main() {
	config := configuration.Read()

	err := config.Verify()
	if err != nil {
		panic(err)
	}

	registry := compression.NewRegistry()
	registry.RegisterDefault()

	utilsI := utils.New(config)

	repo := repository.New(utilsI)

	viewI := view.New(repo, registry)

	httpI := http.New(viewI, registry)
	httpI.RegisterPaths()
	httpI.Run()
}
