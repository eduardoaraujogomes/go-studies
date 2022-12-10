package main

import (
	"fmt"
	"math"
)

type Matriz struct {
	values []float64
	height int
	width  int
}

func (s *Matriz) Set(values []float64) {
	if len(values) != s.height*s.width {
		fmt.Println("valores setados incorretamente, matriz inválida!")
		return
	}

	s.values = values
}

func (s *Matriz) Print() {
	for row := 0; row < s.height; row++ {
		startAt := s.width * row
		endAt := s.width*row + s.width
		fmt.Println(s.values[startAt:endAt])
	}
}

func (s *Matriz) Max() float64 {
	max := -math.MaxFloat64

	for _, v := range s.values {
		if v > max {
			max = v
		}
	}

	return max
}

func (s *Matriz) Quadratic() bool {
	if s.height == s.width && s.height != 0 {
		return true
	}
	return false
}

func main() {
	matriz := Matriz{
		height: 3,
		width:  3,
	}

	matriz.Set([]float64{1, 2, 3, 4, 5, 22, 7, 8, 9})
	matriz.Print()
	fmt.Println("Valor máximo:", matriz.Max())
	fmt.Println("É quadrática:", matriz.Quadratic())

}
