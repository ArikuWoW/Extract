package initializers

import (
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func LoadEnv() {
	if err := godotenv.Load(); err != nil {
		logrus.Fatal("Error loading .env file")
	}
}
