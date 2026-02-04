//go:build ignore

package main

import "fmt"

func main() {
	idade := map[string]int{}
	idade["erick"] = 28
	idade["vini"] = 4
	fmt.Println(idade)
	fmt.Println(idade["erick"])
	fmt.Println(idade["vini"])

	anoNasc := map[string]int{
		"erick": 1995,
		"vini":  2008,
	}
	fmt.Println(anoNasc)
	fmt.Println(anoNasc["erick"])
	fmt.Println(anoNasc["vini"])

	anoNasc["golangDoZero"] = 2024
	fmt.Println(anoNasc)
}

// 2 - Maps: Heterogêneos
// pode misturar tipos
// estrutura chave - valor
// [key] = value
// chave tem um tipo, e o valor pode ter outro tipo
// map[k]v -> k = chave, v = valor

// map[string]int
// {"erick": 28, "bento": 4}
// map[string]string
// {"erick": "gomes", "vinicios": "gomes"}
