package config

import (
	"github.com/genof420/ferremas-api/internal/envutil"
)

type Config struct {
	General  General
	HTTP     HTTP
	Database Database
}

type General struct {
	Debug bool
}

type HTTP struct {
	Address string
	Port    int
	SSL     bool
	SSLCert string
	SSLKey  string
}

type Database struct {
	Host     string
	Port     int
	Username string
	Password string
	Database string
}

func Get() *Config {
	return &Config{
		General: General{
			Debug: envutil.GetEnvBool("DEBUG_MODE"),
		},
		HTTP: HTTP{
			Address: envutil.GetEnv("ADDRESS", "127.0.0.1"),
			Port:    envutil.GetEnvInt("PORT", 80),
			SSL:     envutil.GetEnvBool("SSL_ENABLED"),
			SSLCert: envutil.GetEnv("SSL_CERT", ""),
			SSLKey:  envutil.GetEnv("SSL_KEY", ""),
		},
		Database: Database{
			Host:     envutil.GetEnv("DB_HOST", "localhost"),
			Port:     envutil.GetEnvInt("DB_PORT", 5432),
			Username: envutil.GetEnv("DB_USERNAME", "postgres"),
			Password: envutil.GetEnv("DB_PASSWORD", "toor"),
			Database: envutil.GetEnv("DB_NAME", "ferremas"),
		},
	}
}
