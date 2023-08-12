package compression

import "WebCompressor/internal/utils"

type GZipCompressor struct {
	compressorBase
}

func NewGZipCompressor(utils *utils.Utils) *GZipCompressor {
	return &GZipCompressor{
		compressorBase{
			utils: utils,
		},
	}
}

func (c *GZipCompressor) Mimetype() string {
	return "application/gzip"
}
func (c *GZipCompressor) Extension() string {
	return "tar.gz"
}
func (c *GZipCompressor) Compress(targetPath string) State {
	// placeholder
	return newState(c)
}
