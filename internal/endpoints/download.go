package endpoints

import (
	"github.com/gin-gonic/gin"
)

type downloadEndpoint struct {
}

//goland:noinspection GoExportedFuncWithUnexportedType
func NewDownloadEndpoint() *downloadEndpoint {
	return &downloadEndpoint{}
}

func (e *downloadEndpoint) Handle(c *gin.Context) {
	c.AbortWithStatus(501)
}
func (*downloadEndpoint) Path() string {
	return "/download/:id/*filename"
}
func (*downloadEndpoint) Method() string {
	return "GET"
}
