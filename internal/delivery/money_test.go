package delivery_test

import (
	"testing"

	"fabriciolfj.github/study/internal/delivery"
)

func somar(ds []delivery.Money) delivery.Money {
	var sum delivery.Money
	for _, d := range ds {
		sum += d
	}

	return sum
}

func TestSplitNuncaPerdeCentavo(t *testing.T) {
	for valor := delivery.Money(-500); valor <= 500; valor++ {
		for parts := 1; parts <= 13; parts++ {
			partes := valor.Split(parts)

			if len(partes) != parts {
				t.Fatalf("valor=%d parts=%d: len = %d", valor, parts, len(partes))
			}
			if soma := somar(partes); soma != valor {
				t.Fatalf("valor=%d parts=%d: soma = %d", valor, parts, soma)
			}
			// Nenhuma parte pode diferir de outra em mais de 1 centavo.
			menor, maior := partes[0], partes[0]
			for _, p := range partes {
				if p < menor {
					menor = p
				}
				if p > maior {
					maior = p
				}
			}
			if maior-menor > 1 {
				t.Fatalf("valor=%d parts=%d: diferenca %d entre partes", valor, parts, maior-menor)
			}
		}
	}
}
