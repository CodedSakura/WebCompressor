package endpoints

import (
	"WebCompressor/internal/compression"
	"github.com/gin-gonic/gin"
)

type compressEndpoint struct {
	registry *compression.CompressorRegistry
}

var activeStates []*compression.State

//goland:noinspection GoExportedFuncWithUnexportedType
func NewCompressEndpoint(registry *compression.CompressorRegistry) *compressEndpoint {
	return &compressEndpoint{registry: registry}
}

func (e *compressEndpoint) Handle(c *gin.Context) {
	extension := c.Param("extension")
	pathParam := c.Param("path")

	compressor := e.registry.GetByExtension(extension)

	state, err := (*compressor).Compress(pathParam)

	activeStates = append(activeStates, state)

	if err != nil {
		println(err)
		c.AbortWithStatus(404)
		return
	}

	c.JSON(200, state)
}
func (*compressEndpoint) Path() string {
	return "/compress/:extension/*path"
}
func (*compressEndpoint) Method() string {
	return "POST"
}
