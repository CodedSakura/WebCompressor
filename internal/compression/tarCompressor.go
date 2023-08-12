package compression

import (
	"WebCompressor/internal/configuration"
	"archive/tar"
	"io"
	"io/fs"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"
)

type TarCompressor struct {
	config *configuration.Configuration
}

func NewTarCompressor(config *configuration.Configuration) *TarCompressor {
	return &TarCompressor{config: config}
}

func (c *TarCompressor) Mimetype() string {
	return "application/x-tar"
}
func (c *TarCompressor) Extension() string {
	return "tar"
}
func (c *TarCompressor) Compress(targetPath string) (*State, error) {
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

		w := tar.NewWriter(file)
		defer w.Close()

		tarRootPath := path.Join(c.config.RootPath, targetPath)
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

				trimmedPath := strings.TrimPrefix(path, tarRootPath)
				if info.IsDir() {
					trimmedPath = trimmedPath + "/"
				}
				trimmedPath = strings.TrimPrefix(trimmedPath, "/")
				trimmedPath = "./" + trimmedPath
				header.Name = trimmedPath

				err = w.WriteHeader(header)
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

				_, err = io.Copy(w, file)
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
