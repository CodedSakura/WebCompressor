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
	id := c.Param("id")

	for _, state := range activeStates {
		if state.Id.String() == id {
			c.File(state.Path)
			return
		}
	}

	c.AbortWithStatus(404)
}
func (*downloadEndpoint) Path() string {
	return "/download/:id/*filename"
}
func (*downloadEndpoint) Method() string {
	return "GET"
}
