package config

import "github.com/spf13/viper"

type Config struct {
	APIPORT   string `mapstructure:"APIPORT"`
	SECRETKEY string `mapstructure:"JWTKEY"`
	USERPORT  string `mapstructure:"USERPORT"`
}

func LoadConfig() (*Config, error) {
	var config Config
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}
	err = viper.Unmarshal(&config)
	if err != nil {
		return nil, err
	}
	return &config, nil

}