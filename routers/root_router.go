package routers

import (
	"github.com/gorilla/mux"
	"github.com/hecomp/ipCountryValidatorApi/handlers"
)

func SetRootRouter(router *mux.Router) *mux.Router {
	router.Handle("/", handlers.RootHandler)
	return router
}