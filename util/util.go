package util

import (
	"log"
)

const TAMANHO_AMOSTRA = 10000
const HOST = "127.0.0.1"
const PORTA = 12345
const NAMING_PORT = 1414
const MIOP_REQUEST = 1

type Request struct {
	MoedaDestino string
	Valor        float64
}

type Reply struct {
	Result float64
}

type Args struct {
	MoedaDestino string
	Valor        float64
}

func ChecaErro(err error, msg string) {

	if err != nil {
		log.Fatalf("%s!!: %s", msg, err)
	}
}
