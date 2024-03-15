package main

import "fmt"

func funcao1() {
	fmt.Println("Executando função 1")
}

func funcao2() {
	fmt.Println("Executando função 2")
}

func alunoEstaAprovado(n1, n2 float32) bool {
	defer fmt.Println("Média calculada, o resultado será retornado")
	fmt.Println("Entrando na função para verificar se o aluno está aprovado")

	result := (n1 + n2) / 2

	if result >= 6 {
		return true
	}

	return false
}

func main() {
	//DEFER = ADIAR
	defer funcao1()
	funcao2()

	fmt.Println(alunoEstaAprovado(5, 6))
}
