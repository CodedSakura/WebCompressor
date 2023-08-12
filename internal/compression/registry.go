package compression

import "WebCompressor/internal/utils"

type CompressorRegistry struct {
	Registered []Compressor
}

func NewRegistry() *CompressorRegistry {
	return &CompressorRegistry{Registered: []Compressor{}}
}

func (r *CompressorRegistry) Register(compressors ...Compressor) {
	r.Registered = append(r.Registered, compressors...)
}

func (r *CompressorRegistry) RegisterDefault(utils *utils.Utils) {
	r.Registered = append(
		r.Registered,
		NewZipCompressor(utils),
		NewTarCompressor(utils),
		NewGZipCompressor(utils),
	)
}

func (r *CompressorRegistry) GetByExtension(ext string) *Compressor {
	for _, compressor := range r.Registered {
		if compressor.Extension() == ext {
			return &compressor
		}
	}
	return nil
}
