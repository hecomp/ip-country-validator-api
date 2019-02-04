package common

import "github.com/hecomp/ipCountryValidatorApi/models"

func ContainsIp(c []models.Country, ip string) bool {
	for _, n := range c {
		if ip == n.Ip {
			return true
		}
	}
	return false
}
