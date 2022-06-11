package controllers

import (
	"github.com/eduardoaraujogomes/go-studies/models"
	"github.com/gin-gonic/gin"
)

func GetAllStudents(c *gin.Context) {
	c.JSON(200, models.Students)
}

func Greetings(c *gin.Context) {
	name := c.Params.ByName("name")
	c.JSON(200, gin.H{
		"API says:": "Hello " + name + ", is everthing ok?",
	})
}
