package main

import "fmt"

func main() {
	// i := 0

	// for i < 10 {
	// 	i++
	// 	fmt.Println("Encrementado i")
	// 	time.Sleep(time.Second)
	// }

	// fmt.Println(i)

	// for j := 0; j < 10; j++ {
	// 	fmt.Println("Incrementado j ", j)
	// 	time.Sleep(time.Second)
	// }

	//nomes := [3]string{"Mateus", "Eduardo", "João"}

	//Caso não queira alguma da duas variaveis é so substituir por _:
	//	for _, nome := range nomes {
	//	fmt.Println(nome)
	//}

	// for indice, nome := range nomes {
	// 	fmt.Println(indice, nome)
	// }

	for indice, letra := range "PALAVRA" {
		fmt.Println(indice, string(letra))
	}

	usuario := map[string]string{
		"nome":      "Leonardo",
		"sobrenome": "Silva",
	}

	for chave, valor := range usuario {
		fmt.Println(chave, valor)
	}

}
