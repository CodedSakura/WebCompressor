package compression

type TarCompressor struct {
}

func NewTarCompressor() *TarCompressor {
	return &TarCompressor{}
}

func (c *TarCompressor) Mimetype() string {
	return "application/x-tar"
}
func (c *TarCompressor) Extension() string {
	return "tar"
}
func (c *TarCompressor) Compress(targetPath string) (*State, error) {
	// placeholder
	return newState(c), nil
}
