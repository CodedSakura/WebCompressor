package compression

import "WebCompressor/internal/utils"

type TarCompressor struct {
	compressorBase
}

func NewTarCompressor(utils *utils.Utils) *TarCompressor {
	return &TarCompressor{
		compressorBase{
			utils: utils,
		},
	}
}

func (c *TarCompressor) Mimetype() string {
	return "application/x-tar"
}
func (c *TarCompressor) Extension() string {
	return "tar"
}
func (c *TarCompressor) Compress(targetPath string) State {
	// placeholder
	return newState(c)
}
