package handlers

import (
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/hecomp/ipCountryValidatorApi/middleware"
	"github.com/hecomp/ipCountryValidatorApi/transport"
)

//
var ipSvc  = middleware.LoggingMonitorService()

// PostValidateIpCountryHandler expose the service to the network.
// httptransport act as a handler
var (
	PostValidateIpCountryHandler = httptransport.NewServer(
		transport.MakePostValidateIpCountry(ipSvc),
		transport.DecodeIpCountry,
		transport.EncodeResponse)
)

var (
	RootHandler = httptransport.NewServer(
		transport.MakeRoot(),
		transport.DecodeRoot,
		transport.EncodeResponse)
)

