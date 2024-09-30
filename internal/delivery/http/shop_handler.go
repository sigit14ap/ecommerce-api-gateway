package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sigit14ap/api-gateway/internal/usecase"
	"github.com/sigit14ap/api-gateway/pkg/logger"
	"github.com/sirupsen/logrus"
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
		logger.Error("Failed Register Shop : Invalid request payload", nil)
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	response, err := handler.shopUsecase.Register(headers, payload)
	if err != nil {
		logger.Error("Failed Register Shop", logrus.Fields{"status": response.StatusCode})
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	logger.Info("Register Shop", logrus.Fields{"status": response.StatusCode})
	context.JSON(response.StatusCode, response.Body)
}

func (handler *ShopHandler) Login(context *gin.Context) {
	headers := context.Request.Header
	var payload map[string]interface{}
	if err := context.ShouldBindJSON(&payload); err != nil {
		logger.Error("Failed Login Shop : Invalid request payload", nil)
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	response, err := handler.shopUsecase.Login(headers, payload)
	if err != nil {
		logger.Error("Failed Login Shop", logrus.Fields{"status": response.StatusCode})
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	logger.Info("Login Shop", logrus.Fields{"status": response.StatusCode})
	context.JSON(response.StatusCode, response.Body)
}

func (handler *ShopHandler) Me(context *gin.Context) {
	headers := context.Request.Header
	response, err := handler.shopUsecase.Me(headers)
	if err != nil {
		logger.Error("Failed Data Shop", logrus.Fields{"status": response.StatusCode})
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	logger.Info("Data Shop", logrus.Fields{"status": response.StatusCode})
	context.JSON(response.StatusCode, response.Body)
}
