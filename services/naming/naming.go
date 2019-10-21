package naming

import (
	"RMid/distribution/clientproxy"
)

type NamingService struct {
	Repository map[string]clientproxy.ConversorProxy
}

func (naming *NamingService) Register(name string, proxy clientproxy.ConversorProxy) bool {

	r := false

	// check if repository is already created
	if len(naming.Repository) == 0 {
		naming.Repository = make(map[string]clientproxy.ConversorProxy)
	}

	// check if the service is already registered
	_, ok := naming.Repository[name]

	if ok {
		r = false // service already registered
	} else { // service not registered
		naming.Repository[name] = clientproxy.ConversorProxy{TypeName: proxy.TypeName, Host: proxy.Host, Port: proxy.Port}
		r = true
	}

	return r
}

func (naming NamingService) Lookup(name string) clientproxy.ConversorProxy {

	return naming.Repository[name]
}

func (naming NamingService) List(name string) map[string]clientproxy.ConversorProxy {

	return naming.Repository
}
