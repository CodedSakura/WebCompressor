package http

import (
	"WebCompressor/internal/endpoints"
	"context"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

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
