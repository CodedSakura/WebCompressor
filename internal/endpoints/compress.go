package endpoints

import (
	"github.com/gin-gonic/gin"
)

type compressEndpoint struct {
}

//goland:noinspection GoExportedFuncWithUnexportedType
func NewCompressEndpoint() *compressEndpoint {
	return &compressEndpoint{}
}

func (e *compressEndpoint) Handle(c *gin.Context) {
	c.AbortWithStatus(501)
}
func (*compressEndpoint) Path() string {
	return "/compress/*path"
}
func (*compressEndpoint) Method() string {
	return "POST"
}
