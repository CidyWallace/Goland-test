package enderecos

import "strings"

//TipoDeEndereco verifica se um endereço tem um tipo válido e o retorna
func TipoDeEndereco(endereco string) string {
	tiposValidos := []string{"rua", "avenida", "estrada", "rodovia"}

	//strings.ToLower(): Faz uma string ficar minúscula
	primeiraPalavraDoEnderecao := strings.Split(strings.ToLower(endereco), " ")[0]

	enderecoValido := false

	for _, tipo := range tiposValidos {
		if primeiraPalavraDoEnderecao == tipo {
			enderecoValido = true
		}
	}

	if enderecoValido {
		return strings.Title(primeiraPalavraDoEnderecao)
	}
	return "Tipo Inválido"
}
