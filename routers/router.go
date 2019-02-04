package routers

import "github.com/gorilla/mux"

func InitRouters() *mux.Router  {
	router := mux.NewRouter().StrictSlash(false)
	router = SetIpCountryRouter(router)
	router = SetRootRouter(router)
	return router
}