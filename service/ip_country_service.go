package service

import (
	"errors"
	"github.com/hecomp/ipCountryValidatorApi/common"
	"github.com/hecomp/ipCountryValidatorApi/models"
)

type IpCountryService interface {
	ValidateIpCountry(client models.GatewayClient) (bool, error)
}

var (
	ErrEmpty = errors.New("empty gatewayclient")
)

type IpCountryValidatorService struct {}

func (IpCountryValidatorService) ValidateIpCountry(client models.GatewayClient) (bool, error) {
	if client.IP == "" &&  len(client.Countries) == 0 {
		return false, ErrEmpty
	}
	isListed := common.ContainsIp(client.Countries, client.IP)
	if !isListed {
		return false, nil
	}
	return true, nil
}