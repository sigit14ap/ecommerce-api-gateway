package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sigit14ap/api-gateway/internal/usecase"
)

type ShopHandler struct {
	shopUsecase *usecase.ShopUsecase
}

func NewShopHandler(shopUsecase *usecase.ShopUsecase) *ShopHandler {
	return &ShopHandler{shopUsecase: shopUsecase}
}

func (handler *ShopHandler) Register(context *gin.Context) {
	headers := context.Request.Header
	var payload map[string]interface{}
	if err := context.ShouldBindJSON(&payload); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	response, err := handler.shopUsecase.Register(headers, payload)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(response.StatusCode, response.Body)
}

func (handler *ShopHandler) Login(context *gin.Context) {
	headers := context.Request.Header
	var payload map[string]interface{}
	if err := context.ShouldBindJSON(&payload); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	response, err := handler.shopUsecase.Login(headers, payload)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(response.StatusCode, response.Body)
}

func (handler *ShopHandler) Me(context *gin.Context) {
	headers := context.Request.Header
	response, err := handler.shopUsecase.Me(headers)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(response.StatusCode, response.Body)
}
