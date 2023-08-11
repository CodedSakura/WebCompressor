package api

import (
	"WebCompressor/internal/compression"
	"WebCompressor/internal/repository"
)

type Api struct {
	repository         *repository.Repository
	compressorRegistry *compression.CompressorRegistry
}

func New(repository *repository.Repository, compressorRegistry *compression.CompressorRegistry) *Api {
	return &Api{repository: repository, compressorRegistry: compressorRegistry}
}
