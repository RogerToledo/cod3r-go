package matematica

import "testing"

const erroPadrao = "Valor esperado %v, valor encontrado %v."

func TestMedia(t *testing.T) {
	valorEsperado := 8.62
	origem := []float64{6.0, 10.0, 9.5, 9.0}
	valor := Media(origem)

	if valorEsperado != valor {
		t.Errorf(erroPadrao, valorEsperado, valor)
	}
}

func TestSoma(t *testing.T) {
	valorEsperado := 18.0
	valores := make([]float64, 10)
	valores = append(valores, 5.0, 2.0, 10.0, 1.0)

	valor := Arit(valores, "+")

	if valorEsperado != valor {
		t.Errorf(erroPadrao, valorEsperado, valor)
	}
}

func TestSubtracao(t *testing.T) {
	valorEsperado := 1.0
	valores := 	[]float64{10.0, 5.0, 1.0, 3.0}

	valor := Arit(valores, "-")

	if valorEsperado != valor {
		t.Errorf(erroPadrao, valorEsperado, valor)
	}
}

func TestMultiplicacao(t *testing.T) {
	valorEsperado := 30.0
	valores := 	[]float64{1.0, 2.0, 5.0, 3.0}

	valor := Arit(valores, "*")

	if valorEsperado != valor {
		t.Errorf(erroPadrao, valorEsperado, valor)
	}
}

func TestDivisao(t *testing.T) {
	valorEsperado := 1.0
	valores := 	[]float64{30.0, 2.0, 5.0, 3.0}

	valor := Arit(valores, "/")

	if valorEsperado != valor {
		t.Errorf(erroPadrao, valorEsperado, valor)
	}
}

func TestArit(t *testing.T) {
	testes := []struct {
		operador string
		testCase []float64
		esperado float64

	}{
		{"+", []float64{4.0, 6.0, 10.0, 5.0}, 25},
		{"+", []float64{-10, 6.0, 10.0, 9}, 15},
		{"+", []float64{0.0, 1.0, 1.0, 0.0}, 2.0},
		{"-", []float64{0.0, 1.0, 1.0, 0.0}, -2.0},
		{"-", []float64{10, 5.0, 1.0, 0.0}, 4.0},
		{"-", []float64{0.0, 1.0, 1.0, 0.0}, -2.0},
		{"*", []float64{0.0, 1.0, 1.0, 0.0}, 0.0},
		{"*", []float64{5.0, 1.0, 2.0, 3.0}, 30.0},
		{"*", []float64{4.0, 3.0, 100.0, 2.0}, 2400.0},
		{"/", []float64{0.0, 1.0}, 0.0},
		{"/", []float64{2.0, 0.0}, 2.0},
		{"/", []float64{2400.0, 0.0, 100.0, 3.0}, 8.0},
	}

	for _, teste := range testes {
		t.Logf("Massa: %v", teste)
		atual := Arit(teste.testCase, teste.operador)

		if atual != teste.esperado {
			t.Errorf(erroPadrao, teste.esperado, atual)
		}
	}
}

func TestMedia2(t *testing.T) { 
	msg := "Média"
	testes := []struct {
		msg string
		testCase []float64
		esperado interface{}
	}{
		{msg, []float64{5.0}, 5.0},
		{msg, []float64{5.0, 5.0, 5.0, 5.0}, 5.0},
		{msg, []float64{5.0, 10.0, 2.0, 8.5}, 6.38},
		{msg, []float64{10.0, 0.0, 3.0, 8.5, 9.5}, 6.2},
		{msg, []float64{}, 0.0},
	}

	for _, teste := range testes {
		t.Logf(teste.msg)

		atual := Media(teste.testCase)

		if atual != teste.esperado {
			t.Errorf(erroPadrao, teste.esperado, atual)
		}
	}
}

