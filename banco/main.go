package main

import (
	"banco/contas"
	"fmt"
)

func PagarBoleto(conta verificarConta, valorDoBoleto float64) {
	conta.Sacar(valorDoBoleto)
}

type verificarConta interface {
	Sacar(valor float64) string
}

func main() {
	contaDoDenis := contas.ContaPoupanca{}
	contaDoDenis.Depositar(100)
	contaDoDenis.Sacar(55)
	PagarBoleto(&contaDoDenis, 30)

	fmt.Println(contaDoDenis.ObterSaldo())

	contaDaLuisa := contas.ContaCorrente{}
	contaDaLuisa.Depositar(500)
	PagarBoleto(&contaDaLuisa, 200)

	fmt.Println(contaDaLuisa.ObterSaldo())
}

// contaExemplo := contas.ContaCorrente{}
// contaExemplo.Depositar(100)

// fmt.Println(contaExemplo.ObterSaldo())

// clienteBruno := clientes.Titular{"Bruno", "123.111.123.11", "Desenvolvedor"}
// contaDoBruno := contas.ContaCorrente{clienteBruno, 123, 121311, 100.0}

// fmt.Println(contaDoBruno)

// contaDaSilvia := contas.ContaCorrente{Titular: "Silvia", Saldo: 300}
// contaDoGustavo := contas.ContaCorrente{Titular: "Gustavo", Saldo: 100}

// status := contaDaSilvia.Transferir(200, &contaDoGustavo)

// fmt.Println(status)
// fmt.Println(contaDaSilvia)
// fmt.Println(contaDoGustavo)

// contaDaSilvia := ContaCorrente{}
// contaDaSilvia.titular = "Silvia"
// contaDaSilvia.saldo = 500

// fmt.Println(contaDaSilvia.saldo)

// fmt.Println(contaDaSilvia.Sacar(100))
// fmt.Println(contaDaSilvia.Sacar(-100))

// fmt.Println(contaDaSilvia.saldo)

// fmt.Println(contaDaSilvia.Depositar(100))
// fmt.Println(contaDaSilvia.Depositar(-100))

// status, valor := contaDaSilvia.Depositar(200)
// fmt.Println(status, valor)

// fmt.Println(contaDaSilvia.saldo)

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
