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
	id := c.Param("id")

	for _, state := range activeStates {
		if state.Id.String() == id {
			c.JSON(200, state)
			return
		}
	}

	c.AbortWithStatus(404)
}
func (*statusEndpoint) Path() string {
	return "/status/:id"
}
func (*statusEndpoint) Method() string {
	return "GET"
}
