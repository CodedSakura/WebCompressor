package view

import "github.com/gin-gonic/gin"

func (v *View) Download(c *gin.Context) {
	c.AbortWithStatus(501)
}
