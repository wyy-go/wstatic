package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/wyy-go/wstatic"
)


func main() {
	r := gin.Default()

	r.Use(wstatic.New(wstatic.WithUrlPrefix("/"),
		wstatic.WithRoot("./form"),
		wstatic.WithIndexes(false)))

	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "test")
	})

	// Listen and Server in 0.0.0.0:8080
	if err := r.Run(":8880"); err != nil {
		log.Fatal(err)
	}
}
