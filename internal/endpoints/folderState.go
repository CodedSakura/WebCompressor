package endpoints

import (
	"WebCompressor/internal/compression"
	"WebCompressor/internal/configuration"
	"WebCompressor/internal/directorySize"
	"github.com/gin-gonic/gin"
	"os"
	path2 "path"
	"time"
)

type folderStateEndpoint struct {
	config             *configuration.Configuration
	compressorRegistry *compression.CompressorRegistry
}

//goland:noinspection GoExportedFuncWithUnexportedType
func NewFolderStateEndpoint(config *configuration.Configuration, compressorRegistry *compression.CompressorRegistry) *folderStateEndpoint {
	return &folderStateEndpoint{config: config, compressorRegistry: compressorRegistry}
}

type folderStateFileInfo struct {
	Name         string
	Size         int64
	LastModified time.Time
}

func (e *folderStateEndpoint) Handle(c *gin.Context) {
	pathParam := c.Param("path")
	path := path2.Join(e.config.RootPath, pathParam)

	dir, err := os.ReadDir(path)
	if err != nil {
		c.AbortWithStatus(404)
		return
	}

	var folders, files []folderStateFileInfo

	for _, entry := range dir {
		info, err := entry.Info()
		if err != nil {
			c.AbortWithStatus(404)
			return
		}
		fileInfo := folderStateFileInfo{
			Name:         info.Name(),
			Size:         info.Size(),
			LastModified: info.ModTime(),
		}
		if entry.IsDir() {
			fileInfo.Size, err = directorySize.Calculate(
				path2.Join(path, fileInfo.Name),
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

	var compressors []gin.H
	for _, compressor := range e.compressorRegistry.Registered {
		compressors = append(compressors, gin.H{
			"extension": compressor.Extension(),
			"mimeType":  compressor.Mimetype(),
		})
	}

	c.JSON(200, gin.H{
		"folders":     folders,
		"files":       files,
		"compressors": compressors,
	})
}
func (*folderStateEndpoint) Path() string {
	return "/state/*path"
}
func (*folderStateEndpoint) Method() string {
	return "GET"
}
