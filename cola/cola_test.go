package cola_test

import (
	TDACola "tdas/cola"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestColaVacia(t *testing.T) {
	t.Log("Hacemos pruebas con una cola sin elementos")
	cola := TDACola.CrearColaEnlazada[int]()
	require.NotNil(t, cola)
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.Desencolar() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.VerPrimero() })
	require.EqualValues(t, true, cola.EstaVacia())
}

func TestVolumen(t *testing.T) {
	t.Log("Hacemos pruebas encolando una gran cantidad de elementos, luego desapilando todos hasta que quede vacía")
	tam := 10000
	cola := TDACola.CrearColaEnlazada[int]()
	for i := 0; i < tam; i++ {
		cola.Encolar(i)
		require.EqualValues(t, 0, cola.VerPrimero())
	}
	for j := 0; j < tam; j++ {
		require.EqualValues(t, j, cola.Desencolar())
	}
	require.EqualValues(t, true, cola.EstaVacia())
}

func ColaGenerica[T any](t *testing.T, valor1, valor2, valor3 T) {
	t.Log("Hacemos pruebas encolando y desencolando elementos que pueden ser de cualquier tipo")
	cola := TDACola.CrearColaEnlazada[T]()
	var (
		elemento1 T = valor1
		elemento2 T = valor2
		elemento3 T = valor3
	)
	//Testeo comportamiento al encolar
	cola.Encolar(elemento1)
	require.EqualValues(t, false, cola.EstaVacia())
	require.EqualValues(t, elemento1, cola.VerPrimero())
	cola.Encolar(elemento2)
	require.EqualValues(t, elemento1, cola.VerPrimero())
	cola.Encolar(elemento3)
	require.EqualValues(t, elemento1, cola.VerPrimero())

	//Testeo comportamiento al desencolar
	require.EqualValues(t, elemento1, cola.Desencolar())
	require.EqualValues(t, elemento2, cola.Desencolar())
	require.EqualValues(t, false, cola.EstaVacia())
	require.EqualValues(t, elemento3, cola.Desencolar())

	//Testo comportamiento cuando está vacía
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.Desencolar() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.VerPrimero() })
	require.EqualValues(t, true, cola.EstaVacia())
}

func TestColaDeEnteros(t *testing.T) {
	t.Log("Hacemos pruebas con una cola de números enteros")
	var (
		valor1 = 3
		valor2 = 0
		valor3 = 5
	)
	ColaGenerica(t, valor1, valor2, valor3)
}

func TestColaDeCadenas(t *testing.T) {
	t.Log("Hacemos pruebas con una cola de cadenas")
	var (
		valor1 = "hola"
		valor2 = "como"
		valor3 = "estas"
	)
	ColaGenerica(t, valor1, valor2, valor3)
}

func TestColaDeBooleanos(t *testing.T) {
	t.Log("Hacemos pruebas con una cola de booleanos")
	var (
		valor1 = false
		valor2 = true
		valor3 = false
	)
	ColaGenerica(t, valor1, valor2, valor3)
}

func TestColaDeFlotantes(t *testing.T) {
	t.Log("Hacemos pruebas con una cola de números flotantes")
	var (
		valor1 = 12.5
		valor2 = 33.9
		valor3 = 51.2
	)
	ColaGenerica(t, valor1, valor2, valor3)
}
