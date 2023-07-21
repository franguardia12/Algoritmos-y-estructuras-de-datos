package mesa_electoral

import (
	"rerepolez/diseno_alumnos/errores"
	TDAVoto "rerepolez/diseno_alumnos/votos"
	TDACola "tdas/cola"
)

func ComandoIngresar(ingresado []string, padrones []int) (TDAVoto.Votante, error) {
	if len(ingresado) != 2 {
		err := errores.ErrorParametros{}
		return nil, err
	}
	dni_str := ingresado[1]
	dni, err := TDAVoto.ConvertirAEntero(dni_str)
	if err != nil {
		err = errores.DNIError{}
		return nil, err
	}
	if dni <= 0 {
		err = errores.DNIError{}
		return nil, err
	}
	if !EstaEnPadron(dni, padrones) {
		err = errores.DNIFueraPadron{}
		return nil, err
	}
	votante := TDAVoto.CrearVotante(dni)
	return votante, nil
}

func ComandoVotar(ingresado []string, votantes TDACola.Cola[TDAVoto.Votante], partidos []TDAVoto.Partido, VotantesPasados []TDAVoto.Votante) (error, error) {
	if len(ingresado) != 3 {
		err := errores.ErrorParametros{}
		return err, nil
	}
	tipoVoto, numeroLista := ingresado[1], ingresado[2]
	if votantes.EstaVacia() {
		err := errores.FilaVacia{}
		return err, nil
	}
	if !TDAVoto.TipoVotoEsValido(tipoVoto) {
		err := errores.ErrorTipoVoto{}
		return err, nil
	}
	numero, err := TDAVoto.ConvertirAEntero(numeroLista)
	if err != nil {
		err = errores.ErrorAlternativaInvalida{}
		return err, nil
	}
	if numero < 0 || numero > len(partidos)-1 {
		err = errores.ErrorAlternativaInvalida{}
		return err, nil
	}
	votante := votantes.VerPrimero()
	votante = TDAVoto.ChequearVotanteFraudulento(votante, VotantesPasados)
	voto := TDAVoto.DefinirVoto(tipoVoto)
	error_voto := votante.Votar(voto, numero)
	if error_voto != nil {
		return nil, error_voto
	}
	return nil, nil
}

func ComandoDeshacer(ingresado []string, votantes TDACola.Cola[TDAVoto.Votante], VotantesPasados []TDAVoto.Votante) (TDAVoto.Votante, error) {
	if votantes.EstaVacia() {
		err := errores.FilaVacia{}
		return nil, err
	}
	votante := votantes.VerPrimero()
	votante = TDAVoto.ChequearVotanteFraudulento(votante, VotantesPasados)
	err := votante.Deshacer()
	if err != nil {
		return votante, err
	}
	return votante, nil
}

func ComandoFinVotar(votantes TDACola.Cola[TDAVoto.Votante], partidos []TDAVoto.Partido, VotantesPasados []TDAVoto.Votante) (TDAVoto.Voto, error, error) {
	var voto TDAVoto.Voto
	if votantes.EstaVacia() {
		err := errores.FilaVacia{}
		return voto, err, nil
	}
	votante := votantes.VerPrimero()
	votante = TDAVoto.ChequearVotanteFraudulento(votante, VotantesPasados)
	voto_emitido, err := votante.FinVoto(partidos)
	if err != nil {
		return voto, nil, err
	}
	return voto_emitido, nil, nil
}
