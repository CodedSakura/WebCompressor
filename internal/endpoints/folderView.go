package endpoints

import (
	"WebCompressor/internal/compression"
	"WebCompressor/internal/configuration"
	"WebCompressor/internal/directorySize"
	"WebCompressor/internal/repository"
	"github.com/gin-gonic/gin"
	"os"
	path2 "path"
	"strings"
	"time"
)

type FolderViewEndpoint struct {
	repository         *repository.Repository
	config             *configuration.Configuration
	compressorRegistry *compression.CompressorRegistry
}

func NewFolderViewEndpoint(repository *repository.Repository, config *configuration.Configuration, compressorRegistry *compression.CompressorRegistry) *FolderViewEndpoint {
	return &FolderViewEndpoint{repository: repository, config: config, compressorRegistry: compressorRegistry}
}

type fileInfo struct {
	Name         string
	Size         int64
	LastModified time.Time
}

func (e *FolderViewEndpoint) Handle(c *gin.Context) {
	pathParam := c.Param("path")

	var path []string
	for _, pathPart := range strings.Split(pathParam, "/") {
		if len(pathPart) > 0 {
			path = append(path, pathPart)
		}
	}

	dir, err := os.ReadDir(
		path2.Join(e.config.RootPath, pathParam),
	)
	if err != nil {
		c.AbortWithStatus(404)
		return
	}

	var folders, files []fileInfo

	for _, entry := range dir {
		info, err := entry.Info()
		if err != nil {
			c.AbortWithStatus(404)
			return
		}
		fileInfo := fileInfo{
			Name:         info.Name(),
			Size:         info.Size(),
			LastModified: info.ModTime(),
		}
		if entry.IsDir() {
			fileInfo.Size, err = directorySize.Calculate(
				path2.Join(e.config.RootPath, pathParam, fileInfo.Name),
			)
			if err != nil {
				c.AbortWithStatus(404)
				return
			}
			folders = append(folders, fileInfo)
		} else {
			files = append(files, fileInfo)
		}
	}

	println("ff: ", len(folders), len(files))

	c.HTML(200, "folderView.tmpl", gin.H{
		"path":        path,
		"folders":     folders,
		"files":       files,
		"compressors": e.compressorRegistry.Registered,
	})
}
func (*FolderViewEndpoint) Path() string {
	return "/view/*path"
}
func (*FolderViewEndpoint) Method() string {
	return "GET"
}
