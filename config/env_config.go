package config

import "github.com/spf13/viper"

type Config struct {
	LimitedByIP    bool   `mapstructure:"IS_LIMITED_BY_IP"`
	LimitedByToken bool   `mapstructure:"IS_LIMITED_BY_TOKEN"`
	ServerPort     string `mapstructure:"SERVER_PORT"`
	RedisHost      string `mapstructure:"REDIS_HOST"`
	RedisPort      string `mapstructure:"REDIS_PORT"`
	RedisPassword  string `mapstructure:"REDIS_PASSWORD"`
	RedisDB        int    `mapstructure:"REDIS_DB"`
	Driver         string `mapstructure:"DRIVER"`
}

func LoadConfig(envPath string) (*Config, error) {
	var cfg Config
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath(envPath)
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
