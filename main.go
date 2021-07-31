package main

import (
	"github.com/gin-gonic/gin"
)

type developer struct {
	Name  string ` json "name"`
	Emoji string `json "emoji"`
}

var developers = []developer{
	{Name: "John Doe", Emoji: ":)"},
	{Name: "Mary Doe", Emoji: ";)"},
}

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"devlopers": developers})
	})

	router.Run(":3000")
}
