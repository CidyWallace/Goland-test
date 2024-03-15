package main

import "fmt"

func main() {
	canal := make(chan string, 2)

	canal <- "Olá mundo"
	canal <- "Programando em Go"

	menssagem := <-canal
	menssagem2 := <-canal

	fmt.Println(menssagem)
	fmt.Println(menssagem2)
}
