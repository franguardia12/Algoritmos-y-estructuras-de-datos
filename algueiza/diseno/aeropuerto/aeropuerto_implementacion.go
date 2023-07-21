package aeropuerto

import (
	"fmt"
	"strconv"
	"strings"
	TDAHeap "tdas/cola_prioridad"
	TDADiccionario "tdas/diccionario"
	TDAPila "tdas/pila"
)

func compararVuelos(a, b Vuelo) int {
	//Compara por prioridad. Si tienen la misma, desempata por codigo de vuelo
	p1, p2 := a.ObtenerPrioridad(), b.ObtenerPrioridad()
	prior1, _ := strconv.Atoi(p1)
	prior2, _ := strconv.Atoi(p2)
	res := prior1 - prior2
	if res == 0 {
		return strings.Compare(a.ObtenerCodigoVuelo(), b.ObtenerCodigoVuelo()) * -1
	}
	return res
}

func CompararFechaYCodigo(a, b FechaYCodigo) int {
	fecha1, fecha2 := a.fecha, b.fecha
	res := strings.Compare(fecha1, fecha2)
	if res == 0 {
		return strings.Compare(a.codigo, b.codigo)
	}
	return res
}

type aeropuertoImplementacion struct {
	vuelos     TDADiccionario.Diccionario[string, Vuelo]
	conexiones TDADiccionario.Diccionario[Conexion, TDADiccionario.DiccionarioOrdenado[FechaYCodigo, Vuelo]]
	tablero    TDADiccionario.DiccionarioOrdenado[string, TDADiccionario.DiccionarioOrdenado[string, string]]
}

func CrearAeropuerto() Aeropuerto {
	aeropuerto := new(aeropuertoImplementacion)
	aeropuerto.tablero = TDADiccionario.CrearABB[string, TDADiccionario.DiccionarioOrdenado[string, string]](strings.Compare)
	aeropuerto.vuelos = TDADiccionario.CrearHash[string, Vuelo]()
	aeropuerto.conexiones = TDADiccionario.CrearHash[Conexion, TDADiccionario.DiccionarioOrdenado[FechaYCodigo, Vuelo]]()
	return aeropuerto
}

func (aeropuerto *aeropuertoImplementacion) AgregarVuelo(vuelo Vuelo) {
	codigo, fecha := vuelo.ObtenerCodigoVuelo(), vuelo.ObtenerFecha()
	origen, destino := vuelo.ObtenerOrigenYdestino()
	conexion := Conexion{origen, destino}
	fechaYCodigo := FechaYCodigo{fecha, codigo}

	if aeropuerto.vuelos.Pertenece(codigo) {
		vueloAnterior := aeropuerto.vuelos.Obtener(codigo)
		fechaAnterior := vueloAnterior.ObtenerFecha()
		fechaYCodigoAnterior := FechaYCodigo{fechaAnterior, codigo}
		origenAnterior, destinoAnterior := vueloAnterior.ObtenerOrigenYdestino()
		conexionAnterior := Conexion{origenAnterior, destinoAnterior}
		aeropuerto.conexiones.Obtener(conexionAnterior).Borrar(fechaYCodigoAnterior)
		aeropuerto.tablero.Obtener(fechaAnterior).Borrar(codigo)
	}
	if !aeropuerto.tablero.Pertenece(fecha) {
		dicc := TDADiccionario.CrearABB[string, string](strings.Compare)
		dicc.Guardar(codigo, codigo)
		aeropuerto.tablero.Guardar(fecha, dicc)
	} else {
		aeropuerto.tablero.Obtener(fecha).Guardar(codigo, codigo)
	}
	if !aeropuerto.conexiones.Pertenece(conexion) {
		dicc := TDADiccionario.CrearABB[FechaYCodigo, Vuelo](CompararFechaYCodigo)
		dicc.Guardar(fechaYCodigo, vuelo)
		aeropuerto.conexiones.Guardar(conexion, dicc)
	} else {
		aeropuerto.conexiones.Obtener(conexion).Guardar(fechaYCodigo, vuelo)
	}
	aeropuerto.vuelos.Guardar(codigo, vuelo)

}

func (aeropuerto *aeropuertoImplementacion) InfoVuelo(codigo string) (string, error) {
	if !aeropuerto.vuelos.Pertenece(codigo) {
		return "", fmt.Errorf("El vuelo no pertenece al aeropuerto")
	}
	return aeropuerto.vuelos.Obtener(codigo).ObtenerInfo(), nil
}

func (aeropuerto *aeropuertoImplementacion) MostrarPrioridad(k int) {
	iter := aeropuerto.vuelos.Iterador()
	arreglo := []Vuelo{}
	for iter.HaySiguiente() {
		_, vuelo := iter.VerActual()
		arreglo = append(arreglo, vuelo)
		iter.Siguiente()
	}
	heap := TDAHeap.CrearHeapArr[Vuelo](arreglo, compararVuelos)

	for i := 0; i < k && !heap.EstaVacia(); i++ {
		vuelo := heap.Desencolar()
		prioridad, codigo := vuelo.ObtenerPrioridad(), vuelo.ObtenerCodigoVuelo()
		fmt.Println(prioridad + " - " + codigo)
	}
}

func (aeropuerto *aeropuertoImplementacion) SiguienteVuelo(origen, destino, fecha string) Vuelo {
	conexion := Conexion{origen, destino}
	if !aeropuerto.conexiones.Pertenece(conexion) {
		fmt.Println("No hay vuelo registrado desde " + origen + " hacia " + destino + " desde " + fecha)
		return nil
	}
	desde := FechaYCodigo{fecha, "-0000"}
	iter := aeropuerto.conexiones.Obtener(conexion).IteradorRango(&desde, nil)
	for iter.HaySiguiente() {
		_, vuelo := iter.VerActual()
		return vuelo
	}
	fmt.Println("No hay vuelo registrado desde " + origen + " hacia " + destino + " desde " + fecha)
	return nil
}

func (aeropuerto *aeropuertoImplementacion) VerTableroASC(k int, desde string, hasta string) {
	i := 0
	iter := aeropuerto.tablero.IteradorRango(&desde, &hasta)
	for iter.HaySiguiente() && i < k {
		fecha, vuelos := iter.VerActual()
		iter2 := vuelos.Iterador()
		for iter2.HaySiguiente() && i < k {
			codigo, _ := iter2.VerActual()
			fmt.Println(fecha + " - " + codigo)
			i++
			iter2.Siguiente()
		}
		iter.Siguiente()
	}
}

func (aeropuerto *aeropuertoImplementacion) VerTableroDESC(k int, desde string, hasta string) {
	pila := TDAPila.CrearPilaDinamica[string]()
	iter := aeropuerto.tablero.IteradorRango(&desde, &hasta)
	for iter.HaySiguiente() {
		fecha, vuelos := iter.VerActual()
		iter2 := vuelos.Iterador()
		for iter2.HaySiguiente() {
			codigo, _ := iter2.VerActual()
			cadena := fecha + " - " + codigo
			pila.Apilar(cadena)
			iter2.Siguiente()
		}
		iter.Siguiente()
	}

	i := 0
	for !pila.EstaVacia() && i < k {
		fmt.Println(pila.Desapilar())
		i++
	}
}

func (aeropuerto *aeropuertoImplementacion) Borrar(desde, hasta string) {
	iter := aeropuerto.tablero.IteradorRango(&desde, &hasta)
	for iter.HaySiguiente() {
		fecha, vuelos := iter.VerActual()
		iter2 := vuelos.Iterador()
		for iter2.HaySiguiente() {
			codigo, _ := iter2.VerActual()
			vuelo := aeropuerto.vuelos.Obtener(codigo)
			aeropuerto.vuelos.Borrar(codigo)
			origen, destino := vuelo.ObtenerOrigenYdestino()
			conexion := Conexion{origen, destino}
			fechaYCodigo := FechaYCodigo{fecha, codigo}
			aeropuerto.conexiones.Obtener(conexion).Borrar(fechaYCodigo)
			fmt.Println(vuelo.ObtenerInfo())
			iter2.Siguiente()
		}
		aeropuerto.tablero.Borrar(fecha)
		iter.Siguiente()
	}
}
