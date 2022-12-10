package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var total float64

	res, err := os.ReadFile("./Aula 8/data.csv")

	if err != nil {
		panic("error to read the file")
	}

	data := strings.Split(string(res), ";")

	for i := 0; i < len(data); i++ {
		column := strings.Split(data[i], ",")

		if i != 0 {
			price, err := strconv.ParseFloat(column[1], 64)

			if err != nil {
				fmt.Println("It was not possible to convert the price")
			}

			amount, err := strconv.ParseInt(column[2], 10, 64)

			if err != nil {
				fmt.Println("It was not possible to convert the amount")
			}

			totalProduct := price * float64(amount)
			total += totalProduct
		}

		for j := 0; j < len(column); j++ {
			fmt.Printf("%s\t\t", column[j])

			if j == len(column)-1 {
				fmt.Printf("\n")
			}
		}

		fmt.Printf("\nTotal\t\t%.2f\n", total)

	}
}
