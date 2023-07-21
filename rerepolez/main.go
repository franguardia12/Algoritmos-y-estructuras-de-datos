package main

import (
	"bufio"
	"fmt"
	"os"
	"rerepolez/diseno_alumnos/errores"
	TDAVoto "rerepolez/diseno_alumnos/votos"
	Mesa "rerepolez/mesa_electoral"
	"strings"
	"tdas/cola"
)

func main() {
	impugnados := 0
	archivo_lista, archivo_padron, err := Mesa.LeerArchivos()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer archivo_lista.Close()
	defer archivo_padron.Close()

	VotantesPasados := []TDAVoto.Votante{}
	padrones := Mesa.LeerPadrones(archivo_padron)
	partidos := Mesa.LeerPartidos(archivo_lista)
	votantes := cola.CrearColaEnlazada[TDAVoto.Votante]()

	entrada := bufio.NewScanner(os.Stdin)
	for entrada.Scan() {
		ingresado := strings.Split(entrada.Text(), " ")
		comando := ingresado[0]
		switch comando {
		case "ingresar":
			votante, err := Mesa.ComandoIngresar(ingresado, padrones)
			if err != nil {
				fmt.Println(err)
				continue
			}
			votantes.Encolar(votante)
			fmt.Println("OK")
		case "votar":
			err1, err2 := Mesa.ComandoVotar(ingresado, votantes, partidos, VotantesPasados)
			if err1 != nil {
				fmt.Println(err1)
				continue
			} else if err2 != nil {
				fmt.Println(err2)
				votantes.Desencolar()
				continue
			}
			fmt.Println("OK")

		case "deshacer":
			if len(ingresado) != 1 {
				continue
			}
			votante, err := Mesa.ComandoDeshacer(ingresado, votantes, VotantesPasados)
			if err != nil {
				fmt.Println(err)
				var yaVoto bool
				if votante != nil {
					_, yaVoto = TDAVoto.YaFueVotante(votante, VotantesPasados)
				}
				if yaVoto {
					votantes.Desencolar()
				}
				continue
			}
			fmt.Println("OK")
		case "fin-votar":
			if len(ingresado) != 1 {
				continue
			}
			voto_emitido, err1, err2 := Mesa.ComandoFinVotar(votantes, partidos, VotantesPasados)
			if err1 != nil {
				fmt.Println(err1)
				continue
			} else if err2 != nil {
				fmt.Println(err2)
				votantes.Desencolar()
				continue
			}
			if voto_emitido.EstaImpugnado() {
				impugnados++
			}
			VotantesPasados = append(VotantesPasados, votantes.Desencolar())
			fmt.Println("OK")
		default:
			continue
		}
	}

	if !votantes.EstaVacia() {
		err := errores.ErrorCiudadanosSinVotar{}
		fmt.Println(err)
	}
	TDAVoto.Imprimir_resultado(partidos, impugnados)
}
