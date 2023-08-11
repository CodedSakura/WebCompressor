package http

import (
	"WebCompressor/internal/api"
	"WebCompressor/internal/compression"
	"WebCompressor/internal/view"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

type Http struct {
	view               *view.View
	api                *api.Api
	compressorRegistry *compression.CompressorRegistry
	gin                *gin.Engine
}

func New(view *view.View, api *api.Api, compressorRegistry *compression.CompressorRegistry) *Http {
	gin.ForceConsoleColor()

	engine := gin.Default()

	return &Http{view: view, api: api, compressorRegistry: compressorRegistry, gin: engine}
}

func (h *Http) RegisterPaths() {
	h.gin.LoadHTMLGlob("internal/view/*.tmpl")

	h.gin.Static("/assets", "./assets")

	h.gin.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/view/")
	})

	h.gin.GET("/view/*path", h.view.FolderView)
	h.gin.GET("/raw/*path", h.view.RawView)

	h.gin.GET("/download/:id", h.view.Download)

	apiGroup := h.gin.Group("/api")
	{
		apiGroup.GET("/view/*path", h.api.View)
		apiGroup.POST("/compress/:extension", h.api.Compress)
		apiGroup.GET("/status/:id", h.api.Status)
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
