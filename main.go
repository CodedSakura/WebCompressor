package main

import (
	"WebCompressor/internal/api"
	"WebCompressor/internal/compression"
	"WebCompressor/internal/configuration"
	"WebCompressor/internal/http"
	"WebCompressor/internal/repository"
	"WebCompressor/internal/utils"
	"WebCompressor/internal/view"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		fx.Provide(
			http.New,
			api.New,
			view.New,
			repository.New,
			utils.New,
			configuration.Read,
			compression.NewRegistry,
		),
		fx.Invoke(
			func(_ *gin.Engine, utils *utils.Utils, registry *compression.CompressorRegistry) {
				registry.RegisterDefault(utils)
			},
		),
	).Run()
}
