package config

import (
	"github.com/spf13/viper"
)

type Redis struct {
	Host            string
	Port            string
	DB              int
	CacheExpiration int
}

func NewConfig() *Config {
	viper.AutomaticEnv()
	// postgres
	viper.SetDefault("DB_HOST", "postgres")
	viper.SetDefault("DB_PORT", "5432")
	viper.SetDefault("DB_USER", "postgres")
	viper.SetDefault("DB_PASSWORD", "12345")
	viper.SetDefault("DB_NAME", "postgres")

	// redis
	viper.SetDefault("REDIS_HOST", "localhost")
	viper.SetDefault("REDIS_PORT", 6379)
	viper.SetDefault("REDIS_DB", 0)
	viper.SetDefault("CACHE_EXPIRATION", 20)

	// service
	viper.SetDefault("SERVICE_PORT", 8080)

	return &Config{
		PgDB: Pg{
			Host:     viper.GetString("DB_HOST"),
			Port:     viper.GetString("DB_PORT"),
			User:     viper.GetString("DB_USER"),
			Password: viper.GetString("DB_PASSWORD"),
			Name:     viper.GetString("DB_NAME"),
		},
		Redis: Redis{
			Host:            viper.GetString("REDIS_HOST"),
			Port:            viper.GetString("REDIS_PORT"),
			DB:              viper.GetInt("REDIS_DB"),
			CacheExpiration: viper.GetInt("CACHE_EXPIRATION"),
		},
		Server: Server{
			Port: viper.GetString("SERVICE_PORT"),
		},
	}
}
