package api

import (
	"WebCompressor/internal/utils"
	"github.com/gin-gonic/gin"
)

type Api struct {
	utils *utils.Utils
}

func New(utils *utils.Utils) *Api {
	return &Api{utils: utils}
}

func (a *Api) View(c *gin.Context) {
	c.AbortWithStatus(501)
}

func (a *Api) Compress(c *gin.Context) {
	c.AbortWithStatus(501)
}

func (a *Api) Status(c *gin.Context) {
	c.AbortWithStatus(501)
}
