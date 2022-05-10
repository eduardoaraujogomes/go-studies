package main

import "fmt"

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func (c *ContaCorrente) Sacar(valorDoSaque float64) string {
	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.saldo

	if podeSacar {
		c.saldo -= valorDoSaque
		return "Saque realizado com sucesso!"
	}
	return "Saldo insuficiente"
}

func main() {
	contaDaSilvia := ContaCorrente{}
	contaDaSilvia.titular = "Silvia"
	contaDaSilvia.saldo = 500

	fmt.Println(contaDaSilvia.saldo)

	fmt.Println(contaDaSilvia.Sacar(100))
	fmt.Println(contaDaSilvia.Sacar(-100))

	fmt.Println(contaDaSilvia.saldo)

	// var contaDoEduardo ContaCorrente = ContaCorrente{}

	// contaGuilherme := ContaCorrente{
	// 	titular: "Guilherme",
	// 	saldo:   123.5,
	// }

	// contaDoEduardo := ContaCorrente{
	// 	titular:       "Eduardo",
	// 	numeroAgencia: 125,
	// 	numeroConta:   121121,
	// 	saldo:         122.4,
	// }

	// contaDaBruna := ContaCorrente{
	// 	"Bruna",
	// 	222,
	// 	1212455,
	// 	522.4,
	// }

	// // var contaDaCris *ContaCorrente
	// // contaDaCris = new(ContaCorrente)

	// contaDaCris := new(ContaCorrente)
	// contaDaCris.titular = "Cris"

	// fmt.Println(contaDoEduardo)
	// fmt.Println(contaDaBruna)
	// fmt.Println(contaGuilherme)
	// fmt.Println(*contaDaCris)

}
