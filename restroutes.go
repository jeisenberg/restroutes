package restroutes

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/mux"
	"net/http"
	"reflect"
)

type handlerInterface interface {
	HandleFunc(string, func(http.ResponseWriter, *http.Request)) *mux.Route
}

type initializeInterface interface {
	Initialize()
}

type Route struct {
	Receiver      interface{}
	Method        string
	RequestMethod string
}

type Routes map[string]Route

func Register(m handlerInterface, routes Routes) {
	for k, v := range routes {
		s := reflect.ValueOf(v.Receiver).MethodByName(v.Method)
		methodIface := s.Interface()
		method := methodIface.(func(w http.ResponseWriter, r *http.Request))
		m.HandleFunc(k, method).Methods(v.RequestMethod)
	}
}

func RegisterGin(e *gin.Engine, routes Routes) {
	for k, v := range routes {
		r := reflect.ValueOf(v.Receiver)
		if _, ok := r.Interface().(initializeInterface); ok {
			r.MethodByName("Initialize").Call([]reflect.Value{})
		}
		s := reflect.ValueOf(v.Receiver).MethodByName(v.Method)
		methodIface := s.Interface()
		method := methodIface.(func(c *gin.Context))
		f := reflect.ValueOf(e).MethodByName(v.RequestMethod)
		arg := []reflect.Value{reflect.ValueOf(k), reflect.ValueOf(method)}
		f.Call(arg)
	}
}
