package repository

import (
	"os"
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
	var pathParts []string
	for _, pathPart := range strings.Split(relativePath, "/") {
		if len(pathPart) > 0 {
			pathParts = append(pathParts, pathPart)
		}
	}

	dir, err := r.utils.ReadDir(relativePath)
	if err != nil {
		return nil, nil, nil, err
	}

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
