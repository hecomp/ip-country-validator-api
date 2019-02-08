package models

type GatewayClient struct {
	IP        string    `json:"ip,omitempty"`
	Countries []Country `json:"countries,omitempty"`
}
