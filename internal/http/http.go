package http

import (
	"WebCompressor/internal/api"
	"WebCompressor/internal/compression"
	"WebCompressor/internal/view"
	"context"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	"net/http"
	"os"
)

type Http struct {
	view               *view.View
	api                *api.Api
	compressorRegistry *compression.CompressorRegistry
	gin                *gin.Engine
}

func New(lc fx.Lifecycle, api *api.Api, view *view.View) *gin.Engine {
	gin.ForceConsoleColor()

	engine := gin.Default()

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			engine.LoadHTMLGlob("internal/view/*.tmpl")

			registerPaths(engine, api, view)

			go engine.Run()
			return nil
		},
	})

	return engine
}

func registerPaths(engine *gin.Engine, api *api.Api, view *view.View) {
	engine.LoadHTMLGlob("internal/view/*.tmpl")

	engine.Static("/assets", "./assets")

	engine.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/view/")
	})

	engine.GET("/view/*path", view.FolderView)
	engine.GET("/raw/*path", view.RawView)

	engine.GET("/download/:id", view.Download)

	apiGroup := engine.Group("/api")
	{
		apiGroup.GET("/view/*path", api.View)
		apiGroup.POST("/compress/:extension", api.Compress)
		apiGroup.GET("/status/:id", api.Status)
	}
}

func (h *Http) Run() {
	err := h.gin.Run()
	if err != nil {
		println("Failed to start webserver")
		os.Exit(1)
		return
	}
}
