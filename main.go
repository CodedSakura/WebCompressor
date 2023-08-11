package main

import (
	"WebCompressor/pkg/configuration"
	"WebCompressor/pkg/repository"
	"WebCompressor/pkg/view"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

func main() {
	config := configuration.Read()

	err := config.Verify()
	if err != nil {
		panic(err)
	}

	repo := repository.NewRepository(config)

	viewI := view.NewView(config, repo)

	gin.ForceConsoleColor()

	r := gin.Default()

	r.LoadHTMLGlob("templates/*")

	r.Static("/assets", "./assets")

	r.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/view/")
	})

	r.GET("/view/*path", viewI.FolderView)

	err = r.Run()
	if err != nil {
		println("Failed to start webserver")
		os.Exit(1)
		return
	}
}
