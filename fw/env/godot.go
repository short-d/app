package env

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var _ Env = (*GoDotEnv)(nil)

type GoDotEnv struct {
}

func (g GoDotEnv) GetVar(key string, defaultValue string) string {
	val := os.Getenv(key)
	if val == "" {
		return defaultValue
	}
	return val
}

func (g GoDotEnv) AutoLoadDotEnvFile() {
	_, err := os.Stat(".env")
	if os.IsNotExist(err) {
		return
	}

	err = godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func NewGoDotEnv() GoDotEnv {
	return GoDotEnv{}
}
