package enderecos_test

import (
	. "introducao-testes/enderecos"
	"testing"
)

type cenarioDeTeste struct {
	enderecoInserido string
	retornoEsperado  string
}

func TestTipoDeEndereco(t *testing.T) {
	// TestNomeDaFuncaoTestada(ponteiro do tipo T)

	cenariosDeTeste := []cenarioDeTeste{
		{"Rua Wenceslau Borini", "Rua"},
		{"Beco América", "Tipo Inválido"},
		{"Avenida Translada", "Avenida"},
		{"Avenida Paulista", "Avenida"},
		{"RODOVIA IMIGRANTES", "Rodovia"},
		{"estrada dos cabalos", "Estrada"},
		{"Wandinha", "Tipo Inválido"},
		{"", "Tipo Inválido"},
	}

	for _, cenario := range cenariosDeTeste {
		retornoRecebido := TipoDeEndereco(cenario.enderecoInserido)
		if retornoRecebido != cenario.retornoEsperado {
			t.Errorf("O tipo recebido %s é diferente do esperado %s",
				retornoRecebido,
				cenario.retornoEsperado,
			)
		}
	}

}
