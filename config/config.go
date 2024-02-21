package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Server   Server
	Database Database
}

type Server struct {
	Port string
}

type Database struct {
	User string
	Pass string
	Host string
	Port string
	Name string
}

func GetConfig() Config {

	viper.SetConfigFile("config.json")
	if err := viper.ReadInConfig(); err != nil {
		panic("Failed to get config.")
	}

	return Config{
		Server: Server{
			Port: viper.GetString("server.port"),
		},
		Database: Database{
			User: viper.GetString("database.user"),
			Pass: viper.GetString("database.pass"),
			Host: viper.GetString("database.host"),
			Port: viper.GetString("database.port"),
			Name: viper.GetString("database.name"),
		},
	}
}
