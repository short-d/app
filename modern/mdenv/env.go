package mdenv

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/short-d/app/fw"
)

var _ fw.Environment = (*GoDotEnv)(nil)

type GoDotEnv struct {
}

func (g GoDotEnv) GetEnv(key string, defaultValue string) string {
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
