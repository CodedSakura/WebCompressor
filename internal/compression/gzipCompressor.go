package compression

type GZipCompressor struct {
}

func NewGZipCompressor() *GZipCompressor {
	return &GZipCompressor{}
}

func (c *GZipCompressor) Mimetype() string {
	return "application/gzip"
}
func (c *GZipCompressor) Extension() string {
	return "tar.gz"
}
func (c *GZipCompressor) Compress(targetPath string) (State, error) {
	// placeholder
	return newState(c), nil
}
