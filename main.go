package main

import (
	"WebCompressor/internal/configuration"
	"WebCompressor/internal/http"
	"WebCompressor/internal/repository"
	"WebCompressor/internal/view"
)

func main() {
	config := configuration.Read()

	err := config.Verify()
	if err != nil {
		panic(err)
	}

	repo := repository.New(config)

	viewI := view.New(repo)

	httpI := http.New(viewI)
	httpI.RegisterPaths()
	httpI.Run()
}
