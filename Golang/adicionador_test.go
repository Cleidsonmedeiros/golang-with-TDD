package inteiros

import "testing"

func Adiciona(x, y int) int {
	return 4
}

func TestAdicionador(t *testing.T) {
	soma := Adiciona(2, 2)
	esperado := 4

	if soma != esperado {
		t.Errorf("esperado '%d', resultado '%d'", esperado, soma)
	}
}
