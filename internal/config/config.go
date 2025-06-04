package config

import "os"

type Config struct {
	Dir string
	Key string
	BD  struct {
		Host, Port, Username, Password, BDName, SSLMode string
	}
}

func getEnv(key, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return def
}

func NewConfig() *Config {
	return &Config{
		Dir: getEnv("DIR", "../../static"),
		Key: getEnv("APP_KEY", "key"),
		BD: struct {
			Host, Port, Username, Password, BDName, SSLMode string
		}{
			Host:     getEnv("BD_HOST", "db"),
			Port:     getEnv("BD_PORT", "5432"),
			Username: getEnv("BD_USERNAME", "postgres"),
			Password: getEnv("BD_PASSWORD", "1111"),
			BDName:   getEnv("BD_NAME", "bookstore"),
			SSLMode:  getEnv("BD_SSLMODE", "disable"),
		},
	}
}
