package main

import (
	"WebCompressor/internal/compression"
	"WebCompressor/internal/configuration"
	"WebCompressor/internal/endpoints"
	"WebCompressor/internal/http"
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
			endpoints.AsEndpoint(endpoints.NewRootEndpoint),
			endpoints.AsEndpoint(endpoints.NewStatusEndpoint),
			configuration.Read,
			fx.Annotate(
				compression.NewRegistry,
				fx.ParamTags(`group:"compressors"`),
			),
			compression.AsCompressor(compression.NewZipCompressor),
			compression.AsCompressor(compression.NewTarCompressor),
			compression.AsCompressor(compression.NewGZipCompressor),
		),
		fx.Invoke(
			func(*gin.Engine) {},
		),
	).Run()
}
