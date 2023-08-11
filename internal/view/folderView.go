package view

import "github.com/gin-gonic/gin"

func (v *View) FolderView(c *gin.Context) {
	pathParam := c.Param("path")

	path, folders, files, err := v.repository.GetDirectoryContents(pathParam)

	if err != nil {
		c.AbortWithStatus(404)
		return
	}

	c.HTML(200, "folderView.tmpl", gin.H{
		"path":        path,
		"folders":     folders,
		"files":       files,
		"compressors": v.compressorRegistry.Registered,
	})
}
