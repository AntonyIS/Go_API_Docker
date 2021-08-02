package routes

import (
	"github.com/gin-gonic/gin"
)

type developer struct {
	Name  string ` json "name"`
	Emoji string `json "emoji"`
}

var Developers = []developer{
	{Name: "John Doe", Emoji: ":)"},
	{Name: "Mary Doe", Emoji: ";)"},
}

var router = gin.Default()
