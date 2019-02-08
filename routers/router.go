package routers

import "github.com/gorilla/mux"

func InitRouters() *mux.Router  {
	// Initialize router and dispatcher handling
	router := mux.NewRouter().StrictSlash(false)
	router = SetIpCountryRouter(router)
	router = SetRootRouter(router)
	return router
}