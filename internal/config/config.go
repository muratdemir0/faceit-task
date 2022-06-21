package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	Appname string
	Server  Server
	Mongo   Mongo
}

type Server struct {
	Port string
}

func New(configPath, configName string) (*Config, error) {
	v := viper.New()
	v.AddConfigPath(configPath)
	v.SetConfigName(configName)
	err := v.ReadInConfig()
	if err != nil {
		return nil, err
	}

	c := &Config{}

	if err := v.Unmarshal(c); err != nil {
		return nil, err
	}
	return c, nil
}

func (c Config) Print() {
	fmt.Println(c)
}
