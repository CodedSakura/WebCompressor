package compression

type GZipCompressor struct {
}

func (c GZipCompressor) Mimetype() string {
	return "application/gzip"
}
func (c GZipCompressor) Extension() string {
	return "tar.gz"
}
func (c GZipCompressor) Compress(targetPath string) State {
	// placeholder
	return newState(c, targetPath)
}
