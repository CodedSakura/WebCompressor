package main

import (
	"WebCompressor/internal/configuration"
	"WebCompressor/internal/repository"
	"WebCompressor/internal/view"
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

	repo := repository.New(config)

	viewI := view.New(repo)

	gin.ForceConsoleColor()

	r := gin.Default()

	r.LoadHTMLGlob("internal/view/*.tmpl")

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
