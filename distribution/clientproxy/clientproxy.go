package clientproxy

import (
	"RMid/aux"
	"RMid/distribution/requestor"
)

type ConversorProxy struct {
	Host     string
	Port     int
	Id       int
	TypeName string
}

func NewConversorProxy(host string, porta int, id int, typeName string) ConversorProxy {

	return ConversorProxy{Host: host, Port: porta, Id: id, TypeName: typeName}
}

func (proxy ConversorProxy) Converter(moedaDestino string, valor float32) float32 {

	//Prepara a invocação

	params := make([]interface{}, 2)
	params[0] = moedaDestino
	params[1] = valor
	request := aux.Request{Op: "Converter", Params: params}
	inv := aux.Invocation{Host: proxy.Host, Port: proxy.Port, Request: request}

	//Invoca o Requestor

	req := requestor.Requestor{}
	ter := req.Invoke(inv).([]interface{})

	return float32(ter[0].(float32))
}
