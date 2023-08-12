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
			obj := gin.H{
				"id":       state.Id,
				"cratedAt": state.CreatedTime,
				"progress": state.Progress,
			}

			if state.IsDone() {
				obj["finishedAt"] = state.FinishedTime
			}
			if state.HasSucceeded() {
				obj["downloadUrl"] = "/download/" + id
			}

			c.JSON(200, obj)
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
