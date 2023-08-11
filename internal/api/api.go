package api

import (
	"WebCompressor/internal/compression"
	"WebCompressor/internal/repository"
	"WebCompressor/internal/utils"
)

type Api struct {
	repository         *repository.Repository
	compressorRegistry *compression.CompressorRegistry
	utils              *utils.Utils
}

func New(
	repository *repository.Repository,
	compressorRegistry *compression.CompressorRegistry,
	utils *utils.Utils,
) *Api {
	return &Api{
		repository:         repository,
		compressorRegistry: compressorRegistry,
		utils:              utils,
	}
}
