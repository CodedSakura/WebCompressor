package endpoints

import (
	"github.com/gin-gonic/gin"
)

type statusEndpoint struct {
}

//goland:noinspection GoExportedFuncWithUnexportedType
func NewStatusEndpoint() *statusEndpoint {
	return &statusEndpoint{}
}

func (e *statusEndpoint) Handle(c *gin.Context) {
	c.AbortWithStatus(501)
}
func (*statusEndpoint) Path() string {
	return "/status/:id"
}
func (*statusEndpoint) Method() string {
	return "GET"
}
