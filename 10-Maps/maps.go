package main

import "fmt"

func main() {
	fmt.Println("Maps")

	usuario := map[string]string{
		"nome":      "cidy",
		"sobrenome": "wallace",
	}

	fmt.Println(usuario["sobrenome"])

	usuario2 := map[string]map[string]string{

		"nome": {
			"primeiro": "João",
			"ultimo":   "Pedro",
		},
		"curso": {
			"nome":   "engenharia",
			"campus": "campus 1",
		},
	}

	delete(usuario2, "nome")
	fmt.Println(usuario2)

	usuario2["signo"] = map[string]string{
		"nome": "peixes",
	}

	usuario2["signo"] = map[string]string{
		"nome": "leão",
	}

	fmt.Println(usuario2)
}
