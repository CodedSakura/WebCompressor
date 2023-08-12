package directorySize

import (
	"io/fs"
	"path/filepath"
)

func Calculate(path string) (int64, error) {
	var totalSize int64
	err := filepath.Walk(path, func(_ string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		totalSize += info.Size()
		return err
	})
	return totalSize, err
}

func CountFiles(path string) (int, error) {
	var fileCount int
	err := filepath.Walk(path, func(_ string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		fileCount += 1
		return err
	})
	return fileCount, err
}
