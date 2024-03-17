package formas

import (
	"math"
	"testing"
)

func TestArea(t *testing.T) {
	//TDD - test Driven Development
	t.Run("Restangulo", func(t *testing.T) {
		r := Retangulo{12, 10}
		areaEsperada := float64(120)
		areaRecebida := r.Area()

		if areaEsperada != areaRecebida {
			//t.Errorf("Area recebida %f é diferente da área esperada %f", areaRecebida, areaEsperada)
			t.Fatalf("Area recebida %f é diferente da área esperada %f", areaRecebida, areaEsperada)
		}
	})

	t.Run("Circulo", func(t *testing.T) {
		c := Circulo{10}
		areaEsperada := float64(math.Pi * 100)
		areaRecebida := c.Area()

		if areaEsperada != areaRecebida {
			t.Fatalf("Area recebida %f é diferente da área esperada %f", areaRecebida, areaEsperada)
		}
	})
}
