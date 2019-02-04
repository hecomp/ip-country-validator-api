package middleware

import (
	"github.com/go-kit/kit/log"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
	kitprometheus "github.com/go-kit/kit/metrics/prometheus"
	"github.com/hecomp/ipCountryValidatorApi/service"
	"os"
)

func LoggingMonitorService() service.IpCountryService {
	var logger log.Logger
	logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)

	fieldKeys := []string{"method", "error"}

	var ics  service.IpCountryService
	ics = service.IpCountryValidatorService{}
	ics = NewLoggingService(log.With(logger, "component", "ip-country-validator"), ics)
	ics = NewInstrumentingMiddleware(
		kitprometheus.NewCounterFrom(stdprometheus.CounterOpts{
			Namespace: "ip_country_validator_api",
			Subsystem: "ip_country_validator_service",
			Name:      "request_count",
			Help:      "Number of requests received.",
		}, fieldKeys),
		kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
			Namespace: "ip_country_validator_api",
			Subsystem: "ip_country_validator_service",
			Name:      "request_latency_microseconds",
			Help:      "Total duration of requests in microseconds.",
		}, fieldKeys),
		ics,
	)
	return ics
}
