package mdenv

import (
	"log"
	"os"
	"path"

	"github.com/byliuyang/app/fw"
	"github.com/joho/godotenv"
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
	workDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	_, err = os.Stat(path.Join(workDir, ".env"))
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
