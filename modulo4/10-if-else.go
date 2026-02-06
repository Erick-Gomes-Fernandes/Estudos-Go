//go:build ignore

package main

import "fmt"

// IF / ELSE
// SE / SENÃO

func main() {

	numero := 3
	condicao := 3
	// if (condição) { <ação> }
	if numero == condicao { // true
		fmt.Println("Valor é igual a", condicao)
	} else {
		fmt.Println("Valor não é igual a", condicao)
	}

	if numero == 1 {
		fmt.Println("Valor é igual a 1")
	} else if numero == 2 {
		fmt.Println("Valor é igual a 2")
	} else {
		fmt.Println("Valor é diferente de 1 e 2")
	}

	fmt.Println(1 + 1)
	fmt.Println(2 - 1)
	fmt.Println(2 * 2)
	fmt.Println(10 / 2)
	fmt.Println(10 % 2)

	if numero%2 == 0 {
		fmt.Println(numero, "é par")
	} else {
		fmt.Printf("%d é impar", numero)
	}
}

// Operações
// Soma: 1 + 1
// Subtração: 2 - 1
// Multiplicação: 2 * 2
// Divisão: 10 / 2
// Resto da divisão por x: 10 % 2
