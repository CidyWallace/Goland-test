package main

import "fmt"

var num int

func init() {
	fmt.Println("Executando a função Init")
	num = 10
}

func main() {
	fmt.Println("Executando a função main")
	fmt.Println(num)
}
