package main

import (
	"fmt"
)

/*
   Complete a seguinte função para que a mesma devolva todos os possíveis números de 4 dígitos,
   onde cada um seja menor ou igual a<maxDigit>, e a soma dos dígitos de cada número gerado seja 21
   Exemplo maxDigit=6: 3666, 4566...
*/

func main() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT
	var maxDigit int
	fmt.Println("Put the max Digit")
	fmt.Scan(&maxDigit)

	const sum = 21
	maxNumber := maxDigit * 1111

	numbers := make([]int, 0)

	for number := 1111; number <= maxNumber; number++ {
		firstNumber := number / 1000
		secondNumber := (number / 100) % 10
		thirdNumber := (number / 10) % 10
		fourthNumber := number % 10

		if firstNumber+secondNumber+thirdNumber+fourthNumber == sum {
			if firstNumber <= maxDigit && secondNumber <= maxDigit && thirdNumber <= maxDigit && fourthNumber <= maxDigit {
				numbers = append(numbers, number)
				fmt.Println(number)
			}
		}
	}
	if len(numbers) < 1 {
		fmt.Println("null")
	}
}
