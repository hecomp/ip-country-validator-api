package routers

import (
	"github.com/gorilla/mux"
	"github.com/hecomp/ipCountryValidatorApi/handlers"
)

func SetIpCountryRouter(router *mux.Router) *mux.Router {
	router.Handle("/validate/ip", handlers.PostValidateIpCountryHandler).Methods("POST")
	return router
}