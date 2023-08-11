package view

import "github.com/gin-gonic/gin"

func (v *View) RawView(c *gin.Context) {
	pathParam := c.Param("path")

	if !v.utils.FileExists(pathParam) {
		c.AbortWithStatus(404)
		return
	}

	fullPath := v.utils.GetAbsolutePath(pathParam)

	c.File(fullPath)
}
