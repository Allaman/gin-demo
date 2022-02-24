package main

import (
	"embed"
	"html/template"
	"io/fs"
	"net/http"

	"github.com/gin-gonic/gin"
)

//go:embed assets templates
var embeddedFiles embed.FS

func main() {
	router := gin.Default()
	templ := template.Must(template.New("").ParseFS(embeddedFiles, "templates/*"))
	router.SetHTMLTemplate(templ)

	router.StaticFS("/public", http.FS(embeddedFiles))

	router.GET("/favicon.png", func(c *gin.Context) {
		c.FileFromFS(".", FaviconFS())
	})

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "Hello Gin!",
		})
	})
	router.GET("/ping", func(c *gin.Context) {
		c.HTML(http.StatusOK, "ping.tmpl", gin.H{
			"title": "Pong",
		})
	})
	router.Run(":8080")
}

func FaviconFS() http.FileSystem {
	sub, err := fs.Sub(embeddedFiles, "assets/favicon.png")
	if err != nil {
		panic(err)
	}
	return http.FS(sub)
}
