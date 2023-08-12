package main

import (
	"WebCompressor/internal/compression"
	"WebCompressor/internal/configuration"
	"WebCompressor/internal/endpoints"
	"WebCompressor/internal/http"
	"WebCompressor/internal/repository"
	"WebCompressor/internal/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		fx.Provide(
			fx.Annotate(
				http.New,
				fx.ParamTags(`group:"endpoints"`),
			),
			endpoints.AsEndpoint(endpoints.NewCompressEndpoint),
			endpoints.AsEndpoint(endpoints.NewDownloadEndpoint),
			endpoints.AsEndpoint(endpoints.NewFolderStateEndpoint),
			endpoints.AsEndpoint(endpoints.NewFolderViewEndpoint),
			endpoints.AsEndpoint(endpoints.NewRawEndpoint),
			endpoints.AsEndpoint(endpoints.NewStatusEndpoint),
			//api.New,
			//view.New,
			repository.New,
			utils.New,
			configuration.Read,
			compression.NewRegistry,
		),
		fx.Invoke(
			func(_ *gin.Engine, utils *utils.Utils, registry *compression.CompressorRegistry) {
				registry.RegisterDefault(utils)
			},
			//func(*gin.Engine) {},
		),
	).Run()
}
