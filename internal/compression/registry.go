package compression

type CompressorRegistry struct {
	Registered []Compressor
}

func NewRegistry(compressors []Compressor) *CompressorRegistry {
	registry := CompressorRegistry{Registered: []Compressor{}}
	println("reg: ", compressors, len(compressors))
	registry.Register(compressors...)
	return &registry
}

func (r *CompressorRegistry) Register(compressors ...Compressor) {
	r.Registered = append(r.Registered, compressors...)
}

func (r *CompressorRegistry) GetByExtension(ext string) *Compressor {
	for _, compressor := range r.Registered {
		if compressor.Extension() == ext {
			return &compressor
		}
	}
	return nil
}
