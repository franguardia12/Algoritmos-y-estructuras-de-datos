package funciones

import (
	TDAAeropuerto "algueiza/diseno/aeropuerto"
	"bufio"
	"os"
	"strings"
)

func LeerArchivos(ruta string) (*os.File, error) {
	archivo, err := os.Open(ruta)
	if err != nil {
		return nil, err
	}
	return archivo, nil
}

func IngresarVuelos(archivo *os.File, aeropuerto TDAAeropuerto.Aeropuerto) {
	scaner := bufio.NewScanner(archivo)
	for scaner.Scan() {
		linea := scaner.Text()
		linea = strings.TrimSuffix(linea, "\n")

		//Guarda los elementos del split en un arreglo de tamaño CANT_INFORMACION (10)
		var informacion [TDAAeropuerto.CANT_INFORMACION]string
		lista := strings.Split(linea, ",")
		copy(informacion[:], lista)
		vuelo := TDAAeropuerto.CrearVuelo(informacion)
		aeropuerto.AgregarVuelo(vuelo)
	}
}
