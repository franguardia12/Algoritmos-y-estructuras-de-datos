package aeropuerto

const (
	FLIGHT_NUMBER = iota
	AIRLINE
	ORIGIN_AIRPORT
	DESTINATION_AIRPORT
	TAIL_NUMBER
	PRIORITY
	DATE
	DEPARTURE_DELAY
	AIR_TIME
	CANCELLED
)

const CANT_INFORMACION = CANCELLED + 1

// Vuelo modela a un vuelo específico, con toda la información asociada a él.
type Vuelo interface {
	//Devuelve la prioridad del vuelo.
	ObtenerPrioridad() string

	//Devuelve la fecha del vuelo.
	ObtenerFecha() string

	//Devuelve el codigo del vuelo.
	ObtenerCodigoVuelo() string

	//Devuelve el origen y el destino del vuelo.
	ObtenerOrigenYdestino() (string, string)

	//Devuelve la información del vuelo.
	ObtenerInfo() string
}
