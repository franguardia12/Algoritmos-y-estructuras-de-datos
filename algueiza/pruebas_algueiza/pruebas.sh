#!/usr/bin/env bash

set -eu

PROGRAMA="$1"
BASENAME=$(basename "$PROGRAMA")
RET=0
VALGRIND_RETRY=""

# Correr con diff y sin Valgrind.
echo "Ejecución de pruebas unitarias de $BASENAME:"
echo ""

for t in *.test; do
    b=${t%.test}
    ret=0
    echo "$b $(< $t)"

    # TODO: mostrar diff de stderr sin importar la comparación de stdout.

    ($PROGRAMA <${b}_in >${b}_actual_out 2>${b}_actual_err && \
        diff -u ${b}_out ${b}_actual_out && \
        diff -u ${b}_err ${b}_actual_err && \
        echo "OK") || { RET=$?; ret=$RET; echo "ERROR"; }

    if [[ $ret -gt 128 ]]; then
        # Si $RET es mayor que 128, el programa murió por una señal, en
        # particular por la señal de valor numérico $RET - 128; por ejemplo
        # 11 (segmentation fault) o 6 (abort).
        #
        # Para este tipo de fallos, corremos de nuevo bajo Valgrind la prueba,
        # para dar una posibilidad de saber, sin acceso a la prueba en sí, por
        # dónde viene el error.
        #
        # TODO: quizá usar una lista fija de señales si hay algunas para las que
        # no se debería correr de nuevo bajo Valgrind.
        VALGRIND_RETRY="$VALGRIND_RETRY ${b}"
    fi

    echo ""
done

exit $RET
