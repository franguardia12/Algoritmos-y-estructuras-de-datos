package comandos

import (
	TDAAeropuerto "algueiza/diseno/aeropuerto"
	"algueiza/diseno/funciones"
	"fmt"
	"strconv"
	"strings"
)

func ComandoAgregarArchivo(aeropuerto TDAAeropuerto.Aeropuerto, ingresado []string) error {
	if len(ingresado) != 2 {
		return fmt.Errorf("Error en comando agregar_archivo")
	}
	ruta := ingresado[1]
	archivo, err := funciones.LeerArchivos(ruta)
	if err != nil {
		return fmt.Errorf("Error en comando agregar_archivo")
	}
	defer archivo.Close()
	funciones.IngresarVuelos(archivo, aeropuerto)
	return nil

}

func ComandoInfoVuelo(aeropuerto TDAAeropuerto.Aeropuerto, ingresado []string) (string, error) {
	if len(ingresado) != 2 {
		return "", fmt.Errorf("Error en comando info_vuelo")
	}
	codigo := ingresado[1]
	res, err := aeropuerto.InfoVuelo(codigo)
	if err != nil {
		return "", fmt.Errorf("Error en comando info_vuelo")
	}
	return res, nil
}

func ComandoPrioridadVuelos(aeropuerto TDAAeropuerto.Aeropuerto, ingresado []string) error {
	if len(ingresado) != 2 {
		return fmt.Errorf("Error en comando prioridad_vuelos")
	}

	k := ingresado[1]

	convertido, err := strconv.Atoi(k)
	if err != nil || convertido < 0 {
		return fmt.Errorf("Error en comando prioridad_vuelos")
	}
	aeropuerto.MostrarPrioridad(convertido)
	return nil
}

func ComandoSiguienteVuelo(aeropuerto TDAAeropuerto.Aeropuerto, ingresado []string) (string, error) {
	if len(ingresado) != 4 {
		return "", fmt.Errorf("Error en comando siguiente_vuelo")
	}
	origen, destino, fecha := ingresado[1], ingresado[2], ingresado[3]
	vuelo := aeropuerto.SiguienteVuelo(origen, destino, fecha)
	if vuelo == nil {
		return "", nil
	}
	res := vuelo.ObtenerInfo()
	return res, nil
}

func ComandoVerTablero(aeropuerto TDAAeropuerto.Aeropuerto, ingresado []string) error {
	if len(ingresado) != 5 {
		return fmt.Errorf("Error en comando ver_tablero")
	}
	k := ingresado[1]
	modo := ingresado[2]
	desde, hasta := ingresado[3], ingresado[4]
	convertido, err := strconv.Atoi(k)
	if err != nil || (modo != "asc" && modo != "desc") {
		return fmt.Errorf("Error en comando ver_tablero")
	}
	if strings.Compare(desde, hasta) > 0 {
		return nil
	}
	if modo == "asc" {
		aeropuerto.VerTableroASC(convertido, desde, hasta)
	} else {
		aeropuerto.VerTableroDESC(convertido, desde, hasta)
	}
	return nil

}

func ComandoBorrar(aeropuerto TDAAeropuerto.Aeropuerto, ingresado []string) error {
	if len(ingresado) != 3 {
		return fmt.Errorf("Error en comando borrar")
	}
	desde, hasta := ingresado[1], ingresado[2]
	if strings.Compare(desde, hasta) > 0 {
		return nil
	}
	aeropuerto.Borrar(desde, hasta)
	return nil
}
