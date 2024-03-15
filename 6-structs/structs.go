package main

import "fmt"

type usuario struct {
	nome     string
	idade    uint8
	endereco endereco
}

type endereco struct {
	rua    string
	numero uint8
}

func main() {
	fmt.Println("Arquivo Struct")

	var user usuario
	user.nome = "Wallace"
	user.idade = 20

	ender := endereco{"Rua Paulo Jóse dos Santos", 224}

	user2 := usuario{"Marcos", 35, ender}

	//Usado quando se tem apenas uma das informações necessárias
	user3 := usuario{nome: "Castro"}

	fmt.Println(user)
	fmt.Println(user2)
	fmt.Println(user3)

}
