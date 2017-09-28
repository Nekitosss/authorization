package authorization

import (
	n "github.com/Nekitosss/authorization/network"
	"net/http"

	"github.com/gorilla/mux"
)


type Route struct {
	Name string

	Methot string

	Pattern string

	HandlerFunc http.HandlerFunc
}


var routes = []Route{
	Route{"Register", "POST", "/v1/register", n.Register},
	Route{"ValidateRegister", "GET", "/v1/verify_register/{id}", n.VerifyRegistration},
	Route{"Login", "POST", "/v1/login", n.Login},
	Route{"ValidateSession", "POST", "/v1/validate_session", n.ValidateSession},
}


func addRoutes(router *mux.Router) {

	for _, route := range routes {
		router.
		Methods(route.Methot).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)

	}

}
