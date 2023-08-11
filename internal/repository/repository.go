package repository

import (
	"WebCompressor/internal/utils"
)

type Repository struct {
	utils *utils.Utils
}

func New(utils *utils.Utils) *Repository {
	return &Repository{utils: utils}
}
