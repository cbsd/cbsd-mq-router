package main

import (
	"fmt"
	"os"
	"encoding/json"
)

type Config struct {
	Broker		string	`json:"broker"`
	CbsdColor	bool	`json:"cbsdcolor"`
	CbsdEnv		string	`json:"cbsdenv"`
	Logfile		string	`json:"logfile"`
	BeanstalkConfig		`json:"beanstalkd"`
}

func LoadConfiguration(file string) (Config,error) {
	var config Config
	configFile, err := os.Open(file)
	defer configFile.Close()

	if err != nil {
		fmt.Println(err.Error())
		return config, err
	}

	jsonParser := json.NewDecoder(configFile)
	err = jsonParser.Decode(&config)

	if err != nil {
		fmt.Printf("config error: %s: %s\n", file,err.Error())
		return config, err
	}

	return config, err
}
