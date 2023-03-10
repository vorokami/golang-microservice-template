package config

import (
	"github.com/spf13/viper"
	"go.uber.org/zap/zapcore"
)

type DbConfig struct {
	ConnectionString string `mapstructure:"connectionstring"`
	Provider         string `mapstructure:"provider"`
}

type Config struct {
	DbUsers  DbConfig `mapstructure:"dbUsers"`
	DbPoints DbConfig `mapstructure:"dbPoints"`
	Server   struct {
		Port string `mapstructure:"port"`
	} `mapstructure:"server"`
	Zap struct {
		Level    zapcore.Level `mapstructure:"level"`
		LogsPath string        `mapstructure:"logspath"`
	} `mapstructure:"zap"`
	FIREBASE_AUTH_KEY map[string]string `mapstructure:"FIREBASE_AUTH_KEY"`
}

var vp *viper.Viper

func LoadConfig() (Config, error) {
	vp = viper.New()
	var config Config

	vp.SetConfigName("config")
	vp.SetConfigType("json")
	vp.AddConfigPath("./config")
	vp.AddConfigPath(".")
	err := vp.ReadInConfig()
	if err != nil {
		return Config{}, err
	}

	err = vp.Unmarshal(&config)
	if err != nil {
		return Config{}, err
	}

	return config, nil
}
