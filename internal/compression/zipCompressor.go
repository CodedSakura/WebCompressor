package compression

type ZipCompressor struct {
}

func (c *ZipCompressor) Mimetype() string {
	return "application/zip"
}
func (c *ZipCompressor) Extension() string {
	return "zip"
}
func (c *ZipCompressor) Compress(targetPath string) State {
	// placeholder
	return newState(c, targetPath)
}
