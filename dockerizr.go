package main

import (
	"encoding/json"
	"os"
)

func main() {
	configFile, err := os.Open("dockerizr.json")
	if err != nil {
		panic(err)
	}
	config := Config{}
	jsonParser := json.NewDecoder(configFile)
	if err = jsonParser.Decode(&config); err != nil {
		panic(err)
	}
	config.setAppPort()
	copyAndMoveTemplates(config)
}
