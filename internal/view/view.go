package view

import (
	"WebCompressor/internal/configuration"
	"WebCompressor/internal/repository"
	"github.com/gin-gonic/gin"
)

type View struct {
	configuration configuration.Configuration
	repository    repository.Repository
}

func NewView(configuration configuration.Configuration, repository repository.Repository) View {
	return View{configuration: configuration, repository: repository}
}

func (v View) FolderView(c *gin.Context) {
	pathParam := c.Param("path")

	path, folders, files, err := v.repository.GetDirectoryContents(pathParam)

	if err != nil {
		c.AbortWithStatus(404)
		return
	}

	c.HTML(200, "folder-view.tmpl", gin.H{
		"path":    path,
		"folders": folders,
		"files":   files,
	})
}
