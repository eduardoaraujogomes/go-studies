package main

import (
	"digital-house/internal/menu"
	"digital-house/internal/tickets"
	"fmt"
	"os"
	"strings"
)

func main() {
	menu.Introduction()
	for {
		menu.ShowMenu()
		switch menu.GetUserInput() {
		case 1:
			fmt.Print("Type the location: ")
			var location string
			fmt.Scan(&location)
			total, err := tickets.GetTotalTickets(strings.ToUpper(location))

			if err != nil {
				fmt.Println(err)
				continue
			}

			fmt.Printf("\n\nTotal of tickets by location: %d \n\n\n", total)
		case 2:
			fmt.Println("Type one of the periods options: ")
			menu.PeriodMenu()
			var period string
			fmt.Scan(&period)
			totalByPeriod, err := tickets.GetCountByPeriod(period)

			if err != nil {
				fmt.Println(err)
				continue
			}

			fmt.Printf("\n\nTotal of flights by period: %d \n\n\n", totalByPeriod)

		case 3:
			totalByCountry, err := tickets.AverageDestination()

			if err != nil {
				fmt.Println(err)
				continue
			}

			fmt.Printf("\n\nAverage of flights: %0.2f \n\n\n", totalByCountry)
		case 0:
			fmt.Println("\n\nExiting program...")
			os.Exit(0)

		default:
			fmt.Printf("\n\nInvalid input! Try again.\n\n\n")
		}

	}

}
