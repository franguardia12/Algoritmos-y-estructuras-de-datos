package mesa_electoral

import (
	"bufio"
	"fmt"
	"os"
	"rerepolez/diseno_alumnos/errores"
	TDAVoto "rerepolez/diseno_alumnos/votos"
	"strings"
)

func LeerArchivos() (*os.File, *os.File, error) {
	params := os.Args[1:]
	if len(params) != 2 {
		err := errores.ErrorParametros{}
		return nil, nil, err
	}
	ruta_archivo_lista, ruta_archivo_padron := params[0], params[1]
	archivo_lista, err1 := os.Open(ruta_archivo_lista)
	if err1 != nil {
		err1 = errores.ErrorLeerArchivo{}
		return nil, nil, err1
	}

	archivo_padron, err2 := os.Open(ruta_archivo_padron)
	if err2 != nil {
		err2 := errores.ErrorLeerArchivo{}
		return nil, nil, err2
	}

	return archivo_lista, archivo_padron, nil
}

// Como el arreglo de padrones está ordenado realizo búsquedas binarias en él
func EstaEnPadron(dni int, padrones []int) bool {
	return _estaEnPadron(dni, padrones, 0, len(padrones))
}

func _estaEnPadron(dni int, padrones []int, inicio, fin int) bool {
	//busqueda binaria
	if inicio >= fin {
		return false
	}
	medio := (inicio + fin) / 2
	if padrones[medio] == dni {
		return true
	} else if padrones[medio] > dni {
		return _estaEnPadron(dni, padrones, inicio, medio)
	} else {
		return _estaEnPadron(dni, padrones, medio+1, fin)
	}
}

func LeerPartidos(archivo_lista *os.File) []TDAVoto.Partido {
	partidos := []TDAVoto.Partido{}
	scaner_partidos := bufio.NewScanner(archivo_lista)
	for scaner_partidos.Scan() {
		linea := scaner_partidos.Text()
		linea = strings.TrimSuffix(linea, "\n")
		lista := strings.Split(linea, ",")
		nombre := lista[0]
		var candidatos [TDAVoto.CANT_VOTACION]string
		for j := 0; j < 3; j++ {
			i := j + 1
			candidatos[j] = lista[i]
		}
		partido := TDAVoto.CrearPartido(nombre, candidatos)
		partidos = append(partidos, partido)
	}
	partido_en_blanco := TDAVoto.CrearVotosEnBlanco()
	partidos = append(partidos, partido_en_blanco)
	return partidos
}

func LeerPadrones(archivo_padron *os.File) []int {
	padrones := []int{}
	scaner_padrones := bufio.NewScanner(archivo_padron)
	for scaner_padrones.Scan() {
		dni_str := scaner_padrones.Text()
		dni, err := TDAVoto.ConvertirAEntero(dni_str)
		if err != nil {
			err := errores.DNIError{}
			fmt.Println(err)
		}
		if dni <= 0 {
			err := errores.DNIError{}
			fmt.Println(err)
		}
		padrones = append(padrones, dni)
	}
	radixSort(padrones)
	return padrones
}

func radixSort(arr []int) {
	max := maximo(arr)

	for exp := 1; max/exp > 0; exp *= 10 {
		countingSort(arr, exp)
	}
}

func maximo(arreglo []int) int {
	max := arreglo[0]
	for _, numero := range arreglo {
		if numero > max {
			max = numero
		}
	}
	return max
}

func countingSort(arr []int, exp int) {
	largo := len(arr)
	salida := make([]int, largo)
	frecuencias := make([]int, 10)

	for i := 0; i < largo; i++ {
		indice := (arr[i] / exp) % 10
		frecuencias[indice]++
	}

	for i := 1; i < 10; i++ {
		frecuencias[i] += frecuencias[i-1]
	}

	for i := largo - 1; i >= 0; i-- {
		indice := (arr[i] / exp) % 10
		salida[frecuencias[indice]-1] = arr[i]
		frecuencias[indice]--
	}

	for i := 0; i < largo; i++ {
		arr[i] = salida[i]
	}
}
