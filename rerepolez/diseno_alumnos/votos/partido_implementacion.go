package votos

import (
	"strconv"
)

type partidoImplementacion struct {
	nombre           string
	candidatos       [CANT_VOTACION]string
	votos_candidatos [CANT_VOTACION]int
}

type partidoEnBlanco struct {
	votos [CANT_VOTACION]int
}

func CrearPartido(nombre string, candidatos [CANT_VOTACION]string) Partido {
	partido := new(partidoImplementacion)
	partido.nombre = nombre
	partido.candidatos = candidatos
	return partido
}

func CrearVotosEnBlanco() Partido {
	partidoenblanco := new(partidoEnBlanco)
	return partidoenblanco
}

func (partido *partidoImplementacion) VotadoPara(tipo TipoVoto) {
	partido.votos_candidatos[tipo] += 1
}

func (partido partidoImplementacion) ObtenerResultado(tipo TipoVoto) string {
	votos := partido.votos_candidatos[tipo]
	nombre_partido := partido.nombre
	candidato := partido.candidatos[tipo]
	cadena := strconv.Itoa(votos)
	resultado := nombre_partido + " - " + candidato + ": " + cadena + " voto"
	return Pluralizar(votos, resultado)
}

func (blanco *partidoEnBlanco) VotadoPara(tipo TipoVoto) {
	blanco.votos[tipo]++
}

func (blanco partidoEnBlanco) ObtenerResultado(tipo TipoVoto) string {
	votos := blanco.votos[tipo]
	cadena := strconv.Itoa(votos)
	resultado := "Votos en Blanco: " + cadena + " voto"
	return Pluralizar(votos, resultado)

}
