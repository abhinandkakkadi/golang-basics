package util

import "github.com/spf13/viper"

type DbConfig struct {
	Address  string `mapstructure:"address"`
	Port     string `mapstructure:"port"`
	Password string `mapstructure:"password"`
	User     string `mapstructure:"user"`
}

type Config struct {
	Db     DbConfig     `mapstructure:"db"`
	Server ServerConfig `mapstructure:"server"`
}

type ServerConfig struct {
	Port string `mapstructure:"port"`
}

// this is what viper.New() actually returns
var vp *viper.Viper

func LoadConfig() (Config,error) {

	vp = viper.New()
	var config Config

	vp.SetConfigName("config")
	vp.SetConfigType("json")

	// viper can search in multiple directories
	vp.AddConfigPath(".")
	vp.AddConfigPath("./util")

	err := vp.ReadInConfig()
	if err != nil {
		return Config{},err
	}

	err = vp.Unmarshal(&config)
	if err != nil {
		return Config{},err
	}
	
	return config,nil
}