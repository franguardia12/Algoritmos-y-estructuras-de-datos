package votos

import (
	"rerepolez/diseno_alumnos/errores"
	TDAPila "tdas/pila"
)

type VotoIndividual struct {
	tipoVoto    TipoVoto
	alternativa int
}

type votanteImplementacion struct {
	dni         int
	fraudulento bool
	pila_votos  TDAPila.Pila[VotoIndividual]
	votos       Voto
}

func CrearVotante(dni int) Votante {
	votante := new(votanteImplementacion)
	votante.dni = dni
	votante.pila_votos = TDAPila.CrearPilaDinamica[VotoIndividual]()
	return votante
}

func (votante votanteImplementacion) LeerDNI() int {
	return votante.dni
}

func (votante *votanteImplementacion) Votar(tipo TipoVoto, alternativa int) error {
	if votante.fraudulento {
		err := errores.ErrorVotanteFraudulento{}
		err.Dni = votante.LeerDNI()
		return err
	}
	if alternativa == 0 {
		votante.votos.Impugnado = true
	}
	//Revisar fin-voto para una posible incorporación de esta parte
	voto := new(VotoIndividual)
	voto.tipoVoto = tipo
	voto.alternativa = alternativa
	//En esta parte que apilamos el voto con esas dos partes, podríamos hacer un struct de los votos que vamos contando (que tenga como campos al tipo voto y
	//la alternativa la que vota). De esa forma solo accedemos a cada campo según sea necesario
	votante.pila_votos.Apilar(*voto)
	return nil
}

func (votante *votanteImplementacion) Deshacer() error {
	if votante.fraudulento {
		err := errores.ErrorVotanteFraudulento{}
		err.Dni = votante.LeerDNI()
		return err
	}
	if votante.pila_votos.EstaVacia() {
		return errores.ErrorNoHayVotosAnteriores{}
	}
	voto := votante.pila_votos.Desapilar()
	if voto.alternativa == 0 {
		votante.votos.Impugnado = false
	}
	return nil
}

func (votante *votanteImplementacion) FinVoto(partidos []Partido) (Voto, error) {
	if votante.fraudulento {
		err := errores.ErrorVotanteFraudulento{}
		err.Dni = votante.LeerDNI()
		return votante.votos, err
	}

	if votante.votos.Impugnado {
		return votante.votos, nil
	}

	for !votante.pila_votos.EstaVacia() {
		voto := votante.pila_votos.Desapilar()
		if votante.votos.VotoPorTipo[voto.tipoVoto] != 0 {
			continue
		}
		votante.votos.VotoPorTipo[voto.tipoVoto] = voto.alternativa
	}
	votos_en_blanco := partidos[len(partidos)-1]
	for tipovoto, alternativa := range votante.votos.VotoPorTipo {
		if alternativa == 0 {
			votos_en_blanco.VotadoPara(TipoVoto(tipovoto))
			continue
		}
		partidos[alternativa-1].VotadoPara(TipoVoto(tipovoto))
	}
	votante.fraudulento = true
	return votante.votos, nil
}
