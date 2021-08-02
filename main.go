package main

import (
	"example.com/routes"
	"github.com/gin-gonic/gin"
)

var developers = routes.Developers

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"devlopers": developers})
	})

	router.Run(":3000")
}
