package api

import "github.com/gin-gonic/gin"

func (a *Api) View(c *gin.Context) {
	pathParam := c.Param("path")

	path, folders, files, err := a.repository.GetDirectoryContents(pathParam)

	if err != nil {
		c.AbortWithStatus(404)
		return
	}

	var compressors []gin.H
	for _, compressor := range a.compressorRegistry.Registered {
		compressors = append(compressors, gin.H{
			"extension": compressor.Extension(),
			"mimeType":  compressor.Mimetype(),
		})
	}

	c.JSON(200, gin.H{
		"path":        path,
		"folders":     folders,
		"files":       files,
		"compressors": compressors,
	})
}
