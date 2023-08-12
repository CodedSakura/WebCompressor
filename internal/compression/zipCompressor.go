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
	"time"
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
func (c *ZipCompressor) Compress(targetPath string) (*State, error) {
	state := newState(c)

	go func() {
		// https://stackoverflow.com/a/63233911/8672525
		err := os.MkdirAll(filepath.Dir(state.Path), 0775)
		if err != nil {
			state.Progress = -1
			state.FinishedTime = time.Now()
			return
		}

		file, err := os.Create(state.Path)
		if err != nil {
			state.Progress = -1
			state.FinishedTime = time.Now()
			return
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

				trimmedPath := strings.TrimPrefix(path, zipRootPath)
				trimmedPath = strings.TrimPrefix(trimmedPath, "/")
				f, err := w.Create(trimmedPath)
				if err != nil {
					return err
				}

				_, err = io.Copy(f, file)
				return err
			},
		)
		state.Progress = 1
		state.FinishedTime = time.Now()
		if err != nil {
			state.Progress = -1
			return
		}
	}()

	return state, nil
}
