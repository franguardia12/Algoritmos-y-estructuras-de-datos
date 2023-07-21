package diccionario_test

import (
	"fmt"
	"math/rand"
	TDADiccionario "tdas/diccionario"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

var TAMS_VOLUMEN = []int{12500, 25000, 50000, 100000, 200000, 400000}

func compararCadenas(cadena1, cadena2 string) int {
	if cadena1 > cadena2 {
		return 1
	} else if cadena1 < cadena2 {
		return -1
	} else {
		return 0
	}
}

func compararNumeros(a, b int) int {
	if a > b {
		return 1
	} else if a < b {
		return -1
	} else {
		return 0
	}
}

func TestDiccVacio(t *testing.T) {
	t.Log("Comprueba que Diccionario vacio no tiene claves")
	dic := TDADiccionario.CrearABB[string, string](compararCadenas)
	require.EqualValues(t, 0, dic.Cantidad())
	require.False(t, dic.Pertenece("A"))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Obtener("A") })
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Borrar("A") })
}

func TestDiccClaveDefault(t *testing.T) {
	t.Log("Prueba sobre un ABB vacío que si justo buscamos la clave que es el default del tipo de dato, " +
		"sigue sin existir")
	dic := TDADiccionario.CrearABB[string, string](compararCadenas)
	require.False(t, dic.Pertenece(""))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Obtener("") })
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Borrar("") })

	dicNum := TDADiccionario.CrearABB[int, string](compararNumeros)
	require.False(t, dicNum.Pertenece(0))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dicNum.Obtener(0) })
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dicNum.Borrar(0) })
}

func TestUnElemento(t *testing.T) {
	t.Log("Comprueba que Diccionario con un elemento tiene esa Clave, unicamente")
	dic := TDADiccionario.CrearABB[string, int](compararCadenas)
	dic.Guardar("A", 10)
	require.EqualValues(t, 1, dic.Cantidad())
	require.True(t, dic.Pertenece("A"))
	require.False(t, dic.Pertenece("B"))
	require.EqualValues(t, 10, dic.Obtener("A"))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Obtener("B") })
}

func TestDiccGuardar(t *testing.T) {
	t.Log("Guarda algunos pocos elementos en el diccionario, y se comprueba que en todo momento funciona acorde")
	clave1 := "Perro"
	clave2 := "Gato"
	clave3 := "Vaca"
	valor1 := "guau"
	valor2 := "miau"
	valor3 := "moo"
	claves := []string{clave1, clave2, clave3}
	valores := []string{valor1, valor2, valor3}

	dic := TDADiccionario.CrearABB[string, string](compararCadenas)
	require.False(t, dic.Pertenece(claves[0]))
	require.False(t, dic.Pertenece(claves[0]))
	dic.Guardar(claves[0], valores[0])
	require.EqualValues(t, 1, dic.Cantidad())
	require.True(t, dic.Pertenece(claves[0]))
	require.True(t, dic.Pertenece(claves[0]))
	require.EqualValues(t, valores[0], dic.Obtener(claves[0]))
	require.EqualValues(t, valores[0], dic.Obtener(claves[0]))

	require.False(t, dic.Pertenece(claves[1]))
	require.False(t, dic.Pertenece(claves[2]))
	dic.Guardar(claves[1], valores[1])
	require.True(t, dic.Pertenece(claves[0]))
	require.True(t, dic.Pertenece(claves[1]))
	require.EqualValues(t, 2, dic.Cantidad())
	require.EqualValues(t, valores[0], dic.Obtener(claves[0]))
	require.EqualValues(t, valores[1], dic.Obtener(claves[1]))

	require.False(t, dic.Pertenece(claves[2]))
	dic.Guardar(claves[2], valores[2])
	require.True(t, dic.Pertenece(claves[0]))
	require.True(t, dic.Pertenece(claves[1]))
	require.True(t, dic.Pertenece(claves[2]))
	require.EqualValues(t, 3, dic.Cantidad())
	require.EqualValues(t, valores[0], dic.Obtener(claves[0]))
	require.EqualValues(t, valores[1], dic.Obtener(claves[1]))
	require.EqualValues(t, valores[2], dic.Obtener(claves[2]))
}

func TestReemplazarDato(t *testing.T) {
	t.Log("Guarda un par de claves, y luego vuelve a guardar, buscando que el dato se haya reemplazado")
	clave := "Perro"
	clave2 := "Gato"
	dic := TDADiccionario.CrearABB[string, string](compararCadenas)
	dic.Guardar(clave, "guau")
	dic.Guardar(clave2, "miau")
	require.True(t, dic.Pertenece(clave))
	require.True(t, dic.Pertenece(clave2))
	require.EqualValues(t, "guau", dic.Obtener(clave))
	require.EqualValues(t, "miau", dic.Obtener(clave2))
	require.EqualValues(t, 2, dic.Cantidad())

	dic.Guardar(clave, "miu")
	dic.Guardar(clave2, "baubau")
	require.True(t, dic.Pertenece(clave))
	require.True(t, dic.Pertenece(clave2))
	require.EqualValues(t, 2, dic.Cantidad())
	require.EqualValues(t, "miu", dic.Obtener(clave))
	require.EqualValues(t, "baubau", dic.Obtener(clave2))
}

func TestReemplazarDatoHopscotch(t *testing.T) {
	t.Log("Guarda bastantes claves, y luego reemplaza sus datos. Luego valida que todos los datos sean " +
		"correctos. Para una implementación Hopscotch, detecta errores al hacer lugar o guardar elementos.")

	dic := TDADiccionario.CrearABB[int, int](compararNumeros)
	for i := 0; i < 500; i++ {
		dic.Guardar(i, i)
	}
	for i := 0; i < 500; i++ {
		dic.Guardar(i, 2*i)
	}
	ok := true
	for i := 0; i < 500 && ok; i++ {
		ok = dic.Obtener(i) == 2*i
	}
	require.True(t, ok, "Los elementos no fueron actualizados correctamente")
}

func TestDiccBorrar(t *testing.T) {
	t.Log("Guarda algunos pocos elementos en el diccionario, y se los borra, revisando que en todo momento " +
		"el diccionario se comporte de manera adecuada")
	clave1 := "Perro"
	clave2 := "Gato"
	clave3 := "Vaca"
	valor1 := "guau"
	valor2 := "miau"
	valor3 := "moo"
	claves := []string{clave1, clave2, clave3}
	valores := []string{valor1, valor2, valor3}
	dic := TDADiccionario.CrearABB[string, string](compararCadenas)

	require.False(t, dic.Pertenece(claves[0]))
	require.False(t, dic.Pertenece(claves[0]))
	dic.Guardar(claves[0], valores[0])
	dic.Guardar(claves[1], valores[1])
	dic.Guardar(claves[2], valores[2])

	require.True(t, dic.Pertenece(claves[2]))
	require.EqualValues(t, valores[2], dic.Borrar(claves[2]))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Borrar(claves[2]) })
	require.EqualValues(t, 2, dic.Cantidad())
	require.False(t, dic.Pertenece(claves[2]))

	require.True(t, dic.Pertenece(claves[0]))
	require.EqualValues(t, valores[0], dic.Borrar(claves[0]))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Borrar(claves[0]) })
	require.EqualValues(t, 1, dic.Cantidad())
	require.False(t, dic.Pertenece(claves[0]))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Obtener(claves[0]) })

	require.True(t, dic.Pertenece(claves[1]))
	require.EqualValues(t, valores[1], dic.Borrar(claves[1]))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Borrar(claves[1]) })
	require.EqualValues(t, 0, dic.Cantidad())
	require.False(t, dic.Pertenece(claves[1]))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Obtener(claves[1]) })
}

func TestReutilizacionDeBorrados(t *testing.T) {
	t.Log("Prueba que no haya problema reinsertando un elemento borrado")
	dic := TDADiccionario.CrearABB[string, string](compararCadenas)
	clave := "hola"
	dic.Guardar(clave, "mundo!")
	dic.Borrar(clave)
	require.EqualValues(t, 0, dic.Cantidad())
	require.False(t, dic.Pertenece(clave))
	dic.Guardar(clave, "mundooo!")
	require.True(t, dic.Pertenece(clave))
	require.EqualValues(t, 1, dic.Cantidad())
	require.EqualValues(t, "mundooo!", dic.Obtener(clave))
}

func TestClavesNumericas(t *testing.T) {
	t.Log("Valida que no solo funcione con strings")
	dic := TDADiccionario.CrearABB[int, string](compararNumeros)
	clave := 10
	valor := "Gatito"

	dic.Guardar(clave, valor)
	require.EqualValues(t, 1, dic.Cantidad())
	require.True(t, dic.Pertenece(clave))
	require.EqualValues(t, valor, dic.Obtener(clave))
	require.EqualValues(t, valor, dic.Borrar(clave))
	require.False(t, dic.Pertenece(clave))
}

func TestClaveVacia1(t *testing.T) {
	t.Log("Guardamos una clave vacía (i.e. \"\") y deberia funcionar sin problemas")
	dic := TDADiccionario.CrearABB[string, string](compararCadenas)
	clave := ""
	dic.Guardar(clave, clave)
	require.True(t, dic.Pertenece(clave))
	require.EqualValues(t, 1, dic.Cantidad())
	require.EqualValues(t, clave, dic.Obtener(clave))
}

func TestValorNulo1(t *testing.T) {
	t.Log("Probamos que el valor puede ser nil sin problemas")
	dic := TDADiccionario.CrearABB[string, *int](compararCadenas)
	clave := "Pez"
	dic.Guardar(clave, nil)
	require.True(t, dic.Pertenece(clave))
	require.EqualValues(t, 1, dic.Cantidad())
	require.EqualValues(t, (*int)(nil), dic.Obtener(clave))
	require.EqualValues(t, (*int)(nil), dic.Borrar(clave))
	require.False(t, dic.Pertenece(clave))
}

func mezclar(arr1 []string, arr2 []int) {
	//mezcla aleatoriamente dos arreglos
	for i := len(arr1) - 1; i > 0; i-- {
		j := generarNumeroAleatorio(len(arr1) - 1)
		arr1[i], arr1[j] = arr1[j], arr1[i]
		arr2[i], arr2[j] = arr2[j], arr2[i]
	}
}

func ejecutarPruebaVolumen1(b *testing.B, n int) {
	dic := TDADiccionario.CrearHash[string, int]()

	claves := make([]string, n)
	valores := make([]int, n)

	for i := 0; i < n; i++ {
		valores[i] = i
		claves[i] = fmt.Sprintf("%08d", i)
	}

	mezclar(claves, valores)

	for i := range claves {
		dic.Guardar(claves[i], valores[i])
	}

	require.EqualValues(b, n, dic.Cantidad(), "La cantidad de elementos es incorrecta")

	/* Verifica que devuelva los valores correctos */
	ok := true
	for i := 0; i < n; i++ {
		ok = dic.Pertenece(claves[i])
		if !ok {
			break
		}
		ok = dic.Obtener(claves[i]) == valores[i]
		if !ok {
			break
		}
	}

	require.True(b, ok, "Pertenece y Obtener con muchos elementos no funciona correctamente")
	require.EqualValues(b, n, dic.Cantidad(), "La cantidad de elementos es incorrecta")

	/* Verifica que borre y devuelva los valores correctos */
	for i := 0; i < n; i++ {
		ok = dic.Borrar(claves[i]) == valores[i]
		if !ok {
			break
		}
	}

	require.True(b, ok, "Borrar muchos elementos no funciona correctamente")
	require.EqualValues(b, 0, dic.Cantidad())
}

func BenchmarkDiccionario1(b *testing.B) {
	b.Log("Prueba de stress del Diccionario. Prueba guardando distinta cantidad de elementos (muy grandes), " +
		"ejecutando muchas veces las pruebas para generar un benchmark. Valida que la cantidad " +
		"sea la adecuada. Luego validamos que podemos obtener y ver si pertenece cada una de las claves geeneradas, " +
		"y que luego podemos borrar sin problemas")
	for _, n := range TAMS_VOLUMEN {
		b.Run(fmt.Sprintf("Prueba %d elementos", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				ejecutarPruebaVolumen1(b, n)
			}
		})
	}
}

func TestIterarDiccVacio(t *testing.T) {
	t.Log("Iterar sobre diccionario vacio es simplemente tenerlo al final")
	dic := TDADiccionario.CrearABB[string, int](compararCadenas)
	iter := dic.Iterador()
	require.False(t, iter.HaySiguiente())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })
}

func buscar(clave string, claves []string) int {
	for i, c := range claves {
		if c == clave {
			return i
		}
	}
	return -1
}

func TestDiccIterar(t *testing.T) {
	t.Log("Guardamos 3 valores en un Diccionario, e iteramos validando que las claves sean todas diferentes " +
		"pero pertenecientes al diccionario. Además los valores de VerActual y Siguiente van siendo correctos entre sí")
	clave1 := "Perro"
	clave2 := "Gato"
	clave3 := "Vaca"
	valor1 := "guau"
	valor2 := "miau"
	valor3 := "moo"
	claves := []string{clave1, clave2, clave3}
	valores := []string{valor1, valor2, valor3}
	dic := TDADiccionario.CrearABB[string, string](compararCadenas)
	dic.Guardar(claves[0], valores[0])
	dic.Guardar(claves[1], valores[1])
	dic.Guardar(claves[2], valores[2])
	iter := dic.Iterador()

	require.True(t, iter.HaySiguiente())
	primero, _ := iter.VerActual()
	require.NotEqualValues(t, -1, buscar(primero, claves))

	iter.Siguiente()
	segundo, segundo_valor := iter.VerActual()
	require.NotEqualValues(t, -1, buscar(segundo, claves))
	require.EqualValues(t, valores[buscar(segundo, claves)], segundo_valor)
	require.NotEqualValues(t, primero, segundo)
	require.True(t, iter.HaySiguiente())

	iter.Siguiente()
	require.True(t, iter.HaySiguiente())
	tercero, _ := iter.VerActual()
	require.NotEqualValues(t, -1, buscar(tercero, claves))
	require.NotEqualValues(t, primero, tercero)
	require.NotEqualValues(t, segundo, tercero)
	iter.Siguiente()

	require.False(t, iter.HaySiguiente())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })
}

func TestIterNoLlegaAlFinal(t *testing.T) {
	t.Log("Crea un iterador y no lo avanza. Luego crea otro iterador y lo avanza.")
	dic := TDADiccionario.CrearABB[string, string](compararCadenas)
	claves := []string{"B", "A", "C"}
	dic.Guardar(claves[0], "")
	dic.Guardar(claves[1], "")
	dic.Guardar(claves[2], "")

	dic.Iterador()
	iter2 := dic.Iterador()
	iter2.Siguiente()
	iter3 := dic.Iterador()
	primero, _ := iter3.VerActual()
	iter3.Siguiente()
	segundo, _ := iter3.VerActual()
	iter3.Siguiente()
	tercero, _ := iter3.VerActual()
	iter3.Siguiente()
	require.False(t, iter3.HaySiguiente())
	require.NotEqualValues(t, primero, segundo)
	require.NotEqualValues(t, tercero, segundo)
	require.NotEqualValues(t, primero, tercero)
	require.NotEqualValues(t, -1, buscar(primero, claves))
	require.NotEqualValues(t, -1, buscar(segundo, claves))
	require.NotEqualValues(t, -1, buscar(tercero, claves))
}

func TestIterInternoClaves(t *testing.T) {
	t.Log("Valida que todas las claves sean recorridas (y una única vez) con el iterador interno")
	clave1 := "Perro"
	clave2 := "Gato"
	clave3 := "Vaca"
	claves := []string{clave1, clave2, clave3}
	dic := TDADiccionario.CrearABB[string, *int](compararCadenas)
	dic.Guardar(claves[0], nil)
	dic.Guardar(claves[1], nil)
	dic.Guardar(claves[2], nil)

	cs := []string{"", "", ""}
	cantidad := 0
	cantPtr := &cantidad

	dic.Iterar(func(clave string, dato *int) bool {
		cs[cantidad] = clave
		*cantPtr = *cantPtr + 1
		return true
	})

	require.EqualValues(t, 3, cantidad)
	require.NotEqualValues(t, -1, buscar(cs[0], claves))
	require.NotEqualValues(t, -1, buscar(cs[1], claves))
	require.NotEqualValues(t, -1, buscar(cs[2], claves))
	require.NotEqualValues(t, cs[0], cs[1])
	require.NotEqualValues(t, cs[0], cs[2])
	require.NotEqualValues(t, cs[2], cs[1])
}

func TestIterInternoValores(t *testing.T) {
	t.Log("Valida que los datos sean recorridas correctamente (y una única vez) con el iterador interno")
	clave1 := "Gato"
	clave2 := "Perro"
	clave3 := "Vaca"
	clave4 := "Burrito"
	clave5 := "Hamster"

	dic := TDADiccionario.CrearABB[string, int](compararCadenas)
	dic.Guardar(clave1, 6)
	dic.Guardar(clave2, 2)
	dic.Guardar(clave3, 3)
	dic.Guardar(clave4, 4)
	dic.Guardar(clave5, 5)

	factorial := 1
	ptrFactorial := &factorial
	dic.Iterar(func(_ string, dato int) bool {
		*ptrFactorial *= dato
		return true
	})

	require.EqualValues(t, 720, factorial)
}

func generarNumeroAleatorio(N int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(N + 1)
}

func TestVolumenIterCorte(t *testing.T) {
	t.Log("Prueba de volumen de iterador interno, para validar que siempre que se indique que se corte" +
		" la iteración con la función visitar, se corte")

	dic := TDADiccionario.CrearABB[int, int](compararNumeros)

	/* Inserta 'n' parejas en el hash */
	for i := 0; i < 10000; i++ {
		numero := generarNumeroAleatorio(10000)
		dic.Guardar(numero, numero)
	}

	seguirEjecutando := true
	siguioEjecutandoCuandoNoDebia := false

	dic.Iterar(func(c int, v int) bool {
		if !seguirEjecutando {
			siguioEjecutandoCuandoNoDebia = true
			return false
		}
		if c%100 == 0 {
			seguirEjecutando = false
			return false
		}
		return true
	})

	require.False(t, seguirEjecutando, "Se tendría que haber encontrado un elemento que genere el corte")
	require.False(t, siguioEjecutandoCuandoNoDebia,
		"No debería haber seguido ejecutando si encontramos un elemento que hizo que la iteración corte")
}

func TestIteradorRango(t *testing.T) {
	t.Log("Prueba el iterador externo con ambos limites DESDE y HASTA definidos")
	dic := TDADiccionario.CrearABB[int, int](compararNumeros)
	dic.Guardar(7, 7)
	dic.Guardar(3, 3)
	dic.Guardar(8, 8)
	dic.Guardar(2, 2)
	dic.Guardar(5, 5)
	desde := 2
	hasta := 7
	iter := dic.IteradorRango(&desde, &hasta)

	require.True(t, iter.HaySiguiente())
	primero, _ := iter.VerActual()
	require.EqualValues(t, 2, primero)

	iter.Siguiente()
	require.True(t, iter.HaySiguiente())
	segundo, _ := iter.VerActual()
	require.EqualValues(t, 3, segundo)

	iter.Siguiente()
	require.True(t, iter.HaySiguiente())
	tercero, _ := iter.VerActual()
	require.EqualValues(t, 5, tercero)

	iter.Siguiente()
	require.True(t, iter.HaySiguiente())
	cuarto, _ := iter.VerActual()
	require.EqualValues(t, 7, cuarto)

}

func TestIteradorRangoConHasta(t *testing.T) {
	t.Log("Prueba el iterador externo solo con el limite HASTA definido")
	dic := TDADiccionario.CrearABB[int, int](compararNumeros)
	dic.Guardar(7, 7)
	dic.Guardar(3, 3)
	dic.Guardar(8, 8)
	dic.Guardar(2, 2)
	dic.Guardar(5, 5)
	hasta := 7
	iter := dic.IteradorRango(nil, &hasta)

	require.True(t, iter.HaySiguiente())
	primero, _ := iter.VerActual()
	require.EqualValues(t, 2, primero)

	iter.Siguiente()
	require.True(t, iter.HaySiguiente())
	segundo, _ := iter.VerActual()
	require.EqualValues(t, 3, segundo)

	iter.Siguiente()
	require.True(t, iter.HaySiguiente())
	tercero, _ := iter.VerActual()
	require.EqualValues(t, 5, tercero)

	iter.Siguiente()
	require.True(t, iter.HaySiguiente())
	cuarto, _ := iter.VerActual()
	require.EqualValues(t, 7, cuarto)

	iter.Siguiente()
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })

}

func TestIteradorRangoConDesde(t *testing.T) {
	t.Log("Prueba el iterador externo solo con el limite DESDE definido")
	dic := TDADiccionario.CrearABB[int, int](compararNumeros)
	dic.Guardar(7, 7)
	dic.Guardar(3, 3)
	dic.Guardar(8, 8)
	dic.Guardar(2, 2)
	dic.Guardar(5, 5)
	desde := 3
	iter := dic.IteradorRango(&desde, nil)

	require.True(t, iter.HaySiguiente())
	primero, _ := iter.VerActual()
	require.EqualValues(t, 3, primero)

	iter.Siguiente()
	require.True(t, iter.HaySiguiente())
	segundo, _ := iter.VerActual()
	require.EqualValues(t, 5, segundo)

	iter.Siguiente()
	require.True(t, iter.HaySiguiente())
	tercero, _ := iter.VerActual()
	require.EqualValues(t, 7, tercero)

	iter.Siguiente()
	require.True(t, iter.HaySiguiente())
	cuarto, _ := iter.VerActual()
	require.EqualValues(t, 8, cuarto)

	iter.Siguiente()
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })

}

func TestIteradorInternoRangos(t *testing.T) {
	t.Log("Prueba el iterador interno con un rango definido")
	dic := TDADiccionario.CrearABB[int, int](compararNumeros)
	dic.Guardar(7, 7)
	dic.Guardar(3, 3)
	dic.Guardar(8, 8)
	dic.Guardar(2, 2)
	dic.Guardar(5, 5)
	desde := 3
	hasta := 7
	factorial := 1
	ptrFactorial := &factorial
	dic.IterarRango(&desde, &hasta, func(_ int, dato int) bool {
		*ptrFactorial *= dato
		return true
	})

	require.EqualValues(t, 105, factorial)
}

func TestIteradorInternoRangosSinHasta(t *testing.T) {
	t.Log("Prueba el iterador interno con un rango sin desde (desde = nil)")
	dic := TDADiccionario.CrearABB[int, int](compararNumeros)
	dic.Guardar(7, 7)
	dic.Guardar(3, 3)
	dic.Guardar(8, 8)
	dic.Guardar(2, 2)
	dic.Guardar(5, 5)
	desde := 5
	factorial := 1
	ptrFactorial := &factorial
	dic.IterarRango(&desde, nil, func(_ int, dato int) bool {
		*ptrFactorial *= dato
		return true
	})

	require.EqualValues(t, 280, factorial)
}
