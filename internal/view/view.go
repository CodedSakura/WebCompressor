package view

import (
	"WebCompressor/internal/compression"
	"WebCompressor/internal/repository"
)

type View struct {
	repository         *repository.Repository
	compressorRegistry *compression.CompressorRegistry
}

func New(repository *repository.Repository, compressorRegistry *compression.CompressorRegistry) *View {
	return &View{repository: repository, compressorRegistry: compressorRegistry}
}
