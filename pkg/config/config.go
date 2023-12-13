package config

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
)

var LocalConfig Config

type Config struct {
	Env       string    `mapstructure:"env", json:"env"`
	Redis     Redis     `mapstructure:"redis", json:"redis"`
	Endpoints Endpoints `mapstructure:"endpoints", json:"endpoints"`
}

type Redis struct {
	Address  string `mapstructure:"address", json:"address"`
	Password string `mapstructure:"password", json:"password"`
}

type Endpoints struct {
	Staging    Endpoint `mapstructure:"staging", json:"staging"`
	Production Endpoint `mapstructure:"production", json:"production"`
}

type Endpoint struct {
	ClientID      string `mapstructure:"client_id", json:"client_id"`
	Credentials   string `mapstructure:"credentials", json:"credentials"`
	Authorization string `mapstructure:"authorization", json:"authorization"`
	Token         string `mapstructure:"token", json:"token"`
	UserInfo      string `mapstructure:"user_info", json:"user_info"`
	Logout        string `mapstructure:"logout", json:"logout"`
}

func init() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")
	viper.AddConfigPath("../../config")
	viper.AddConfigPath("../../../config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("panic: logfile reading failure, %w", err))
	}

	err = viper.Unmarshal(&LocalConfig)
	if err != nil {
		panic(fmt.Errorf("panic: logfile Unmarshal failure, %w", err))
	}
}

func (c *Config) String() (string, error) {
	jsonData, err := json.Marshal(c)
	if err != nil {
		return "", err
	}
	return string(jsonData), nil
}
