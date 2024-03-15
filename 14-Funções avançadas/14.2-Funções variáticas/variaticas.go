package main

import "fmt"

//Esse tipo de parametro me permite mandar quantos valores quiser
func soma(numeros ...int) int {
	total := 0
	for _, valor := range numeros {
		total += valor
	}
	return total
}

func escrever(texto string, num ...int) {
	for _, valor := range num {
		fmt.Println(texto, valor)
	}
}

func main() {
	resultado := soma(1, 2, 3, 4, 5, 100, 250, 50, 35, 27)
	fmt.Println(resultado)

	escrever("Olá Mundo", 1, 2, 3, 4, 5, 6, 7, 8, 9)
}
