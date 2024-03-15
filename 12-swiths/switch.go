package main

import "fmt"

func diaDaSemana(num int8) string {
	switch num {
	case 1:
		return "domingo"
	case 2:
		return "segunda-feira"
	case 3:
		return "terça-feira"
	case 4:
		return "quarta-feira"
	case 5:
		return "quinta-feira"
	case 6:
		return "sexta-feira"
	case 7:
		return "sábado"
	default:
		return "Número Inválido"
	}
}

func main() {
	fmt.Println("Switchs")

	dia := diaDaSemana(20)
	fmt.Println(dia)
}
