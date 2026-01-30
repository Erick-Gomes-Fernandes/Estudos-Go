//go:build ignore

package main

import "fmt"

func main() {
	fmt.Println(soma(42, 13))
	somaDosValores := soma(44, 15)
	fmt.Println(somaDosValores)

	fmt.Println(subtracao(30, 12))
	subtracaoDosValores := subtracao(40, 14)
	fmt.Println(subtracaoDosValores)

	nome, sobrenome := PrintNome("ERICK", "GOMES")
	fmt.Println(nome)
	fmt.Println(sobrenome)
}

// Função começando com letra minúscula:
// Função é PRIVADA
// A Função só pode ser utilizada no próprio pacote
func soma(x int, y int) int {
	return x + y
}

func subtracao(x, y int) int {
	return x - y
}

// Função começando com letra maiúscula:
// Função é PÚBLICA
// A Função pode ser utilizada fora do próprio pacote
// Como utilizar ela fora do pacote main.PrintNome()
func PrintNome(nome, sobrenome string) (string, string) {
	return nome, sobrenome
}
