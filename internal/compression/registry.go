package compression

type CompressorRegistry struct {
	Registered []Compressor
}

func NewRegistry() *CompressorRegistry {
	return &CompressorRegistry{Registered: []Compressor{}}
}

func (r *CompressorRegistry) Register(compressors ...Compressor) {
	r.Registered = append(r.Registered, compressors...)
}

func (r *CompressorRegistry) RegisterDefault() {
	r.Registered = append(r.Registered, &ZipCompressor{}, &TarCompressor{}, &GZipCompressor{})
}
