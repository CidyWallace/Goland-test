package main

import (
	"fmt"
)

func somar(n1 int8, n2 int8) int8 {
	return n1 + n2
}

func calculosMatematicos(n1, n2 int8) (int8, int8) {
	soma := n1 + n2
	subt := n1 - n2
	return soma, subt
}

func main() {
	soma := somar(10, 20)
	fmt.Println(soma)

	var f = func(txt string) string {
		fmt.Println(txt)
		return txt
	}

	retorno := f("resultado da função 1")
	fmt.Println(retorno)

	resultadoSoma, _ := calculosMatematicos(10, 20)
	fmt.Println(resultadoSoma)
}
