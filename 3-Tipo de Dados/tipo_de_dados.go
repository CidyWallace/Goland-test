package main

import (
	"errors"
	"fmt"
)

func main() {
	//NÚMEROS INTEIROS
	var numero int8 = 100
	fmt.Println("int8", numero)

	var numero2 uint = 10000
	fmt.Println("uint", numero2)

	//Int32 = Rune
	var numero3 rune = 3200
	fmt.Println("rune", numero3)

	//Uint8 = Byte
	var numero4 byte = 100
	fmt.Println("byte", numero4)
	//FIM NÚMEROS INTEIROS

	//NÚMEROS REAIS
	var numeroReal1 float32 = 123.45
	fmt.Println("float32", numeroReal1)

	var numeroReal2 float64 = 123000.45
	fmt.Println("float64", numeroReal2)

	numeroReal3 := 123000.45

	fmt.Println("float", numeroReal3)
	//FIM NÚMEROS REAIS

	//CARACTERES
	var str string = "cadeia de caracteres"
	fmt.Println("String explicito", str)

	str2 := "Segunda cadeia de caracteres"
	fmt.Println("String implícito", str2)

	char := 'p'
	fmt.Println(char)
	//FIM CARACTERES

	//Variáveis não inicializadas, possuem o valor 0

	var text string
	fmt.Println(text)

	var num int
	fmt.Println(num)

	var boleano1 bool
	fmt.Println(boleano1)

	var erro1 error = errors.New("Erro interno")
	fmt.Println(erro1)
}
