package repository

import (
	"errors"
	"os"
	"path"
	"strings"
	"time"
)

type file struct {
	Name   string
	Size   int64
	Change time.Time
}
type directory struct {
	*file
	Children int
}

func directoryFromDirEntry(entry os.DirEntry) *directory {
	return &directory{
		Children: 0,
		file:     fileFromDirEntry(entry),
	}
}

func fileFromDirEntry(entry os.DirEntry) *file {
	info, err := entry.Info()
	if err != nil {
		panic(err)
	}
	return &file{
		Name:   entry.Name(),
		Size:   info.Size(),
		Change: info.ModTime(),
	}
}

func (r Repository) GetDirectoryContents(relativePath string) ([]string, []*directory, []*file, error) {
	lookupPath := path.Join(r.configuration.RootPath, relativePath)

	if !dirExists(lookupPath) {
		return nil, nil, nil, errors.New("directory does not exist")
	}

	var pathParts []string
	for _, pathPart := range strings.Split(relativePath, "/") {
		if len(pathPart) > 0 {
			pathParts = append(pathParts, pathPart)
		}
	}

	// very unlikely to encounter error, as path exists
	dir, _ := os.ReadDir(lookupPath)

	var folders []*directory
	var files []*file

	for _, entry := range dir {
		if entry.IsDir() {
			folders = append(folders, directoryFromDirEntry(entry))
		} else {
			files = append(files, fileFromDirEntry(entry))
		}
	}

	return pathParts, folders, files, nil
}
