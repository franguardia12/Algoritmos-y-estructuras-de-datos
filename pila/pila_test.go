package pila_test

import (
	TDAPila "tdas/pila"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPilaVacia(t *testing.T) {
	t.Log("Hacemos pruebas con pila de cantidad cero")
	pila := TDAPila.CrearPilaDinamica[int]()
	require.NotNil(t, pila)
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.Desapilar() })
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.VerTope() })
	require.EqualValues(t, true, pila.EstaVacia())
}

func TestVolumen(t *testing.T) {
	tam := 10000
	pila := TDAPila.CrearPilaDinamica[int]()
	for i := 0; i < tam; i++ {
		pila.Apilar(i)
		require.EqualValues(t, i, pila.VerTope())
	}
	for j := tam - 1; j >= 0; j-- {
		require.EqualValues(t, j, pila.Desapilar())
	}
	require.EqualValues(t, true, pila.EstaVacia())
}

func PilaGenerica[T any](t *testing.T, valor1, valor2, valor3 T) {
	t.Log("Hacemos pruebas con una pila apilando y desapilando elementos que pueden ser de cualquier tipo, también testeamos comportamiento de pila vacía")
	pila := TDAPila.CrearPilaDinamica[T]()
	var (
		elemento1 T = valor1
		elemento2 T = valor2
		elemento3 T = valor3
	)
	//Testeo comportamiento al apilar
	pila.Apilar(valor1)
	require.EqualValues(t, false, pila.EstaVacia())
	require.EqualValues(t, elemento1, pila.VerTope())
	pila.Apilar(valor2)
	require.EqualValues(t, elemento2, pila.VerTope())
	pila.Apilar(valor3)
	require.EqualValues(t, elemento3, pila.VerTope())

	//Testeo comportamiento al desapilar
	require.EqualValues(t, elemento3, pila.Desapilar())
	require.EqualValues(t, elemento2, pila.VerTope())
	require.EqualValues(t, elemento2, pila.Desapilar())
	require.EqualValues(t, false, pila.EstaVacia())
	require.EqualValues(t, elemento1, pila.VerTope())
	require.EqualValues(t, elemento1, pila.Desapilar())

	//Testeo comportamiento cuando está vacía
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.Desapilar() })
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.VerTope() })
	require.EqualValues(t, true, pila.EstaVacia())

}

func TestPilaDeEnteros(t *testing.T) {
	t.Log("Hacemos pruebas con una pila de enteros")
	var (
		valor1 int = 3
		valor2 int = 0
		valor3 int = 5
	)
	PilaGenerica(t, valor1, valor2, valor3)
}

func TestPilaDeCadenas(t *testing.T) {
	t.Log("Hacemos pruebas con una pila de cadenas")
	var (
		valor1 string = "hola"
		valor2 string = "como"
		valor3 string = "estas"
	)
	PilaGenerica(t, valor1, valor2, valor3)
}

func TestPilaDeBooleanos(t *testing.T) {
	t.Log("Hacemos pruebas con una pila de booleanos")
	var (
		valor1 bool = false
		valor2 bool = true
		valor3 bool = false
	)
	PilaGenerica(t, valor1, valor2, valor3)
}

func TestPilaDeFlotantes(t *testing.T) {
	t.Log("Hacemos pruebas con una pila de flotantes")
	var (
		valor1 float32 = 12.5
		valor2 float32 = 33.9
		valor3 float32 = 51.2
	)
	PilaGenerica(t, valor1, valor2, valor3)
}
