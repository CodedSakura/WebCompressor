package main

import (
	"WebCompressor/internal/api"
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

	utilsI := utils.New(config)

	registry := compression.NewRegistry()
	registry.RegisterDefault(utilsI)

	repo := repository.New(utilsI)

	viewI := view.New(repo, registry, utilsI)

	apiI := api.New(repo, registry, utilsI)

	httpI := http.New(viewI, apiI, registry)
	httpI.RegisterPaths()
	httpI.Run()
}
