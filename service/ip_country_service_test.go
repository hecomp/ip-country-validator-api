package service

import (
	"github.com/hecomp/ipCountryValidatorApi/models"
	"testing"
)

func TestValidateIpCountry(t *testing.T) {
	t.Log("")
	ipCountryService := new(IpCountryValidatorService)

	var countries = []models.Country{
		{
			Ip: "5.62.61.108/30",
			Name: "Ruanda",
		},
		{
			Ip: "5.62.61.148/30",
			Name: "Somália",
		},
	}

	clientGateway := models.GatewayClient{
		IP: "5.62.61.148/31",
		Countries: countries,
	}
	if _, err := ipCountryService.ValidateIpCountry(clientGateway); err != nil {
		t.Error(err.Error())
	} else {
		t.Log("Passed")
	}
}
