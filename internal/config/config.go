package config

import "os"

type Config struct {
	BotToken string
}

func Load() Config {
	return Config{
		BotToken: getEnv("BOT_TOKEN", ""),
	}
}

func getEnv(key, fallback string) string {
	val := os.Getenv(key)
	if val == "" {
		return fallback
	}
	return val
}