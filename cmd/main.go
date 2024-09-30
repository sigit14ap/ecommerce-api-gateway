package main

import (
	"github.com/sigit14ap/api-gateway/config"
	delivery "github.com/sigit14ap/api-gateway/internal/delivery/http"
	"github.com/sigit14ap/api-gateway/internal/router"
	"github.com/sigit14ap/api-gateway/internal/usecase"
)

func main() {
	// logger.InitLogger()
	// monitor.InitMetrics()

	cfg := config.LoadConfig()

	userUsecase := usecase.NewUserUsecase(cfg)
	userHandler := delivery.NewUserHandler(userUsecase)

	shopUsecase := usecase.NewShopUsecase(cfg)
	shopHandler := delivery.NewShopHandler(shopUsecase)

	productUsecase := usecase.NewProductUsecase(cfg)
	productHandler := delivery.NewProductHandler(productUsecase)

	warehouseUsecase := usecase.NewWarehouseUsecase(cfg)
	warehouseHandler := delivery.NewWarehouseHandler(warehouseUsecase)

	orderUsecase := usecase.NewOrderUsecase(cfg)
	orderHandler := delivery.NewOrderHandler(orderUsecase)

	router := router.NewRouter(userHandler, shopHandler, productHandler, warehouseHandler, orderHandler)

	router.Run(":" + cfg.AppPort)
}
