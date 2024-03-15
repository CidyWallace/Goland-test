package main

import "fmt"

func main() {
	fmt.Println("Ponteiro")

	var num1 int = 10
	var num2 int = num1
	fmt.Println(num1, num2)

	num1++
	fmt.Println(num1, num2)

	//PONTEIRO É UM REFERÊNCIA DE MÉMORIA
	var num3 int
	var ponteiro *int

	num3 = 100
	ponteiro = &num3
	fmt.Println(num3, ponteiro)

	fmt.Println(num3, *ponteiro) //desreferenciação

	num3++
	fmt.Println(num3, *ponteiro) //desreferenciação

}
