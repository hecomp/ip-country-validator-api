package ipcountryapi

import (
	"github.com/go-kit/kit/log"
	"github.com/hecomp/ipCountryValidatorApi/routers"
	"github.com/urfave/negroni"
	"net/http"
	"os"
)

// RunIpCountryApi initialize server
func RunIpCountryApi(addr string) error {
	logger := log.NewLogfmtLogger(os.Stderr)

	r := routers.InitRouters()

	// Stack of Middleware Handlers
	n := negroni.Classic()
	n.UseHandler(r)

	listener := http.ListenAndServe(addr, n)

	logger.Log("msg", "HTTP", "addr", addr)
	logger.Log("err", listener)
	return listener
}
