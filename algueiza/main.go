package main

import (
	TDAAeropuerto "algueiza/diseno/aeropuerto"
	"algueiza/diseno/comandos"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	algueiza := TDAAeropuerto.CrearAeropuerto()
	entrada := bufio.NewScanner(os.Stdin)
	for entrada.Scan() {
		ingresado := strings.Split(entrada.Text(), " ")
		comando := ingresado[0]
		switch comando {
		case "agregar_archivo":
			err := comandos.ComandoAgregarArchivo(algueiza, ingresado)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				continue
			}
			fmt.Println("OK")
		case "info_vuelo":
			res, err := comandos.ComandoInfoVuelo(algueiza, ingresado)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				continue
			}
			fmt.Println(res)
			fmt.Println("OK")
		case "prioridad_vuelos":
			err := comandos.ComandoPrioridadVuelos(algueiza, ingresado)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				continue
			}
			fmt.Println("OK")
		case "siguiente_vuelo":
			res, err := comandos.ComandoSiguienteVuelo(algueiza, ingresado)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				continue
			}
			if res != "" {
				fmt.Println(res)
			}
			fmt.Println("OK")
		case "ver_tablero":
			err := comandos.ComandoVerTablero(algueiza, ingresado)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				continue
			}
			fmt.Println("OK")
		case "borrar":
			err := comandos.ComandoBorrar(algueiza, ingresado)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				continue
			}
			fmt.Println("OK")
		}
	}
}
