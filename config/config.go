package config

import "github.com/tkanos/gonfig"

//Configuration ...
type Configuration struct {
	DB_NAME  	string
	DB_USERNAME string
	DB_PASSWORD string
	DB_PORT     string
	DB_HOST     string
}

//GetConfig ...
func GetConfig() Configuration {

	conf := Configuration{}
	gonfig.GetConf("config/config.json", &conf)
	
	return conf
	
}