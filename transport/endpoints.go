package transport

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/go-kit/kit/endpoint"
	"github.com/hecomp/ipCountryValidatorApi/models"
	"github.com/hecomp/ipCountryValidatorApi/service"
	"net/http"
)

var (
	ErrBadRouting = errors.New("Error bad routing")
	ErrNotFound   = errors.New("Asset not found\n")
	ErrBadRequest = errors.New("Bad Request")
)

type IndicatorResponse struct {
	IsListed bool   `json:"isListed"`
	Err      string `json:"err,omitempty"`
}

type RootResponse struct {
	Status string `json:"Status"`
	Err    string `json:"err,omitempty"`
}

func MakePostValidateIpCountry(svc service.IpCountryService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(models.GatewayClient)
		isListed, err := svc.ValidateIpCountry(req)
		if err != nil {
			return IndicatorResponse{isListed, err.Error()}, nil
		}
		return IndicatorResponse{isListed, ""},nil
	}
}

func DecodeIpCountry(_ context.Context, r *http.Request) (interface{}, error) {
	var gatewayClient models.GatewayClient

	if r.Body == nil {
		return nil, ErrBadRequest
	}
	err := json.NewDecoder(r.Body).Decode(&gatewayClient)
	if err != nil {
		return nil, errors.New(err.Error())
	} else {
		return gatewayClient, nil
	}
}

func MakeRoot() endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		return RootResponse{"Running IP Country Validator API v1\n", ""},nil
	}
}

func DecodeRoot(_ context.Context, r *http.Request) (interface{}, error) {
	if r.URL.Path != "/" {
		return nil, ErrNotFound
	} else {
		return "", nil
	}
}

func EncodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Add("content-type", "application/json")
	return json.NewEncoder(w).Encode(response)
}