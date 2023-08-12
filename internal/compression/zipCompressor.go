package compression

import (
	"WebCompressor/internal/configuration"
	"archive/zip"
	"io"
	"io/fs"
	"os"
	"path"
	"path/filepath"
	"strings"
)

type ZipCompressor struct {
	config *configuration.Configuration
}

func NewZipCompressor(config *configuration.Configuration) *ZipCompressor {
	return &ZipCompressor{config: config}
}

func (c *ZipCompressor) Mimetype() string {
	return "application/zip"
}
func (c *ZipCompressor) Extension() string {
	return "zip"
}
func (c *ZipCompressor) Compress(targetPath string) (State, error) {
	state := newState(c)

	// https://stackoverflow.com/a/63233911/8672525
	file, err := os.Create(state.Path)
	if err != nil {
		return state, err
	}
	defer file.Close()

	w := zip.NewWriter(file)
	defer w.Close()

	zipRootPath := path.Join(c.config.RootPath, targetPath)
	err = filepath.Walk(
		zipRootPath,
		func(path string, info fs.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if info.IsDir() {
				return nil
			}
			file, err := os.Open(path)
			if err != nil {
				return err
			}
			defer file.Close()

			f, err := w.Create(strings.TrimPrefix(zipRootPath, path))
			if err != nil {
				return err
			}

			_, err = io.Copy(f, file)
			return err
		},
	)

	return state, err
}
