package restroutes

import (
	"net/http"
	"reflect"
)

type handlerInterface interface {
	HandleFunc(w http.ResponseWriter, r *http.Request)
}

type Route struct {
	Receiver interface{}
	Method   string
}

type Routes map[string]Route

func Register(m handlerInterface, routes Routes) {
	for k, v := range routes {
		s := reflect.ValueOf(v.Receiver).MethodByName(v.Method)
		methodIface := s.Interface()
		method := methodIface.(func(w http.ResponseWriter, r *http.Request))
		m.HandleFunc(k, method)
	}
}
