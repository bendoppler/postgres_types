package config

import (
	"fmt"
	"sync"

	"github.com/spf13/viper"
)

type ServerConfig struct {
	DbUsername string
	DbPassword string
	DbHost     string
	DbPort     string
	DbName     string
}

//singleton server config
var (
	instance *ServerConfig
	once     sync.Once
)

//get singleton: first time access create server config, next time return instance
func GetInst() *ServerConfig {
	once.Do(func() {
		instance = readConfig()
	})
	return instance
}

func readConfig() *ServerConfig {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	conf := ServerConfig{}
	//db info
	conf.DbUsername = viper.GetString("database.username")
	conf.DbPassword = viper.GetString("database.password")
	conf.DbHost = viper.GetString("database.host")
	conf.DbPort = viper.GetString("database.port")
	conf.DbName = viper.GetString("database.name")
	return &conf
}
