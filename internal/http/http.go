package http

import (
	"WebCompressor/internal/view"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

type Http struct {
	view *view.View
	gin  *gin.Engine
}

func New(view *view.View) *Http {
	gin.ForceConsoleColor()

	engine := gin.Default()

	return &Http{view: view, gin: engine}
}

func (h Http) RegisterPaths() {
	h.gin.LoadHTMLGlob("internal/view/*.tmpl")

	h.gin.Static("/assets", "./assets")

	h.gin.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/view/")
	})

	h.gin.GET("/view/*path", h.view.FolderView)
}

func (h Http) Run() {
	err := h.gin.Run()
	if err != nil {
		println("Failed to start webserver")
		os.Exit(1)
		return
	}
}
