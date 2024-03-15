package main

import "fmt"

func main() {
	fmt.Println("Estrutura de Controles")

	numero := -5

	if outroNumero := numero; outroNumero > 0 {
		fmt.Println("O número é maior que zero")
	} else {
		fmt.Println("O número é menor que zero")
	}
}
