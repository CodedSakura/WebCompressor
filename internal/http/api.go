package http

import "github.com/gin-gonic/gin"

func apiView(c *gin.Context) {
	c.AbortWithStatus(501)
}

func apiCompress(c *gin.Context) {
	c.AbortWithStatus(501)
}

func apiStatus(c *gin.Context) {
	c.AbortWithStatus(501)
}

func download(c *gin.Context) {
	c.AbortWithStatus(501)
}
