package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

/*
func main() {
	nome := "Douglas"
	idade := 24
	var versao float32 = 1.1

	fmt.Println("Olá, sr.", nome, "sua idade é", idade)
	fmt.Println("Este programa está na versão", versao)

	fmt.Println("O tipo da variável idade é", reflect.TypeOf(idade))
} */

const monitoramentos = 2
const delay = 5

func main() {

	/* 	nome, idade := devolveNomeEIdade()
			para pegar 1 valor só quando tem multiplos valores
			_, idade := devolveNomeEIdade
	   	fmt.Println(nome, idade) */

	/* 	if comando == 1 {
	   		fmt.Println("Monitorando...")
	   	} else if comando == 2 {
	   		fmt.Println("Exibindo Logs...")
	   	} else if comando == 0 {
	   		fmt.Println("Saindo do Programa...")
	   	} else {
	   		fmt.Println("Não conheço este comando")
	   	} */

	exibeIntroducao()
	for {
		exibeMenu()
		switch leComando() {
		case 1:
			iniciarMonitoramento()
		case 2:
			fmt.Println("Exibindo Logs...")
		case 0:
			fmt.Println("Saindo do Programa...")
			os.Exit(0)
		default:
			fmt.Println("Não conheço este comando")
			os.Exit(-1)
		}
	}

}

/* func devolveNomeEIdade() (string, int) {
	nome := "Eduardo"
	idade := 31
	return nome, idade
} */

func exibeIntroducao() {
	nome := "Eduardo"
	versao := 1.1
	fmt.Println("Olá, sr.", nome)
	fmt.Println("Este prgorama está na versão", versao)
}

func exibeMenu() {
	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir")
	fmt.Println("0- Sair do Programa")
}

func leComando() int {
	var comandoLido int
	fmt.Scan(&comandoLido)
	fmt.Println("O comando escolhido foi", comandoLido)

	return comandoLido
}

func iniciarMonitoramento() {
	fmt.Println("Monitorando...")
	sites := []string{"https://random-status-code.herokuapp.com/", "https://www.alura.com.br", "https://www.caelum.com.br"}

	for i := 0; i < monitoramentos; i++ {
		for i, site := range sites {
			fmt.Println("Testando site", i, ":", site)
			testaSite(site)
		}
		time.Sleep(delay * time.Second)
		fmt.Println("")

	}

	fmt.Println("")
}

func testaSite(site string) {
	resp, _ := http.Get(site)

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso!")
	} else {
		fmt.Println("Site:", site, "está com problema. Status Code:", resp.StatusCode)
	}
}

// func exibeNomes() {
// 	nomes := []string{"Douglas", "Eduardo", "Bernardo"}
// 	nomes = append(nomes, "Aparecida")
// 	fmt.Print(cap(nomes))
// 	fmt.Println(nomes)
// 	fmt.Println(len(nomes))
// }
