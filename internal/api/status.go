package api

import "github.com/gin-gonic/gin"

func (a *Api) Status(c *gin.Context) {
	c.AbortWithStatus(501)
}
