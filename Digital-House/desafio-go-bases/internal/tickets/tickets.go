package tickets

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Ticket struct {
	ID          int
	Name        string
	Email       string
	Destination string
	FlightHour  int
	Price       float64
}

func startCSVFile() ([]byte, error) {
	data, err := os.ReadFile("tickets.csv")

	if err != nil {
		return nil, err
	}

	return data, nil
}

func createTicket(name, email, destination string, id, flightHour int, price float64) Ticket {
	return Ticket{
		ID:          id,
		Name:        name,
		Email:       email,
		Destination: destination,
		FlightHour:  flightHour,
		Price:       price,
	}
}

func convertDataToTicket(data []string) (Ticket, error) {
	id, err := strconv.Atoi(data[0])

	if err != nil {
		return Ticket{}, fmt.Errorf("it was not possible to convert the ID")
	}

	name := data[1]

	email := data[2]

	destination := strings.ToUpper(data[3])

	flightHour, err := strconv.Atoi(strings.Split(data[4], ":")[0])

	if err != nil {
		return Ticket{}, fmt.Errorf("it was not possible to convert the time")
	}

	price, err := strconv.ParseFloat(data[5], 64)

	if err != nil {
		return Ticket{}, fmt.Errorf("it was not possible to convert the amount")
	}
	return createTicket(name, email, destination, id, flightHour, price), nil
}

func getAllTickets() (tickets []Ticket, err error) {

	startedFile, err := startCSVFile()
	if err != nil {
		return nil, err
	}

	dataSearatedByLineBreak := strings.Split(string(startedFile), "\n")

	for _, data := range dataSearatedByLineBreak {
		dataSeparatedByComma := strings.Split(data, ",")

		ticket, err := convertDataToTicket(dataSeparatedByComma)
		if err != nil {
			return nil, err
		}

		tickets = append(tickets, ticket)

	}
	return tickets, nil
}

func GetTotalTickets(destination string) (total int, err error) {
	tickets, err := getAllTickets()
	if err != nil {
		return 0, err
	}

	for _, ticket := range tickets {
		if ticket.Destination == destination {
			total++
		}
	}
	if total == 0 {
		return 0, fmt.Errorf("\nticket not found for \"%s\"", destination)
	}
	return total, nil
}

func getQuantityByHourInterval(startHour int, endHour int) (total int, err error) {
	tickets, err := getAllTickets()
	if err != nil {
		return 0, err
	}

	for _, ticket := range tickets {
		if ticket.FlightHour >= startHour && ticket.FlightHour <= endHour {
			total++
		}
	}
	return total, nil
}

func GetCountByPeriod(period string) (totalPeopleFlighting int, err error) {

	switch strings.ToUpper(period) {
	case "DAWN":
		totalPeopleFlighting, err = getQuantityByHourInterval(0, 6)
		if err != nil {
			return 0, err
		}

	case "MORNING":
		totalPeopleFlighting, err = getQuantityByHourInterval(7, 12)
		if err != nil {
			return 0, err
		}
	case "AFTERNOON":
		totalPeopleFlighting, err = getQuantityByHourInterval(13, 19)
		if err != nil {
			return 0, err
		}
	case "NIGHT":
		totalPeopleFlighting, err = getQuantityByHourInterval(20, 23)
		if err != nil {
			return 0, err
		}
	default:
		return 0, fmt.Errorf("\nperiod \"%s\" not found", period)
	}
	return totalPeopleFlighting, nil
}

func totalDistinctsCountries(tickets []Ticket) (total int) {
	distinct := make(map[string]bool)

	for _, ticket := range tickets {
		distinct[ticket.Destination] = true
	}
	total = len(distinct)

	return total
}

func AverageDestination() (average float64, err error) {
	tickets, err := getAllTickets()
	if err != nil {
		return 0, err
	}
	totalDestination := len(tickets)

	average = float64(totalDestination) / float64(totalDistinctsCountries(tickets))

	return average, nil
}
