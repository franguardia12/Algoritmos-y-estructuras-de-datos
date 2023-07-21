package pila

const (
	_CAPACIDAD_INICIAL           int = 5
	_FACTOR_REDIMENSION_AGRANDAR int = 2
	_FACTOR_REDIMENSION_ACHICAR  int = 2
)

/* Definición del struct pila proporcionado por la cátedra. */

type pilaDinamica[T any] struct {
	datos    []T
	cantidad int
}

func CrearPilaDinamica[T any]() Pila[T] {
	pila := new(pilaDinamica[T])
	pila.datos = make([]T, _CAPACIDAD_INICIAL)
	pila.cantidad = 0
	return pila
}

func (pila pilaDinamica[T]) validarPilaNoVacia() {
	if pila.EstaVacia() {
		panic("La pila esta vacia")
	}
}

func (pila pilaDinamica[T]) EstaVacia() bool {
	return pila.cantidad == 0
}

func (pila *pilaDinamica[T]) redimensionarPila(nueva_cap int) {
	nuevo := make([]T, nueva_cap)
	copy(nuevo, pila.datos)
	pila.datos = nuevo
}

func (pila *pilaDinamica[T]) Apilar(valor T) {
	if pila.cantidad == cap(pila.datos) {
		pila.redimensionarPila(cap(pila.datos) * _FACTOR_REDIMENSION_AGRANDAR)
	}
	pila.datos[pila.cantidad] = valor
	pila.cantidad++
}

func (pila *pilaDinamica[T]) Desapilar() T {
	pila.validarPilaNoVacia()
	if pila.cantidad*4 <= cap(pila.datos) {
		pila.redimensionarPila(cap(pila.datos) / _FACTOR_REDIMENSION_ACHICAR)
	}
	pila.cantidad--
	valor := pila.datos[pila.cantidad]
	return valor
}

func (pila *pilaDinamica[T]) VerTope() T {
	pila.validarPilaNoVacia()
	return pila.datos[pila.cantidad-1]
}
