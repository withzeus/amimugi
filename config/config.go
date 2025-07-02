package config

import (
	"os"

	_ "github.com/joho/godotenv/autoload"
)

type LogConfig struct {
	Style string
	Level string
}

type AppConfig struct {
	Name                string
	Port                string
	Mode                string
	Version             string
	CookieDomain        string
	OverflowRedirectUrl string
}

type Auth0Config struct {
	Auth0Domain          string
	Auth0ClientId        string
	Auth0ClientSecret    string
	Auth0CallbackUrl     string
	OverflowClientID     string
	OverflowClientSecret string
}

type AWSConfig struct {
	AWSRegion string
}

var awsConfig AWSConfig
var appConfig AppConfig
var logConfig LogConfig
var auth0Config Auth0Config

type AllConfig struct {
	AppConfig
	AWSConfig
	LogConfig
	Auth0Config
}

var Env AllConfig = initConfig()

func initConfig() AllConfig {
	return AllConfig{
		AppConfig{
			Name:                getEnv("APP_NAME", ""),
			Port:                getEnv("APP_PORT", ""),
			Mode:                getEnv("APP_MODE", "staging"),
			Version:             getEnv("APP_VERSION", ""),
			CookieDomain:        getEnv("COOKIE_DOMAIN", "localhost"),
			OverflowRedirectUrl: getEnv("OVERFLOW_REDIRECT_URL", ""),
		},
		AWSConfig{
			AWSRegion: getEnv("AWS_REGION", ""),
		},
		LogConfig{
			Style: getEnv("LOG_STYLE", "json"),
			Level: getEnv("LOG_LEVEL", "info"),
		},
		Auth0Config{
			Auth0Domain:          getEnv("AUTH0_DOMAIN", ""),
			Auth0ClientId:        getEnv("AUTH0_CLIENT_ID", ""),
			Auth0ClientSecret:    getEnv("AUTH0_CLIENT_SECRET", ""),
			Auth0CallbackUrl:     getEnv("AUTH0_CALLBACK_URL", ""),
			OverflowClientID:     getEnv("OVERFLOW_CLIENT_ID", ""),
			OverflowClientSecret: getEnv("OVERFLOW_CLIENT_SECRET", ""),
		},
	}
}

func getEnv(key, fallback string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return fallback
	}
	return value
}
