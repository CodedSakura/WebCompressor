package endpoints

import (
	"WebCompressor/internal/configuration"
	"github.com/gin-gonic/gin"
	"os"
	"path"
)

type rawEndpoint struct {
	config *configuration.Configuration
}

//goland:noinspection GoExportedFuncWithUnexportedType
func NewRawEndpoint(config *configuration.Configuration) *rawEndpoint {
	return &rawEndpoint{config: config}
}

func (e *rawEndpoint) Handle(c *gin.Context) {
	pathParam := c.Param("path")
	fullPath := path.Join(e.config.RootPath, pathParam)

	stat, err := os.Stat(fullPath)
	if err != nil || stat.IsDir() {
		c.AbortWithStatus(404)
		return
	}

	c.File(fullPath)
}
func (*rawEndpoint) Path() string {
	return "/raw/*path"
}
func (*rawEndpoint) Method() string {
	return "GET"
}
