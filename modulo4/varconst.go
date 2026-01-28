//go:build ignore

package main

import "fmt"

func main() {
	// Variáveis
	var nome string
	nome = "bento"
	fmt.Println(nome)

	nome = "erick"
	fmt.Println(nome)

	var idade int
	idade = 4
	fmt.Println(idade)

	var a = "erick"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b)
	fmt.Println(c)

	var d = true
	fmt.Println(d)

	f := "apple"
	fmt.Println(f)

	numero := 5
	fmt.Println(numero)

	// Constantes
	const idadeBento = 4
	fmt.Println(idadeBento)
}