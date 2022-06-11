package routes

import (
	"github.com/eduardoaraujogomes/go-studies/controllers"
	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()
	r.GET("/students", controllers.GetAllStudents)
	r.GET("/:name", controllers.Greetings)
	r.Run()
}
