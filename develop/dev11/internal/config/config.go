package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

// struct for config file
type Config struct {
	Host string
	Port string
	DB   DBconfig
}

// struct for database config
type DBconfig struct {
	Name string
	// Table    string
	// Username string
	// Password string
	// Host     string
}

// add info from config to config and db struct's
func ReadConfig(configPath string) *Config {
	data, err := ioutil.ReadFile(configPath)
	if err != nil {
		log.Fatal("Could not read config file:", err)
	}
	var conf Config
	err = json.Unmarshal(data, &conf)
	if err != nil {
		log.Fatal("Could not parse config file:", err)
	}
	return &conf
}
