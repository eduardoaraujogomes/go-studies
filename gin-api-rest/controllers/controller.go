package controllers

import (
	"net/http"

	"github.com/eduardoaraujogomes/go-studies/database"
	"github.com/eduardoaraujogomes/go-studies/models"
	"github.com/gin-gonic/gin"
)

func GetAllStudents(c *gin.Context) {
	var students []models.Student
	database.DB.Find(&students)
	c.JSON(200, students)
}

func Greetings(c *gin.Context) {
	name := c.Params.ByName("name")
	c.JSON(200, gin.H{
		"API says:": "Hello " + name + ", is everthing ok?",
	})
}

func CreateStudent(c *gin.Context) {
	var student models.Student
	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	database.DB.Create(&student)
	c.JSON(http.StatusOK, student)
}

func GetStudentByID(c *gin.Context) {
	var student models.Student
	id := c.Params.ByName("id")
	database.DB.First(&student, id)
	if student.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not found": "Student not found",
		})
		return
	}
	c.JSON(http.StatusOK, student)

}

func DeleteStudent(c *gin.Context) {
	var student models.Student
	id := c.Params.ByName("id")
	database.DB.Delete(&student, id)
	c.JSON(http.StatusNoContent, gin.H{})
}

func UpdateStudent(c *gin.Context) {
	var student models.Student
	id := c.Params.ByName("id")
	database.DB.First(&student, id)
	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	database.DB.Model(&student).UpdateColumns(student)
	c.JSON(http.StatusOK, student)
}

func GetStudentByCPF(c *gin.Context) {
	var student models.Student
	cpf := c.Param("cpf")
	database.DB.Where(&models.Student{CPF: cpf}).First(&student)
	if student.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not found": "Student not found",
		})
		return
	}
	c.JSON(http.StatusOK, student)
}
