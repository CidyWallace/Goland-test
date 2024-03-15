package main

import "fmt"

type usuario struct {
	nome  string
	idade uint8
}

func (u usuario) salvar() {
	fmt.Printf("Salvando dados do usuário %s no banco de dados\n", u.nome)
}

func (u usuario) maiorDeIdade() bool {
	return u.idade >= 18
}

func (u *usuario) fazerAniversario() {
	u.idade++
}

func (u *usuario) setNome(n string) {
	u.nome = n
}

func main() {
	user := usuario{"Felipe", 20}
	user.salvar()

	user2 := usuario{"Marcos", 35}
	user2.salvar()

	maiorDeIdade := user2.maiorDeIdade()
	fmt.Println(maiorDeIdade)

	user2.fazerAniversario()
	fmt.Println(user2.idade)

	user2.setNome("Castro")
	user2.salvar()
	fmt.Println(user2.nome)
}
