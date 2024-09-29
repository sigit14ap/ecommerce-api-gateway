package router

import (
	delivery "github.com/sigit14ap/api-gateway/internal/delivery/http"

	"github.com/gin-gonic/gin"
)

func NewRouter(userHandler *delivery.UserHandler, shopHandler *delivery.ShopHandler, productHandler *delivery.ProductHandler) *gin.Engine {
	router := gin.New()
	v1 := router.Group("/api/v1")

	userGroup := v1.Group("/users")
	{
		userGroup.POST("register", userHandler.Register)
		userGroup.POST("login", userHandler.Login)
		userGroup.GET("me", userHandler.Me)
	}

	shopGroup := v1.Group("/shop")
	{
		shopGroup.POST("register", shopHandler.Register)
		shopGroup.POST("login", shopHandler.Login)
		shopGroup.GET("me", shopHandler.Me)
	}

	productsGroup := v1.Group("/products")
	{
		productsGroup.GET("/", productHandler.GetAllProductsWithStock)

		productsGroup.GET("/shop/products", productHandler.GetAllByShopID)
		productsGroup.POST("/shop/products", productHandler.Create)
		productsGroup.GET("/shop/products/:id", productHandler.GetByIDAndShopID)
		productsGroup.PUT("/shop/products/:id", productHandler.Update)
		productsGroup.DELETE("/shop/products/:id", productHandler.Delete)
	}

	return router
}
