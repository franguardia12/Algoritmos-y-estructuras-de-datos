package aeropuerto

import "strings"

type vueloImplementacion struct {
	informacion [CANT_INFORMACION]string
}

func CrearVuelo(informacion [CANT_INFORMACION]string) Vuelo {
	vuelo := new(vueloImplementacion)
	vuelo.informacion = informacion
	return vuelo
}

func (vuelo *vueloImplementacion) ObtenerPrioridad() string {
	return vuelo.informacion[PRIORITY]
}

func (vuelo *vueloImplementacion) ObtenerFecha() string {
	return vuelo.informacion[DATE]
}

func (vuelo *vueloImplementacion) ObtenerCodigoVuelo() string {
	return vuelo.informacion[FLIGHT_NUMBER]
}

func (vuelo *vueloImplementacion) ObtenerOrigenYdestino() (string, string) {
	return vuelo.informacion[ORIGIN_AIRPORT], vuelo.informacion[DESTINATION_AIRPORT]
}

func (vuelo *vueloImplementacion) ObtenerInfo() string {
	return strings.Join(vuelo.informacion[:], " ")
}
