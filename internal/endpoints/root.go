package endpoints

import "github.com/gin-gonic/gin"

type rootEndpoint struct {
}

//goland:noinspection GoExportedFuncWithUnexportedType
func NewRootEndpoint() *rootEndpoint {
	return &rootEndpoint{}
}

func (e *rootEndpoint) Handle(c *gin.Context) {
	c.Redirect(301, "/view")
}
func (*rootEndpoint) Path() string {
	return "/"
}
func (*rootEndpoint) Method() string {
	return "GET"
}
