package main

import "fmt"

type Person struct {
	ID          int
	Name        string
	DateOfBirth string
}

type Employee struct {
	ID       int
	Position int
	Person
}

func (s *Employee) PrintEmployee() {
	fmt.Println(fmt.Sprintf("%d-%d-%s", s.ID, s.Position, s.Person.Name))
}

func main() {
	person := Employee{
		ID:       1,
		Position: 10,
		Person: Person{
			ID:          2,
			Name:        "Gabriel",
			DateOfBirth: "10/10/1995",
		},
	}

	person.PrintEmployee()
}
