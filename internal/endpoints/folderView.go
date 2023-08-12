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

type folderViewEndpoint struct {
	repository         *repository.Repository
	config             *configuration.Configuration
	compressorRegistry *compression.CompressorRegistry
}

//goland:noinspection GoExportedFuncWithUnexportedType
func NewFolderViewEndpoint(repository *repository.Repository, config *configuration.Configuration, compressorRegistry *compression.CompressorRegistry) *folderViewEndpoint {
	return &folderViewEndpoint{repository: repository, config: config, compressorRegistry: compressorRegistry}
}

type folderViewFileInfo struct {
	Name         string
	Size         int64
	LastModified time.Time
}

func (e *folderViewEndpoint) Handle(c *gin.Context) {
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

	var folders, files []folderViewFileInfo

	for _, entry := range dir {
		info, err := entry.Info()
		if err != nil {
			c.AbortWithStatus(404)
			return
		}
		fileInfo := folderViewFileInfo{
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
func (*folderViewEndpoint) Path() string {
	return "/view/*path"
}
func (*folderViewEndpoint) Method() string {
	return "GET"
}
