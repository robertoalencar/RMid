package repository

import (
	"Plataformas-Exercicio03/ConversorMoedasApp/meumiddleware/cliente/proxies"
	"Plataformas-Exercicio03/RobertoMiddleware/distribution/clientproxy"
	"reflect"
)

func CheckRepository(proxy clientproxy.ConversorProxy) interface{} {

	var clientProxy interface{}

	switch proxy.TypeName {

	case reflect.TypeOf(proxies.ConversorProxy{}).String():
		conversorProxy := proxies.NewConversorProxy()
		conversorProxy.Proxy.TypeName = proxy.TypeName
		conversorProxy.Proxy.Host = proxy.Host
		conversorProxy.Proxy.Port = proxy.Port
		clientProxy = conversorProxy
	}

	return clientProxy
}
