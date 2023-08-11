package compression

type TarCompressor struct {
}

func (c *TarCompressor) Mimetype() string {
	return "application/x-tar"
}
func (c *TarCompressor) Extension() string {
	return "tar"
}
func (c *TarCompressor) Compress(targetPath string) State {
	// placeholder
	return newState(c, targetPath)
}
