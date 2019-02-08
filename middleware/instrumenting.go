package middleware

import (
	"fmt"
	"github.com/go-kit/kit/metrics"
	"github.com/hecomp/ipCountryValidatorApi/models"
	"github.com/hecomp/ipCountryValidatorApi/service"
	"time"
)

type InstrumentingMiddleware struct {
	RequestCount   metrics.Counter
	RequestLatency metrics.Histogram
	Service service.IpCountryService
}

// NewInstrumentingMiddleware returns a new instance of IpCountryService with metrics counter and histogram
func NewInstrumentingMiddleware(counter metrics.Counter, latency metrics.Histogram, s service.IpCountryService) service.IpCountryService {
	return &InstrumentingMiddleware{
		RequestCount:   counter,
		RequestLatency: latency,
		Service:        s,
	}
}

// ValidateIpCountry returns `isListed` indicator or error
func (s InstrumentingMiddleware) ValidateIpCountry(client models.GatewayClient) (isListed bool, err error) {
	defer func (begin time.Time) {
		lvs := []string{"method", "ValidateIpCountry", "error", fmt.Sprint(err)}
		s.RequestCount.With(lvs...).Add(1)
		s.RequestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())
	return s.Service.ValidateIpCountry(client)
}


