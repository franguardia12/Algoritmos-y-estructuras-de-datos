package lista

type Lista[T any] interface {

	// EstaVacia devuelve verdadero si la cola no tiene elementos encolados, false en caso contrario.
	EstaVacia() bool

	//InsertarPrimero inserta el dato al inicio de la lista
	InsertarPrimero(T)

	// InsertarUltimo inserta el dato al final de la lista.
	InsertarUltimo(T)

	// BorrarPrimero borra el dato al inicio de la lista. Si la lista se encontraba vacía, entra el pánico con el
	// mensaje 'La lista esta vacia'.
	BorrarPrimero() T

	// VerPrimero devuelve el elemento al inicio de la lista (el primero). Si la lista se encontraba vacía, entra en
	// pánico con el mensaje 'La lista esta vacia'.
	VerPrimero() T

	// VerUltimo devuelve el elemento al final de la lista (el ultimo). Si la lista se encontraba vacía, entra em
	// pánico con el mensaje 'La lista esta vacia'.
	VerUltimo() T

	// Largo devuelve la cantidad de elementos de la lista
	Largo() int

	// Iterar aplica la funcion pasada por parametro a todos los elementos de la lista, hasta que no hayan más
	// elementos, o la función en cuestión devualva false.
	Iterar(visitar func(T) bool)

	//Devuelve una instancia de IteradorLista.
	Iterador() IteradorLista[T]
}

type IteradorLista[T any] interface {

	// Devuelve el elemento actual de la iteracion. Si se la invoca sobre un iterador que ya itero sobre
	// todos los elementos, entra en panico con el mensaje "El iterador termino de iterar"
	VerActual() T

	// Devuelve si hay algun elemento para ver en la posicion actual
	HaySiguiente() bool

	// Avanza una posicion en la iteracion. Si se lo invoca sobre un iterador que ya itero sobre
	// todos los elementos, entra en panico con el mensaje "El iterador termino de iterar"
	Siguiente()

	// Inserta un elemento en la posicion actual
	Insertar(T)

	// Borra el elemento de la posicion actual. Si se la invoca sobre un iterador que ya itero sobre
	// todos los elementos, entra en panico con el mensaje "El iterador termino de iterar"
	Borrar() T
}
