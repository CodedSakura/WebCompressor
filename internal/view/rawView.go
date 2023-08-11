package view

import "github.com/gin-gonic/gin"

func (v *View) RawView(c *gin.Context) {
	c.AbortWithStatus(501)
}
