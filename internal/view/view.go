package view

import (
	"WebCompressor/internal/compression"
	"WebCompressor/internal/repository"
	"WebCompressor/internal/utils"
)

type View struct {
	repository         *repository.Repository
	compressorRegistry *compression.CompressorRegistry
	utils              *utils.Utils
}

func New(
	repository *repository.Repository,
	compressorRegistry *compression.CompressorRegistry,
	utils *utils.Utils,
) *View {
	return &View{
		repository:         repository,
		compressorRegistry: compressorRegistry,
		utils:              utils,
	}
}
