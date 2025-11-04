package utilities

import (
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

func LoadEnv(envFileName string) {
	environment := os.Getenv("APP_ENV")
	if environment == "docker" {
		return
	}
	dir, _ := os.Getwd()
	for {
		try := filepath.Join(dir, envFileName)
		if _, err := os.Stat(try); err == nil {
			if err := godotenv.Load(try); err != nil {
				log.Panicf("could not load .env file: %v", err)
			}
			return
		}

		parent := filepath.Dir(dir)
		if parent == dir {
			log.Panic("no env file found")
		}
		dir = parent
	}
}
