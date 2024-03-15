package main

import "fmt"

func main() {
	//variável com tipo explicito
	var variavel1 string = "variável 1"
	//variável com tipo implicito
	variavel2 := "variável 2"

	fmt.Println(variavel1)
	fmt.Println(variavel2)

	var (
		variavel3 string = "variavel 3"
		variavel4 string = "variavel 4"
	)

	fmt.Println(variavel3, variavel4)

	variavel5, variavel6 := "variavel 5", "variavel 6"

	fmt.Println(variavel5, variavel6)

	const constante1 string = "Constante 1"

	fmt.Println(constante1)

	//invertendo valores entre duas variáveis
	variavel5, variavel6 = variavel6, variavel5

	fmt.Println(variavel5, variavel6)
}
