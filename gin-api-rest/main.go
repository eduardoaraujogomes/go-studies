package main

import (
	"github.com/eduardoaraujogomes/go-studies/models"
	"github.com/eduardoaraujogomes/go-studies/routes"
)

func main() {
	models.Students = []models.Student{
		{Name: "Eduardo", CPF: "019129121", RG: "121212121"},
		{Name: "Ana", CPF: "0133329121", RG: "41414141"},
		{Name: "João", CPF: "1111233", RG: "97744224"},
	}
	routes.HandleRequests()
}
