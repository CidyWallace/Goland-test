// TESTE DE UNIDADE
package enderecos

import (
	"testing"
)

type cenarioDeTeste struct {
	endereceRecebido string
	retornoEsperado  string
}

func TestTipoDeEndereco(t *testing.T) {
	cenarioDeTeste := []cenarioDeTeste{
		{"Rua ABC", "Rua"},
		{"Avenida Paulista", "Avenida"},
		{"Rodovia dos imigrantes", "Rodovia"},
		{"Praça das rosas", "Tipo Inválido"},
		{"Estrada Qualquer", "Estrada"},
		{"RUA DOS BOBOS", "Rua"},
		{"AVENIDA REBOLÇAS", "Avenida"},
		{"", "Tipo Inválido"},
	}

	for _, cenario := range cenarioDeTeste {
		enderecoRecebido := TipoDeEndereco(cenario.endereceRecebido)
		if enderecoRecebido != cenario.retornoEsperado {
			t.Errorf("O tipo recebido é diferente do esperado! Esperava: %s e recebeu: %s",
				enderecoRecebido, cenario.retornoEsperado)
		}
	}

	// if tipoDeEnderecoRecebido != tipoDeEnderecoEsperado {
	// 	t.Errorf("O tipo recebido é diferente do esperado! Esperava: %s e recebeu: %s",
	// 		tipoDeEnderecoEsperado, tipoDeEnderecoRecebido)
	// }
}
