package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AppPort             string
	AppSecret           string
	UserServiceUrl      string
	ShopServiceUrl      string
	ProductServiceUrl   string
	WarehouseServiceUrl string
	OrderServiceUrl     string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	config := &Config{
		AppPort:             os.Getenv("APP_PORT"),
		AppSecret:           os.Getenv("APP_SECRET"),
		UserServiceUrl:      os.Getenv("USER_SERVICE_BASE_URL"),
		ShopServiceUrl:      os.Getenv("SHOP_SERVICE_BASE_URL"),
		ProductServiceUrl:   os.Getenv("PRODUCT_SERVICE_BASE_URL"),
		WarehouseServiceUrl: os.Getenv("WAREHOUSE_SERVICE_BASE_URL"),
		OrderServiceUrl:     os.Getenv("ORDER_SERVICE_BASE_URL"),
	}

	return config
}
