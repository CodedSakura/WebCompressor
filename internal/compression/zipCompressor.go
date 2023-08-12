package compression

import "WebCompressor/internal/utils"

type ZipCompressor struct {
	compressorBase
}

func NewZipCompressor(utils *utils.Utils) *ZipCompressor {
	return &ZipCompressor{
		compressorBase{
			utils: utils,
		},
	}
}

func (c *ZipCompressor) Mimetype() string {
	return "application/zip"
}
func (c *ZipCompressor) Extension() string {
	return "zip"
}
func (c *ZipCompressor) Compress(targetPath string) (State, error) {
	// placeholder
	state := newState(c)
	return state, nil
}
