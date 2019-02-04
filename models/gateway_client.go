package models

type GatewayClient struct {
	IP        string    `json:"ip"`
	Countries []Country `json:"countries"`
}
