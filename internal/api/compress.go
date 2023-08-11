package api

import "github.com/gin-gonic/gin"

func (a *Api) Compress(c *gin.Context) {
	c.AbortWithStatus(501)
}
