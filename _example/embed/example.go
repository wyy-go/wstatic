package main

import (
	"embed"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/wyy-go/wstatic"
)

//go:embed form
var web embed.FS

func main() {
	r := gin.Default()

	r.Use(wstatic.New(wstatic.WithUrlPrefix("/"),
		wstatic.WithRoot("/form"),
		wstatic.WithEmbed(true),
		wstatic.WithEmbedFs(web),
		wstatic.WithIndexes(false)))

	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "test")
	})

	// Listen and Server in 0.0.0.0:8080
	if err := r.Run(":8880"); err != nil {
		log.Fatal(err)
	}
}
