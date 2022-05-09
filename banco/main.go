package main

import "fmt"

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func main() {

	// var contaDoEduardo ContaCorrente = ContaCorrente{}

	contaGuilherme := ContaCorrente{
		titular: "Guilherme",
		saldo:   123.5,
	}

	contaDoEduardo := ContaCorrente{
		titular:       "Eduardo",
		numeroAgencia: 125,
		numeroConta:   121121,
		saldo:         122.4,
	}

	contaDaBruna := ContaCorrente{
		"Bruna",
		222,
		1212455,
		522.4,
	}

	fmt.Println(contaDoEduardo)
	fmt.Println(contaDaBruna)
	fmt.Println(contaGuilherme)

}
