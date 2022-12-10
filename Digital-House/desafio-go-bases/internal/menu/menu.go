package menu

import (
	"fmt"
)

func Introduction() {
	fmt.Println("------------------------------------------------------")
	fmt.Printf("\n\t\tDIGITAL-HOUSE FLIGHTS\n\n")
}

func ShowMenu() {
	fmt.Println("------------------------------------------------------")
	fmt.Printf("Type your option:\n\n")
	fmt.Println("1\t-\tGet total tickets by location")
	fmt.Println("2\t-\tGet total of flights by period")
	fmt.Println("3\t-\tGet average of flights")
	fmt.Printf("0\t-\tExit program\n\n")
	fmt.Println("------------------------------------------------------")
}

func PeriodMenu() {
	fmt.Println("\nDawn")
	fmt.Println("Morning")
	fmt.Println("Afternoon")
	fmt.Printf("Night\n\n\n")
}

func GetUserInput() (inputedCommand int) {
	fmt.Scan(&inputedCommand)
	return inputedCommand
}
