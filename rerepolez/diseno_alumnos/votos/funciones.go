package votos

import (
	"fmt"
	"rerepolez/diseno_alumnos/errores"
	"strconv"
)

func TipoVotoEsValido(voto string) bool {
	if voto != "Presidente" && voto != "Gobernador" && voto != "Intendente" {
		return false
	}
	return true
}

func ConvertirAEntero(str string) (int, error) {
	convertido, err := strconv.Atoi(str)
	return convertido, err
}

func DefinirVoto(voto string) TipoVoto {
	if voto == "Presidente" {
		return PRESIDENTE
	} else if voto == "Intendente" {
		return INTENDENTE
	} else if voto == "Gobernador" {
		return GOBERNADOR
	} else {
		panic(errores.ErrorTipoVoto{})
	}
}

func YaFueVotante(votante Votante, votantes_pasados []Votante) (int, bool) {
	for indice, votante_actual := range votantes_pasados {
		if votante.LeerDNI() == votante_actual.LeerDNI() {
			return indice, true
		}
	}
	return -1, false
}

func ChequearVotanteFraudulento(votante Votante, VotantesPasados []Votante) Votante {
	i, es_fraudulento := YaFueVotante(votante, VotantesPasados)
	if es_fraudulento {
		votante = VotantesPasados[i]
	}
	return votante
}

func Imprimir_resultado(partidos []Partido, impugnados int) {
	voto := PRESIDENTE
	for i := 0; i < 3; i++ {
		if i == 0 {
			fmt.Println("Presidente:")
			voto = PRESIDENTE
		} else if i == 1 {
			fmt.Println("Gobernador:")
			voto = GOBERNADOR
		} else {
			fmt.Println("Intendente:")
			voto = INTENDENTE
		}
		fmt.Println(partidos[len(partidos)-1].ObtenerResultado(voto))
		for j := 0; j < len(partidos)-1; j++ {
			fmt.Println(partidos[j].ObtenerResultado(voto))
		}
		fmt.Println()
	}
	impugnados_str := strconv.Itoa(impugnados)
	res := "Votos Impugnados: " + impugnados_str + " voto"
	res = Pluralizar(impugnados, res)
	fmt.Println(res)
}

func Pluralizar(votos int, cadena string) string {
	if votos == 0 || votos > 1 {
		cadena += "s"
	}
	return cadena
}
