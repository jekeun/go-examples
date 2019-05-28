package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Config struct {
	Database struct {
		Host     string `json:"host"`
		Password string `json:"password"`
	} `json:"database"`
	Host string `json:"host"`
	Port int `json:"port"`
}

func LoadConfiguration(file string) Config {
	var config Config
	configFile, err := os.Open(file)
	defer configFile.Close()
	if err != nil {
		fmt.Println(err.Error())
	}
	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&config)
	return config
}


func main() {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(dir)


	config := LoadConfiguration("./example/config/config.json")

	fmt.Println(config)
}
