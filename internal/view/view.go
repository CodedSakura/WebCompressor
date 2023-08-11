package view

import (
	"WebCompressor/internal/repository"
)

type View struct {
	repository *repository.Repository
}

func New(repository *repository.Repository) *View {
	return &View{repository: repository}
}
