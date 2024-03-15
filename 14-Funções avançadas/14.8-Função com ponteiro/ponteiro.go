package main

import "fmt"

func InverterNumero(numero int) int {
	return numero * -1
}

func InverterNumeroPonteiro(numero *int) {
	*numero = *numero * -1
}

func main() {
	numero := 20
	fmt.Println(InverterNumero(numero))
	fmt.Println(numero)

	fmt.Println("------------------")

	novoNumero := 10
	fmt.Println(novoNumero)
	InverterNumeroPonteiro(&novoNumero)
	fmt.Println(novoNumero)
}
