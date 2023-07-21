package diccionario

import (
	TDAPila "tdas/pila"
)

type nodoAbb[K comparable, V any] struct {
	izquierdo *nodoAbb[K, V]
	derecho   *nodoAbb[K, V]
	clave     K
	dato      V
}

func crearNodoABB[K comparable, V any]() *nodoAbb[K, V] {
	nodo := new(nodoAbb[K, V])
	return nodo
}

type funcCmp[K comparable] func(K, K) int

type abb[K comparable, V any] struct {
	raiz     *nodoAbb[K, V]
	cantidad int
	cmp      funcCmp[K]
}

type iterDiccionarioOrdenado[K comparable, V any] struct {
	elementos TDAPila.Pila[*nodoAbb[K, V]]
	desde     *K
	hasta     *K
	ab        *abb[K, V]
}

func (ab *abb[K, V]) buscarClave(clave K, nodo, anterior *nodoAbb[K, V]) (*nodoAbb[K, V], *nodoAbb[K, V]) {
	if nodo == nil {
		return nodo, anterior
	}
	if ab.cmp(nodo.clave, clave) < 0 {
		return ab.buscarClave(clave, nodo.derecho, nodo)
	} else if ab.cmp(nodo.clave, clave) > 0 {
		return ab.buscarClave(clave, nodo.izquierdo, nodo)
	} else {
		return nodo, anterior
	}
}

func CrearABB[K comparable, V any](funcion_cmp func(K, K) int) DiccionarioOrdenado[K, V] {
	arbol := new(abb[K, V])
	var raiz *nodoAbb[K, V]
	arbol.raiz = raiz
	arbol.cmp = funcion_cmp
	return arbol
}

func (ab *abb[K, V]) Cantidad() int {
	return ab.cantidad
}

func (ab *abb[K, V]) Pertenece(clave K) bool {
	nodo, _ := ab.buscarClave(clave, ab.raiz, nil)
	return nodo != nil
}

func (ab *abb[K, V]) Obtener(clave K) V {
	nodo, _ := ab.buscarClave(clave, ab.raiz, nil)
	if nodo == nil {
		panic("La clave no pertenece al diccionario")
	}
	return nodo.dato
}

func (ab *abb[K, V]) Guardar(clave K, dato V) {
	nuevo := crearNodoABB[K, V]()
	nuevo.clave, nuevo.dato = clave, dato

	nodo, padre := ab.buscarClave(clave, ab.raiz, nil)
	if nodo != nil {
		nodo.dato = dato
	} else {
		if padre == nil {
			ab.raiz = nuevo
		} else if ab.cmp(padre.clave, clave) < 0 {
			padre.derecho = nuevo
		} else {
			padre.izquierdo = nuevo
		}
		ab.cantidad++
	}
}

func (nodo *nodoAbb[K, V]) cantHijos() int {
	if nodo.izquierdo == nil && nodo.derecho == nil {
		return 0
	}
	if (nodo.izquierdo == nil && nodo.derecho != nil) || (nodo.izquierdo != nil && nodo.derecho == nil) {
		return 1
	}
	return 2
}

func (nodo *nodoAbb[K, V]) buscarReemplazante() *nodoAbb[K, V] {
	//mas grande del subarbol izquierdo
	if nodo.derecho == nil {
		return nodo
	}
	return nodo.derecho.buscarReemplazante()
}

func (ab *abb[K, V]) Borrar(clave K) V {
	nodo, padre := ab.buscarClave(clave, ab.raiz, nil)
	if nodo == nil {
		panic("La clave no pertenece al diccionario")
	}
	cant_hijos := nodo.cantHijos()
	borrado := nodo.dato

	if cant_hijos == 0 {
		if padre == nil {
			ab.raiz = nil
		} else if ab.cmp(nodo.clave, padre.clave) < 0 {
			padre.izquierdo = nil
		} else {
			padre.derecho = nil
		}
	} else if cant_hijos == 1 {
		if nodo.izquierdo != nil {
			if padre == nil {
				ab.raiz = nodo.izquierdo
			} else if ab.cmp(nodo.clave, padre.clave) < 0 {
				padre.izquierdo = nodo.izquierdo
			} else {
				padre.derecho = nodo.izquierdo
			}
		} else {
			if padre == nil {
				ab.raiz = nodo.derecho
			} else if ab.cmp(nodo.clave, padre.clave) < 0 {
				padre.izquierdo = nodo.derecho
			} else {
				padre.derecho = nodo.derecho
			}
		}
	} else {
		reemplazante := nodo.izquierdo.buscarReemplazante()
		claveRemplazante := reemplazante.clave
		valorReemplazante := ab.Borrar(reemplazante.clave)
		ab.cantidad++ //este llamado de Borrar reemplazante nos restó 1 en cantidad
		nodo.clave, nodo.dato = claveRemplazante, valorReemplazante
	}
	ab.cantidad--
	return borrado
}

func (ab *abb[K, V]) IterarRango(desde *K, hasta *K, visitar func(clave K, dato V) bool) {
	ab._IterarRango(ab.raiz, desde, hasta, visitar)
}

func (ab *abb[K, V]) _IterarRango(nodo *nodoAbb[K, V], desde *K, hasta *K, visitar func(clave K, dato V) bool) bool {
	if nodo == nil {
		return false
	}

	debeCortar := ab._IterarRango(nodo.izquierdo, desde, hasta, visitar)
	if debeCortar {
		return true
	}

	if desde == nil {
		if (hasta != nil && ab.cmp(nodo.clave, *hasta) >= 0) || !visitar(nodo.clave, nodo.dato) {
			return true
		} else {
			return ab._IterarRango(nodo.derecho, desde, hasta, visitar)
		}
	} else {
		if (hasta == nil && ab.cmp(nodo.clave, *desde) >= 0) || (hasta != nil && ab.cmp(nodo.clave, *desde) >= 0 && ab.cmp(nodo.clave, *hasta) <= 0) {
			if !visitar(nodo.clave, nodo.dato) {
				return true
			}
		}
		ab._IterarRango(nodo.derecho, desde, hasta, visitar)
	}
	return false
}

func (ab *abb[K, V]) apilarElementos(nodo *nodoAbb[K, V], iter iterDiccionarioOrdenado[K, V], pila TDAPila.Pila[*nodoAbb[K, V]]) {
	if nodo == nil {
		return
	}
	if iter.desde == nil || !(ab.cmp(nodo.clave, *iter.desde) < 0) {
		pila.Apilar(nodo)
		ab.apilarElementos(nodo.izquierdo, iter, pila)
	} else {
		ab.apilarElementos(nodo.derecho, iter, pila)
	}
}

func (ab *abb[K, V]) IteradorRango(desde *K, hasta *K) IterDiccionario[K, V] {
	iter := new(iterDiccionarioOrdenado[K, V])
	iter.elementos = TDAPila.CrearPilaDinamica[*nodoAbb[K, V]]()
	iter.desde = desde
	iter.hasta = hasta
	iter.ab = ab
	ab.apilarElementos(iter.ab.raiz, *iter, iter.elementos)
	return iter
}

func (iter *iterDiccionarioOrdenado[K, V]) HaySiguiente() bool {
	cmp := iter.ab.cmp
	if iter.hasta == nil || iter.elementos.EstaVacia() {
		return !iter.elementos.EstaVacia()
	} else {
		return cmp(iter.elementos.VerTope().clave, *iter.hasta) <= 0
	}
}

func (iter *iterDiccionarioOrdenado[K, V]) VerActual() (K, V) {
	if iter.elementos.EstaVacia() || !iter.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	par := iter.elementos.VerTope()
	return par.clave, par.dato
}

func (iter *iterDiccionarioOrdenado[K, V]) Siguiente() {
	if !iter.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	desapilado := iter.elementos.Desapilar()
	if desapilado.derecho != nil {
		iter.ab.apilarElementos(desapilado.derecho, *iter, iter.elementos)
	}
}

func (ab *abb[K, V]) Iterar(visitar func(clave K, dato V) bool) {
	ab.IterarRango(nil, nil, visitar)
}

func (ab *abb[K, V]) Iterador() IterDiccionario[K, V] {
	iter := ab.IteradorRango(nil, nil)
	return iter
}
