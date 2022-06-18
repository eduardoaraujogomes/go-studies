package main

import "fmt"

func main() {
	fmt.Printf("Outro programa em %s!", "Go")
}

// go version - mostra a versão
//godoc -http=:6060 - mostra a documentação de forma offline, na porta 6060
// go env - várias informações onde o go ta instalado, path e outras coisas
// go doc comandos.go - verifica se tem código suspeito
// go build comandos.go - builda o código, para ser executado direto
// go run comandos.go - compila e executa o código
// go get -u github... - instala um pacote
