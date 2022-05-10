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

func (c *ContaCorrente) Depositar(valorDoDeposito float64) (string, float64) {
	if valorDoDeposito > 0 {
		c.saldo += valorDoDeposito
		return "Deposito realizado com sucesso. Saldo atual: ", c.saldo
	}
	return "Valor do depósito menor que zero. Saldo atual: ", c.saldo
}

// função variádica
// func Somando(numeros ...int) int {
// 	resultadoDaSoma := 0
// 	for _, numero := range numeros {
// 		resultadoDaSoma += numero
// 	}
// 	return resultadoDaSoma
// }

func main() {

	contaDaSilvia := ContaCorrente{}
	contaDaSilvia.titular = "Silvia"
	contaDaSilvia.saldo = 500

	fmt.Println(contaDaSilvia.saldo)

	fmt.Println(contaDaSilvia.Sacar(100))
	fmt.Println(contaDaSilvia.Sacar(-100))

	fmt.Println(contaDaSilvia.saldo)

	fmt.Println(contaDaSilvia.Depositar(100))
	fmt.Println(contaDaSilvia.Depositar(-100))

	status, valor := contaDaSilvia.Depositar(200)
	fmt.Println(status, valor)

	fmt.Println(contaDaSilvia.saldo)

	//função variádica
	// numeros := []int{1, 2, 3, 4, 5, 6}
	// fmt.Println(Somando(numeros...))
	// fmt.Println(Somando(1))
	// fmt.Println(Somando(1, 1))
	// fmt.Println(Somando(1, 1, 1))
	// fmt.Println(Somando(1, 1, 2, 4))

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
