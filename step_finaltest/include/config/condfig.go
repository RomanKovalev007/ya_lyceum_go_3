package config

import "os"


const (
	Default_port = "8080"
)

type Config struct{
	Port string
}

func SetupConfig() *Config{
	conf := &Config{}
	serverPort := os.Getenv("SERVER_PORT")
	if serverPort == ""{
		serverPort = Default_port
	}
	
	conf.Port = serverPort

	return conf
}