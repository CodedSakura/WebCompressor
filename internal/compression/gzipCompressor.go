package compression

import (
	"WebCompressor/internal/configuration"
	"WebCompressor/internal/directorySize"
	"archive/tar"
	"compress/gzip"
	"io"
	"io/fs"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"
)

type GZipCompressor struct {
	config *configuration.Configuration
}

func NewGZipCompressor(config *configuration.Configuration) *GZipCompressor {
	return &GZipCompressor{config: config}
}

func (c *GZipCompressor) Mimetype() string {
	return "application/gzip"
}
func (c *GZipCompressor) Extension() string {
	return "tar.gz"
}
func (c *GZipCompressor) Compress(targetPath string) (*State, error) {
	state := newState(c)

	go func() {
		// https://medium.com/@skdomino/taring-untaring-files-in-go-6b07cf56bc07
		// https://stackoverflow.com/a/39647084/8672525
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

		g := gzip.NewWriter(file)
		defer g.Close()

		w := tar.NewWriter(g)
		defer w.Close()

		tarRootPath := path.Join(c.config.RootPath, targetPath)

		totalFileCount, err := directorySize.CountFiles(tarRootPath)
		processedFiles := 0
		if err != nil {
			state.Progress = -1
			state.FinishedTime = time.Now()
			return
		}

		err = filepath.Walk(
			tarRootPath,
			func(path string, info fs.FileInfo, err error) error {
				if err != nil {
					return err
				}

				header, err := tar.FileInfoHeader(info, info.Name())
				if err != nil {
					return err
				}

				trimmedPath := strings.TrimPrefix(
					path,
					filepath.Dir(tarRootPath),
				)
				if info.IsDir() {
					trimmedPath = trimmedPath + "/"
				}
				trimmedPath = strings.TrimPrefix(trimmedPath, "/")
				trimmedPath = "./" + trimmedPath
				header.Name = trimmedPath

				if err := w.WriteHeader(header); err != nil {
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

				if _, err := io.Copy(w, file); err != nil {
					return err
				}
				processedFiles += 1
				state.Progress = float32(processedFiles) / float32(totalFileCount)
				return nil
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
