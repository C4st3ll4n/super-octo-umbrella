package main

import "testing"

func TestSoma(t *testing.T) {
	resultado := Soma(1, 2)
	if resultado != 3 {
		t.Errorf("Soma(1, 2) = %d; esperado 3", resultado)
	}
}

func TestSubtrai(t *testing.T) {
	resultado := Subtrai(1, 2)
	if resultado != -1 {
		t.Errorf("Subtrai(1, 2) = %d; esperado -1", resultado)
	}
}

func TestMultiplica(t *testing.T) {
	resultado := Multiplica(1, 2)
	if resultado != 2 {
		t.Errorf("Multiplica(1, 2) = %d; esperado 2", resultado)
	}
}
