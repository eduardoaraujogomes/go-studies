package routes

import (
	"github.com/eduardoaraujogomes/go-studies/controllers"
	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()
	r.GET("/students", controllers.GetAllStudents)
	r.GET("/:name", controllers.Greetings)
	r.GET("/students/:id", controllers.GetStudentByID)
	r.POST("/students", controllers.CreateStudent)
	r.DELETE("/students/:id", controllers.DeleteStudent)
	r.PATCH("/students/:id", controllers.UpdateStudent)
	r.GET("/students/cpf/:cpf", controllers.GetStudentByCPF)
	r.Run()

}
