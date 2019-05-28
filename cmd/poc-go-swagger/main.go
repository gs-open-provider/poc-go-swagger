package main

import (
	"github.com/gs-open-provider/poc-go-swagger/internal/configs"
	"github.com/gs-open-provider/poc-go-swagger/internal/logger"
)

func main() {
	// Initialize Viper across the application
	configs.InitializeViper()

	// Initialize Logger across the application
	logger.InitializeZapCustomLogger()
}
