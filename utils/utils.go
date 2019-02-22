package utils

import (
	"log"
	"os"

	"github.com/midoblgsm/gogrpcrestswagger-boilerplate/resources"
)

func LoadConfig(logger *log.Logger) (resources.Config, error) {
	logger.Println("entering-utils-loadconfig")
	defer logger.Println("exiting-utils-loadconfig")
	config := resources.Config{}

	auth := os.Getenv("AUTHORIZATION")
	config.Authorization = auth

	return config, nil
}
