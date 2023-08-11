package repository

import (
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
	Size     int64
	Children int
}

func (r *Repository) directoryFromDirEntry(relativePath string, entry os.DirEntry) (*directory, error) {
	file, err := r.fileFromDirEntry(entry)
	if err != nil {
		return nil, err
	}
	size, children, err := r.utils.CalculateDir(path.Join(relativePath, entry.Name()))
	if err != nil {
		return nil, err
	}

	return &directory{
		Children: children,
		Size:     size,
		file:     file,
	}, nil
}

func (r *Repository) fileFromDirEntry(entry os.DirEntry) (*file, error) {
	info, err := entry.Info()
	if err != nil {
		return nil, err
	}

	return &file{
		Name:   entry.Name(),
		Size:   info.Size(),
		Change: info.ModTime(),
	}, nil
}

func (r *Repository) GetDirectoryContents(relativePath string) ([]string, []*directory, []*file, error) {
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
			folder, err := r.directoryFromDirEntry(relativePath, entry)
			if err != nil {
				return nil, nil, nil, err
			}
			folders = append(folders, folder)
		} else {
			file, err := r.fileFromDirEntry(entry)
			if err != nil {
				return nil, nil, nil, err
			}
			files = append(files, file)
		}
	}

	return pathParts, folders, files, nil
}
