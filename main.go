package main

import (
	"encoding/json"
	"github.com/hecomp/ipCountryValidatorApi/ipcountryapi"
	"log"
	"os"
)

type configuration struct {
	ServerAddress string `json:"webserver"`
}

func main() {
	file, err := os.Open("config.json")
	if err!=nil {
		log.Fatal(err)
	}

	config := new(configuration)
	json.NewDecoder(file).Decode(config)
	// Start the server on port config.json
	log.Println("Starting web server on address ", config.ServerAddress)
	ipcountryapi.RunIpCountryApi(config.ServerAddress)
}