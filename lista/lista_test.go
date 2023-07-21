package lista_test

import (
	TDALista "tdas/lista"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestListaVacia(t *testing.T) {
	t.Log("Hacemos pruebas con una lista vacía. Se valida que esté vacia, y que VerPrimero, BorrarPrimero, y VerUltimo sean invalidas")
	lista := TDALista.CrearListaEnlazada[int]()
	require.NotNil(t, lista)
	require.True(t, lista.EstaVacia())
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerPrimero() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerUltimo() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.BorrarPrimero() })
	require.EqualValues(t, 0, lista.Largo())
}

func TestVerPrimero(t *testing.T) {
	t.Log("Prueba de primeros. Se añaden y borran diferentes primeros, validando que sean correctos siempre")
	lista := TDALista.CrearListaEnlazada[int]()
	require.NotNil(t, lista)
	require.True(t, lista.EstaVacia())
	require.EqualValues(t, 0, lista.Largo())

	lista.InsertarPrimero(1)
	require.EqualValues(t, 1, lista.VerPrimero())
	require.False(t, lista.EstaVacia())
	require.EqualValues(t, 1, lista.Largo())

	lista.InsertarPrimero(2)
	require.EqualValues(t, 2, lista.VerPrimero())
	require.False(t, lista.EstaVacia())
	require.EqualValues(t, 2, lista.Largo())

	lista.InsertarPrimero(3)
	require.EqualValues(t, 3, lista.VerPrimero())
	require.False(t, lista.EstaVacia())
	require.EqualValues(t, 3, lista.Largo())

	require.EqualValues(t, 3, lista.BorrarPrimero())
	require.EqualValues(t, 2, lista.VerPrimero())
	require.False(t, lista.EstaVacia())
	require.EqualValues(t, 2, lista.Largo())

	require.EqualValues(t, 2, lista.BorrarPrimero())
	require.EqualValues(t, 1, lista.VerPrimero())
	require.False(t, lista.EstaVacia())
	require.EqualValues(t, 1, lista.Largo())

}

func TestVerUltimo(t *testing.T) {
	t.Log("Prueba de ultimos. Se añaden y borran diferentes ultimos, validando que sean correctos")
	lista := TDALista.CrearListaEnlazada[int]()
	require.NotNil(t, lista)
	require.True(t, lista.EstaVacia())
	require.EqualValues(t, 0, lista.Largo())

	lista.InsertarUltimo(1)
	require.EqualValues(t, 1, lista.VerUltimo())
	require.False(t, lista.EstaVacia())
	require.EqualValues(t, 1, lista.Largo())

	lista.InsertarUltimo(2)
	require.EqualValues(t, 2, lista.VerUltimo())
	require.False(t, lista.EstaVacia())
	require.EqualValues(t, 2, lista.Largo())

	lista.InsertarUltimo(3)
	require.EqualValues(t, 3, lista.VerUltimo())
	require.False(t, lista.EstaVacia())
	require.EqualValues(t, 3, lista.Largo())

	require.EqualValues(t, 1, lista.BorrarPrimero())
	require.EqualValues(t, 3, lista.VerUltimo())
	require.False(t, lista.EstaVacia())
	require.EqualValues(t, 2, lista.Largo())

	require.EqualValues(t, 2, lista.BorrarPrimero())
	require.EqualValues(t, 3, lista.VerUltimo())
	require.False(t, lista.EstaVacia())
	require.EqualValues(t, 1, lista.Largo())

}

func TestVolumenPrimeros(t *testing.T) {
	t.Log("Prueba de volumen. Se verifica que los primeros sean correctos")
	lista := TDALista.CrearListaEnlazada[int]()
	require.NotNil(t, lista)
	require.True(t, lista.EstaVacia())
	require.EqualValues(t, 0, lista.Largo())

	for i := 0; i <= 10000; i++ {
		lista.InsertarPrimero(i)
		require.False(t, lista.EstaVacia())
		require.EqualValues(t, i+1, lista.Largo())
		require.EqualValues(t, i, lista.VerPrimero())
	}

	for i := 10000; i >= 0; i-- {
		require.False(t, lista.EstaVacia())
		require.EqualValues(t, i, lista.VerPrimero())
		require.EqualValues(t, i, lista.BorrarPrimero())
		require.EqualValues(t, i, lista.Largo())
	}
	require.True(t, lista.EstaVacia())
	require.EqualValues(t, 0, lista.Largo())
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerPrimero() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerUltimo() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.BorrarPrimero() })
}

func TestVolumenUltimos(t *testing.T) {
	t.Log("Prueba de volumen. Se verifica que los ultimos sean correctos")
	lista := TDALista.CrearListaEnlazada[int]()
	require.NotNil(t, lista)
	require.True(t, lista.EstaVacia())
	require.EqualValues(t, 0, lista.Largo())

	for i := 0; i <= 10000; i++ {
		lista.InsertarUltimo(i)
		require.False(t, lista.EstaVacia())
		require.EqualValues(t, i+1, lista.Largo())
		require.EqualValues(t, i, lista.VerUltimo())
	}

	j := 0
	for i := 10000; i >= 0; i-- {
		require.False(t, lista.EstaVacia())
		require.EqualValues(t, 10000, lista.VerUltimo())
		require.EqualValues(t, j, lista.BorrarPrimero())
		require.EqualValues(t, i, lista.Largo())
		j++
	}
	require.True(t, lista.EstaVacia())
	require.EqualValues(t, 0, lista.Largo())
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerPrimero() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerUltimo() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.BorrarPrimero() })
}

func TestIntercalados(t *testing.T) {
	t.Log("Prueba con varios elementos. Añadiremos primeros y ultimos, y borraremos primeros y ultimos, aleatoriamente, validando que los primeros y ultimos sean correctos")
	lista := TDALista.CrearListaEnlazada[int]()
	require.NotNil(t, lista)
	require.True(t, lista.EstaVacia())
	require.EqualValues(t, 0, lista.Largo())

	lista.InsertarPrimero(1)
	require.EqualValues(t, lista.VerUltimo(), lista.VerPrimero())
	require.False(t, lista.EstaVacia())
	require.EqualValues(t, 1, lista.Largo())

	lista.InsertarUltimo(2)
	require.EqualValues(t, 2, lista.VerUltimo())
	require.EqualValues(t, 1, lista.VerPrimero())
	require.False(t, lista.EstaVacia())
	require.EqualValues(t, 2, lista.Largo())

	lista.InsertarPrimero(0)
	require.EqualValues(t, 2, lista.VerUltimo())
	require.EqualValues(t, 0, lista.VerPrimero())
	require.False(t, lista.EstaVacia())
	require.EqualValues(t, 3, lista.Largo())

	require.EqualValues(t, 0, lista.BorrarPrimero())
	require.EqualValues(t, 2, lista.VerUltimo())
	require.EqualValues(t, 1, lista.VerPrimero())
	require.False(t, lista.EstaVacia())
	require.EqualValues(t, 2, lista.Largo())

	lista.InsertarUltimo(3)
	require.EqualValues(t, 1, lista.BorrarPrimero())
	require.EqualValues(t, 3, lista.VerUltimo())
	require.EqualValues(t, 2, lista.VerPrimero())
	require.False(t, lista.EstaVacia())
	require.EqualValues(t, 2, lista.Largo())

	require.EqualValues(t, 2, lista.BorrarPrimero())
	require.EqualValues(t, lista.VerUltimo(), lista.VerPrimero())
	require.False(t, lista.EstaVacia())
	require.EqualValues(t, 1, lista.Largo())
}

func TestListaGenerica(t *testing.T) {
	t.Log("Haremos pruebas con una lista con elementos de diferentes tipos")
	lista := TDALista.CrearListaEnlazada[any]()
	require.NotNil(t, lista)
	require.True(t, lista.EstaVacia())
	require.EqualValues(t, 0, lista.Largo())

	lista.InsertarPrimero("Algoritmos")
	require.EqualValues(t, lista.VerUltimo(), lista.VerPrimero())
	require.False(t, lista.EstaVacia())
	require.EqualValues(t, 1, lista.Largo())

	lista.InsertarUltimo(2.0)
	require.EqualValues(t, 2.0, lista.VerUltimo())
	require.EqualValues(t, "Algoritmos", lista.VerPrimero())
	require.False(t, lista.EstaVacia())
	require.EqualValues(t, 2, lista.Largo())

	lista.InsertarUltimo(2023)
	require.EqualValues(t, 2023, lista.VerUltimo())
	require.EqualValues(t, "Algoritmos", lista.VerPrimero())
	require.False(t, lista.EstaVacia())
	require.EqualValues(t, 3, lista.Largo())

	require.EqualValues(t, "Algoritmos", lista.BorrarPrimero())
	require.EqualValues(t, 2023, lista.VerUltimo())
	require.EqualValues(t, 2.0, lista.VerPrimero())
	require.False(t, lista.EstaVacia())
	require.EqualValues(t, 2, lista.Largo())

	require.EqualValues(t, 2.0, lista.BorrarPrimero())
	require.EqualValues(t, lista.VerUltimo(), lista.VerPrimero())
	require.False(t, lista.EstaVacia())
	require.EqualValues(t, 1, lista.Largo())

}

func TestIterInsertarPrimero(t *testing.T) {
	t.Log("Con el iterador, valida que se pueda insertar un elemento sobre una lista vacia")
	lista := TDALista.CrearListaEnlazada[int]()
	iter := lista.Iterador()
	iter.Insertar(1)
	require.EqualValues(t, 1, lista.VerPrimero())
	require.EqualValues(t, 1, lista.VerUltimo())
	require.EqualValues(t, 1, lista.Largo())

	require.EqualValues(t, 1, lista.BorrarPrimero())
	require.True(t, lista.EstaVacia())

}

func TestInsertarMedio(t *testing.T) {
	t.Log("Con el iterador, valida que se pueda insertar un elemento en el medio")
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarUltimo(0)
	lista.InsertarUltimo(1)
	lista.InsertarUltimo(2)
	lista.InsertarUltimo(3)
	lista.InsertarUltimo(4)
	lista.InsertarUltimo(5)

	posicion := 0
	for iter := lista.Iterador(); iter.HaySiguiente(); {
		if posicion == 3 {
			iter.Insertar(10)
		} else {
			iter.Siguiente()
		}
		posicion++
	}

	require.EqualValues(t, 0, lista.VerPrimero())
	require.EqualValues(t, 0, lista.BorrarPrimero())
	require.False(t, lista.EstaVacia())
	require.EqualValues(t, 6, lista.Largo())

	require.EqualValues(t, 1, lista.VerPrimero())
	require.EqualValues(t, 1, lista.BorrarPrimero())
	require.False(t, lista.EstaVacia())
	require.EqualValues(t, 5, lista.Largo())

	require.EqualValues(t, 2, lista.VerPrimero())
	require.EqualValues(t, 2, lista.BorrarPrimero())
	require.False(t, lista.EstaVacia())
	require.EqualValues(t, 4, lista.Largo())

	require.EqualValues(t, 10, lista.VerPrimero())
	require.EqualValues(t, 10, lista.BorrarPrimero())
	require.False(t, lista.EstaVacia())
	require.EqualValues(t, 3, lista.Largo())

	require.EqualValues(t, 3, lista.VerPrimero())
	require.EqualValues(t, 3, lista.BorrarPrimero())
	require.False(t, lista.EstaVacia())
	require.EqualValues(t, 2, lista.Largo())

	require.EqualValues(t, 4, lista.VerPrimero())
	require.EqualValues(t, 4, lista.BorrarPrimero())
	require.False(t, lista.EstaVacia())
	require.EqualValues(t, 1, lista.Largo())

	require.EqualValues(t, 5, lista.VerPrimero())
	require.EqualValues(t, 5, lista.BorrarPrimero())
	require.True(t, lista.EstaVacia())
}

func TestInsertarFinal(t *testing.T) {
	t.Log("Con el iterador, valida que se pueda insertar un elemento al final de la lista")
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarUltimo(0)
	lista.InsertarUltimo(1)
	lista.InsertarUltimo(2)
	lista.InsertarUltimo(3)
	lista.InsertarUltimo(4)
	lista.InsertarUltimo(5)

	posicion := 0
	for iter := lista.Iterador(); posicion <= 6; {
		if posicion == 6 {
			iter.Insertar(10)
		} else {
			iter.Siguiente()
		}
		posicion++
	}

	require.EqualValues(t, 10, lista.VerUltimo())
}

func TestBorrarMedio(t *testing.T) {
	t.Log("Con el iterador, valida que se pueda borrar un elemento en el medio")
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarUltimo(0)
	lista.InsertarUltimo(1)
	lista.InsertarUltimo(2)
	lista.InsertarUltimo(3)
	lista.InsertarUltimo(4)
	lista.InsertarUltimo(5)

	posicion := 0
	for iter := lista.Iterador(); iter.HaySiguiente(); {
		if posicion == 3 {
			iter.Borrar()
		} else {
			iter.Siguiente()
		}
		posicion++
	}

	require.EqualValues(t, 0, lista.VerPrimero())
	require.EqualValues(t, 0, lista.BorrarPrimero())
	require.False(t, lista.EstaVacia())
	require.EqualValues(t, 4, lista.Largo())

	require.EqualValues(t, 1, lista.VerPrimero())
	require.EqualValues(t, 1, lista.BorrarPrimero())
	require.False(t, lista.EstaVacia())
	require.EqualValues(t, 3, lista.Largo())

	require.EqualValues(t, 2, lista.VerPrimero())
	require.EqualValues(t, 2, lista.BorrarPrimero())
	require.False(t, lista.EstaVacia())
	require.EqualValues(t, 2, lista.Largo())

	require.EqualValues(t, 4, lista.VerPrimero())
	require.EqualValues(t, 4, lista.BorrarPrimero())
	require.False(t, lista.EstaVacia())
	require.EqualValues(t, 1, lista.Largo())

	require.EqualValues(t, 5, lista.VerPrimero())
	require.EqualValues(t, 5, lista.BorrarPrimero())
	require.True(t, lista.EstaVacia())
}

func TestBorrarUnicoElemento(t *testing.T) {
	t.Log("Con el iterador, valida que se pueda borrar el primer elemento de una lista")
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarPrimero(1)

	iter := lista.Iterador()
	iter.Borrar()
	require.True(t, lista.EstaVacia())
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerPrimero() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerUltimo() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.BorrarPrimero() })
	require.EqualValues(t, 0, lista.Largo())
}

func TestBorrarPrimero(t *testing.T) {
	t.Log("Con el iterador, valida que se pueda borrar el primer elemento de una lista")
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarPrimero(1)
	lista.InsertarUltimo(2)
	lista.InsertarUltimo(3)

	iter := lista.Iterador()
	require.EqualValues(t, 1, iter.Borrar())
	require.EqualValues(t, 2, lista.Largo())
	require.EqualValues(t, 2, lista.VerPrimero())
}

func TestBorrarUltimo(t *testing.T) {
	t.Log("Con el iterador, valida que se pueda borrar el ultimo elemento de una lista, y este se actualice correctamente")
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarUltimo(0)
	lista.InsertarUltimo(1)
	lista.InsertarUltimo(2)
	lista.InsertarUltimo(3)
	lista.InsertarUltimo(4)
	lista.InsertarUltimo(5)

	posicion := 0
	for iter := lista.Iterador(); iter.HaySiguiente(); {
		if posicion == 5 {
			iter.Borrar()
		} else {
			iter.Siguiente()
		}
		posicion++
	}

	require.EqualValues(t, 4, lista.VerUltimo())
}

func TestIteradorInternoSinCondicionCorte(t *testing.T) {
	t.Log("Haremos pruebas con el iterador interno con una funcion sin condicion de corte")
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarUltimo(0)
	lista.InsertarUltimo(1)
	lista.InsertarUltimo(2)
	lista.InsertarUltimo(3)
	lista.InsertarUltimo(4)
	lista.InsertarUltimo(5)

	contador := 0

	contador_ptr := &contador
	lista.Iterar(func(v int) bool {
		*contador_ptr += v
		return true
	})
	require.EqualValues(t, 15, contador)
}

func TestIteradorInternoConCondicionCorte(t *testing.T) {
	t.Log("Haremos pruebas con el iterador interno con una funcion con condicion de corte")
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarUltimo(0)
	lista.InsertarUltimo(1)
	lista.InsertarUltimo(2)
	lista.InsertarUltimo(3)
	lista.InsertarUltimo(4)
	lista.InsertarUltimo(5)
	lista.InsertarUltimo(6)
	lista.InsertarUltimo(7)
	lista.InsertarUltimo(8)
	lista.InsertarUltimo(9)
	lista.InsertarUltimo(10)

	//Sumaremos el resultado de los primeros cinco
	contador := 0
	resultado := 0
	lista.Iterar(func(v int) bool {
		resultado += v
		contador++
		return contador <= 5
	})

	require.EqualValues(t, 15, resultado)
}
