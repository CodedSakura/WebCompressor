package endpoints

import (
	"github.com/gin-gonic/gin"
	"os"
)

type statusEndpoint struct {
}

//goland:noinspection GoExportedFuncWithUnexportedType
func NewStatusEndpoint() *statusEndpoint {
	return &statusEndpoint{}
}

func (e *statusEndpoint) Handle(c *gin.Context) {
	id := c.Param("id")

	if state, ok := activeStates[id]; ok {
		obj := gin.H{
			"id":        state.Id,
			"createdAt": state.CreatedTime,
			"progress":  state.Progress,
		}

		if state.IsDone() {
			obj["finishedAt"] = state.FinishedTime
		}
		if state.HasSucceeded() {
			obj["downloadUrl"] = "/download/" + id
			stat, err := os.Stat(state.Path)
			if err != nil {
				c.AbortWithStatus(500)
				return
			}
			obj["downloadSize"] = stat.Size()
		}

		c.JSON(200, obj)
		return
	}

	c.AbortWithStatus(404)
}
func (*statusEndpoint) Path() string {
	return "/status/:id"
}
func (*statusEndpoint) Method() string {
	return "GET"
}
