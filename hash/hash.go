package diccionario

import (
	"fmt"
	TDALista "tdas/lista"
)

const (
	CAPACIDAD_INICIAL = 5
	INICIO_BUSCAR_POS = -1
)

type parClaveValor[K comparable, V any] struct {
	clave K
	dato  V
}

type hashAbierto[K comparable, V any] struct {
	tabla    []TDALista.Lista[*parClaveValor[K, V]]
	tam      int
	cantidad int
}

type iterDiccionario[K comparable, V any] struct {
	hash          *hashAbierto[K, V]
	actual        *parClaveValor[K, V]
	iteradorLista TDALista.IteradorLista[*parClaveValor[K, V]]
	posicion      int
}

func hashing_function(data []byte) uint32 {
	const offset32 uint32 = 2166136261
	const prime32 = 16777619
	hash := offset32
	for _, d := range data {
		hash ^= uint32(d)
		hash *= prime32
	}
	return hash
}

func convertirABytes[K comparable](clave K) []byte {
	return []byte(fmt.Sprintf("%v", clave))
}

func (hash *hashAbierto[K, V]) buscarClave(clave K) (TDALista.IteradorLista[*parClaveValor[K, V]], error) {
	pos := int(hashing_function(convertirABytes(clave))) % hash.tam
	iter := hash.tabla[pos].Iterador()
	if hash.tabla[pos].EstaVacia() {
		return iter, fmt.Errorf("la clave no pertenece al diccionario")
	}
	var actual *parClaveValor[K, V]
	for iter.HaySiguiente() {
		actual = iter.VerActual()
		if actual.clave == clave {
			return iter, nil
		}
		iter.Siguiente()
	}
	return iter, fmt.Errorf("la clave no pertenece al diccionario")
}

func (hash *hashAbierto[K, V]) redimensionar(nuevo_tam int) {
	hash.tam = nuevo_tam
	nueva_tabla := make([]TDALista.Lista[*parClaveValor[K, V]], nuevo_tam)
	for i := 0; i < nuevo_tam; i++ {
		lista := TDALista.CrearListaEnlazada[*parClaveValor[K, V]]()
		nueva_tabla[i] = lista
	}
	for _, lista := range hash.tabla {
		iter := lista.Iterador()
		for iter.HaySiguiente() {
			parClaveValor := iter.VerActual()
			clave := parClaveValor.clave
			pos := int(hashing_function(convertirABytes(clave))) % hash.tam
			nueva_tabla[pos].InsertarUltimo(parClaveValor)
			iter.Siguiente()
		}
	}
	hash.tabla = nueva_tabla
}

func CrearHash[K comparable, V any]() Diccionario[K, V] {
	hash := new(hashAbierto[K, V])
	hash.tabla = make([]TDALista.Lista[*parClaveValor[K, V]], CAPACIDAD_INICIAL)
	for i := 0; i < CAPACIDAD_INICIAL; i++ {
		hash.tabla[i] = TDALista.CrearListaEnlazada[*parClaveValor[K, V]]()
	}
	hash.tam = CAPACIDAD_INICIAL
	return hash
}

func (hash hashAbierto[K, V]) Cantidad() int {
	return hash.cantidad
}

func (hash hashAbierto[K, V]) Pertenece(clave K) bool {
	_, err := hash.buscarClave(clave)
	return err == nil
}

func (hash hashAbierto[K, V]) Obtener(clave K) V {
	var valor V
	iter, err := hash.buscarClave(clave)
	if err != nil {
		panic("La clave no pertenece al diccionario")
	}
	valor = iter.VerActual().dato
	return valor
}

func (hash *hashAbierto[K, V]) Guardar(clave K, dato V) {
	factorCarga := hash.cantidad / hash.tam
	if factorCarga > 3 {
		hash.redimensionar(hash.tam * 2)
	}
	var par parClaveValor[K, V]
	par.clave = clave
	par.dato = dato
	pos := int(hashing_function(convertirABytes(clave))) % hash.tam

	iter, err := hash.buscarClave(clave)
	if err != nil {
		hash.tabla[pos].InsertarUltimo(&par)
		hash.cantidad++
	} else {
		guardado := iter.VerActual()
		guardado.dato = dato
	}
}

func (hash *hashAbierto[K, V]) Borrar(clave K) V {
	factorCarga := hash.cantidad / hash.tam
	if factorCarga < 2 && hash.tam > CAPACIDAD_INICIAL {
		hash.redimensionar(hash.tam / 2)
	}

	iter, err := hash.buscarClave(clave)
	if err != nil {
		panic("La clave no pertenece al diccionario")
	}
	borrado := iter.Borrar()
	hash.cantidad--
	return borrado.dato
}

func (hash hashAbierto[K, V]) proximaPosicionNoVacia(pos int) int {
	if pos > hash.tam {
		return pos
	}
	posicion := pos
	for i := range hash.tabla {
		if i <= pos {
			continue
		}
		if !hash.tabla[i].EstaVacia() {
			posicion = i
			break
		}
	}
	return posicion
}

func (hash *hashAbierto[K, V]) Iterador() IterDiccionario[K, V] {
	iter := new(iterDiccionario[K, V])
	iter.hash = hash
	posicion := hash.proximaPosicionNoVacia(INICIO_BUSCAR_POS)
	if posicion == INICIO_BUSCAR_POS {
		return iter
	}
	iter.actual = iter.hash.tabla[posicion].VerPrimero()
	iter.iteradorLista = iter.hash.tabla[posicion].Iterador()
	iter.posicion = posicion
	return iter
}
func (iter *iterDiccionario[K, V]) HaySiguiente() bool {
	return iter.actual != nil
}

func (iter *iterDiccionario[K, V]) Siguiente() {
	if !iter.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	iter.iteradorLista.Siguiente()
	var nueva_pos int
	if !iter.iteradorLista.HaySiguiente() {
		nueva_pos = iter.hash.proximaPosicionNoVacia(iter.posicion)
		if nueva_pos == iter.posicion {
			iter.actual = nil
		} else {
			iter.posicion = nueva_pos
			iter.iteradorLista = iter.hash.tabla[iter.posicion].Iterador()
			iter.actual = iter.hash.tabla[iter.posicion].VerPrimero()
		}
	} else {
		iter.actual = iter.iteradorLista.VerActual()
	}
}

func (iter *iterDiccionario[K, V]) VerActual() (K, V) {
	if !iter.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	return iter.actual.clave, iter.actual.dato
}

func (hash *hashAbierto[K, V]) Iterar(visitar func(clave K, dato V) bool) {
	for _, lista := range hash.tabla {
		if lista.EstaVacia() {
			continue
		}
		iter := lista.Iterador()
		for iter.HaySiguiente() {
			par := iter.VerActual()
			clave, valor := par.clave, par.dato
			if visitar(clave, valor) {
				iter.Siguiente()
			} else {
				return
			}
		}
	}
}
