package middleware

import (
	"github.com/go-kit/kit/log"
	"github.com/hecomp/ipCountryValidatorApi/models"
	"github.com/hecomp/ipCountryValidatorApi/service"
	"time"
)

type LoggingMiddleware struct {
	Logger log.Logger
	Service service.IpCountryService
}

func NewLoggingService(logger log.Logger, s service.IpCountryService) service.IpCountryService {
	return &LoggingMiddleware{logger, s}
}

func (lw LoggingMiddleware) ValidateIpCountry(client models.GatewayClient) (isListed bool, err error) {
	defer func (begin time.Time) {
		_ = lw.Logger.Log(
			"method", "ValidateIpCountry",
			"id", client.IP,
			"countries", len(client.Countries),
			"took", time.Since(begin),
			//"err", err,
			)
	}(time.Now())
	return lw.Service.ValidateIpCountry(client)
}
