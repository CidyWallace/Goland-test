package main

import (
	"fmt"
	"introducao_teste/enderecos"
	"strings"
)

func main() {
	endereco := enderecos.TipoDeEndereco("Avenida Paulista")
	fmt.Println(strings.Title(endereco))
}
