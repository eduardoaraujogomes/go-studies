package main

import "fmt"

func main() {
	/* 	word := "palavra"

	   	fmt.Println("A palavra tem", len(word), "letras.")

	   	for _, letter := range word {
	   		fmt.Println(string(letter))
	   	} */

	number := 5
	months := []string{"Janeiro", "Fevereiro", "Março", "Abril", "Maio", "Junho", "Julho", "Agosto", "Setembro", "Outubro", "Novembro", "Dezembro"}

	fmt.Println(months[number-1])

}
