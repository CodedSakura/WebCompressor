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

	if state, ok := activeStates[id]; ok {
		if state.HasSucceeded() {
			c.File(state.Path)
			return
		}
		if state.HasFailed() {
			c.AbortWithStatus(400)
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
