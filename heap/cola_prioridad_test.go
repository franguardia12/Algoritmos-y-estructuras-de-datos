package cola_prioridad_test

import (
	"strings"
	TDAHeap "tdas/cola_prioridad"
	"testing"

	"github.com/stretchr/testify/require"
)

func compararNumeros(a, b int) int {
	return a - b
}

func TestHeapVacio(t *testing.T) {
	t.Log("Comprueba que un heap vacio se comporte como tal")
	heap := TDAHeap.CrearHeap(compararNumeros)
	require.True(t, heap.EstaVacia())
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.VerMax() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.Desencolar() })
	require.EqualValues(t, 0, heap.Cantidad())
}

func TestUnElemento(t *testing.T) {
	t.Log("Encola un elemento y verifica que el heap se comporte correctamente")
	heap := TDAHeap.CrearHeap(compararNumeros)
	require.True(t, heap.EstaVacia())
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.VerMax() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.Desencolar() })
	require.EqualValues(t, 0, heap.Cantidad())

	heap.Encolar(10)
	require.False(t, heap.EstaVacia())
	require.EqualValues(t, 10, heap.VerMax())
	require.EqualValues(t, 1, heap.Cantidad())
}

func TestVariosElementos(t *testing.T) {
	t.Log("Encola varios elementos y verifica que el heap se comporte correctamente")
	heap := TDAHeap.CrearHeap(compararNumeros)
	require.True(t, heap.EstaVacia())
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.VerMax() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.Desencolar() })
	require.EqualValues(t, 0, heap.Cantidad())

	heap.Encolar(10)
	require.False(t, heap.EstaVacia())
	require.EqualValues(t, 10, heap.VerMax())
	require.EqualValues(t, 1, heap.Cantidad())

	heap.Encolar(5)
	require.False(t, heap.EstaVacia())
	require.EqualValues(t, 10, heap.VerMax())
	require.EqualValues(t, 2, heap.Cantidad())

	heap.Encolar(20)
	require.False(t, heap.EstaVacia())
	require.EqualValues(t, 20, heap.VerMax())
	require.EqualValues(t, 3, heap.Cantidad())
}

func TestDesencolar(t *testing.T) {
	t.Log("Encola y desencola varios elementos hasta que el heap esté vacio, comprobando el maximo cada vez. Comprueba que una vez vacío se comporte como tal-")

	heap := TDAHeap.CrearHeap(compararNumeros)
	require.True(t, heap.EstaVacia())
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.VerMax() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.Desencolar() })
	require.EqualValues(t, 0, heap.Cantidad())

	heap.Encolar(10)
	require.False(t, heap.EstaVacia())
	require.EqualValues(t, 10, heap.VerMax())
	require.EqualValues(t, 1, heap.Cantidad())

	heap.Encolar(5)
	require.False(t, heap.EstaVacia())
	require.EqualValues(t, 10, heap.VerMax())
	require.EqualValues(t, 2, heap.Cantidad())

	heap.Encolar(20)
	require.False(t, heap.EstaVacia())
	require.EqualValues(t, 20, heap.VerMax())
	require.EqualValues(t, 3, heap.Cantidad())

	heap.Desencolar()
	require.False(t, heap.EstaVacia())
	require.EqualValues(t, 10, heap.VerMax())
	require.EqualValues(t, 2, heap.Cantidad())

	heap.Desencolar()
	require.False(t, heap.EstaVacia())
	require.EqualValues(t, 5, heap.VerMax())
	require.EqualValues(t, 1, heap.Cantidad())

	heap.Desencolar()
	require.True(t, heap.EstaVacia())
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.VerMax() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.Desencolar() })
	require.EqualValues(t, 0, heap.Cantidad())
}

func TestStrings(t *testing.T) {
	t.Log("Hace pruebas para un heap de cadenas")
	heap := TDAHeap.CrearHeap(strings.Compare)
	require.True(t, heap.EstaVacia())
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.VerMax() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.Desencolar() })
	require.EqualValues(t, 0, heap.Cantidad())

	heap.Encolar("Hola")
	require.False(t, heap.EstaVacia())
	require.EqualValues(t, "Hola", heap.VerMax())
	require.EqualValues(t, 1, heap.Cantidad())

	heap.Encolar("Algoritmos")
	require.False(t, heap.EstaVacia())
	require.EqualValues(t, "Hola", heap.VerMax())
	require.EqualValues(t, 2, heap.Cantidad())

	heap.Encolar("Y Programacion")
	require.False(t, heap.EstaVacia())
	require.EqualValues(t, "Y Programacion", heap.VerMax())
	require.EqualValues(t, 3, heap.Cantidad())

	heap.Desencolar()
	require.False(t, heap.EstaVacia())
	require.EqualValues(t, "Hola", heap.VerMax())
	require.EqualValues(t, 2, heap.Cantidad())

	heap.Desencolar()
	require.False(t, heap.EstaVacia())
	require.EqualValues(t, "Algoritmos", heap.VerMax())
	require.EqualValues(t, 1, heap.Cantidad())

	heap.Desencolar()
	require.True(t, heap.EstaVacia())
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.VerMax() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.Desencolar() })
	require.EqualValues(t, 0, heap.Cantidad())
}

func TestVolumen(t *testing.T) {
	t.Log("Realiza pruebas de volumen para un heap. Encola y desencola una cantidad de elementos grande")
	heap := TDAHeap.CrearHeap(compararNumeros)
	require.True(t, heap.EstaVacia())
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.VerMax() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.Desencolar() })
	require.EqualValues(t, 0, heap.Cantidad())

	for i := 0; i < 10000; i++ {
		require.EqualValues(t, i, heap.Cantidad())
		heap.Encolar(i)
		require.False(t, heap.EstaVacia())
		require.EqualValues(t, i, heap.VerMax())
		require.EqualValues(t, i+1, heap.Cantidad())
	}

	for i := 10000; i > 0; i-- {
		require.EqualValues(t, i, heap.Cantidad())
		require.EqualValues(t, i-1, heap.VerMax())
		require.False(t, heap.EstaVacia())
		heap.Desencolar()
	}

	require.True(t, heap.EstaVacia())
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.VerMax() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.Desencolar() })
	require.EqualValues(t, 0, heap.Cantidad())
}

func TestCrearHeapArreglo(t *testing.T) {
	t.Log("Comprueba que se cree un heap a partir de un arreglo y se comporte correctamente")
	arreglo := []int{3, 10, 5, 6, 9}
	heap := TDAHeap.CrearHeapArr(arreglo, compararNumeros)

	require.False(t, heap.EstaVacia())
	require.EqualValues(t, 10, heap.VerMax())
	require.EqualValues(t, 5, heap.Cantidad())

	heap.Desencolar()
	require.False(t, heap.EstaVacia())
	require.EqualValues(t, 9, heap.VerMax())
	require.EqualValues(t, 4, heap.Cantidad())

	heap.Desencolar()
	require.False(t, heap.EstaVacia())
	require.EqualValues(t, 6, heap.VerMax())
	require.EqualValues(t, 3, heap.Cantidad())

	heap.Desencolar()
	require.False(t, heap.EstaVacia())
	require.EqualValues(t, 5, heap.VerMax())
	require.EqualValues(t, 2, heap.Cantidad())

	heap.Desencolar()
	require.False(t, heap.EstaVacia())
	require.EqualValues(t, 3, heap.VerMax())
	require.EqualValues(t, 1, heap.Cantidad())

	heap.Desencolar()
	require.True(t, heap.EstaVacia())
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.VerMax() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.Desencolar() })
	require.EqualValues(t, 0, heap.Cantidad())
}

func TestCrearHeapArregloVacio(t *testing.T) {
	t.Log("Crea un heap a partir de un arreglo vacio")
	arr := []int{}
	heap := TDAHeap.CrearHeapArr(arr, compararNumeros)
	require.True(t, heap.EstaVacia())
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.VerMax() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.Desencolar() })
	require.EqualValues(t, 0, heap.Cantidad())

	heap.Encolar(10)
	require.False(t, heap.EstaVacia())
	require.EqualValues(t, 10, heap.VerMax())
	require.EqualValues(t, 1, heap.Cantidad())
}

func TestHeapSort(t *testing.T) {
	t.Log("Comprueba que el HeapSort ordene un arreglo correctamente")
	arreglo := []int{100, -10, 7, 3, 90, 1}
	esperado := []int{-10, 1, 3, 7, 90, 100}
	TDAHeap.HeapSort(arreglo, compararNumeros)
	require.EqualValues(t, esperado, arreglo)
}
