package http

import (
	"WebCompressor/internal/api"
	"WebCompressor/internal/compression"
	"WebCompressor/internal/endpoints"
	"WebCompressor/internal/view"
	"context"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

type Http struct {
	view               *view.View
	api                *api.Api
	compressorRegistry *compression.CompressorRegistry
	gin                *gin.Engine
}

func New(endpoints []endpoints.Endpoint, lc fx.Lifecycle) *gin.Engine {
	gin.ForceConsoleColor()

	engine := gin.Default()

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			engine.LoadHTMLGlob("internal/endpoints/*.tmpl")
			engine.Static("/assets", "./assets")

			for _, endpoint := range endpoints {
				engine.Handle(endpoint.Method(), endpoint.Path(), endpoint.Handle)
			}

			go engine.Run()
			return nil
		},
	})

	return engine
}
