package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sigit14ap/api-gateway/helpers"
	"github.com/sigit14ap/api-gateway/internal/usecase"
	"github.com/sigit14ap/api-gateway/pkg/logger"
	"github.com/sirupsen/logrus"
)

type ProductHandler struct {
	productUsecase *usecase.ProductUsecase
}

func NewProductHandler(productUsecase *usecase.ProductUsecase) *ProductHandler {
	return &ProductHandler{productUsecase: productUsecase}
}

func (handler *ProductHandler) GetAllProductsWithStock(context *gin.Context) {
	headers := context.Request.Header
	response, err := handler.productUsecase.GetAllProductsWithStock(headers)
	if err != nil {
		logger.Error("Failed Get All Product Stock", logrus.Fields{"status": response.StatusCode})
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	logger.Info("Get All Product Stock", logrus.Fields{"status": response.StatusCode})
	context.JSON(response.StatusCode, response.Body)
}

func (handler *ProductHandler) GetAllByShopID(context *gin.Context) {
	headers := context.Request.Header
	response, err := handler.productUsecase.GetAllByShopID(headers)
	if err != nil {
		logger.Error("Failed Get All By Shop", logrus.Fields{"status": response.StatusCode})
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	logger.Info("Get All By Shop", logrus.Fields{"status": response.StatusCode})
	context.JSON(response.StatusCode, response.Body)
}

func (handler *ProductHandler) Create(context *gin.Context) {
	headers := context.Request.Header
	var payload map[string]interface{}
	if err := context.ShouldBindJSON(&payload); err != nil {
		logger.Error("Failed Create Product Invalid request payload", nil)
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	response, err := handler.productUsecase.Create(headers, payload)
	if err != nil {
		logger.Error("Failed Create Product", logrus.Fields{"status": response.StatusCode})
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	logger.Info("Create Product", logrus.Fields{"status": response.StatusCode})
	context.JSON(response.StatusCode, response.Body)
}

func (handler *ProductHandler) GetByIDAndShopID(context *gin.Context) {
	headers := context.Request.Header
	id := context.Param("id")
	productID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		logger.Error("Failed Get Product By ID : invalid Product ID", nil)
		helpers.ErrorResponse(context, http.StatusUnauthorized, "invalid Product ID")
		return
	}

	response, err := handler.productUsecase.GetByIDAndShopID(headers, productID)
	if err != nil {
		logger.Error("Failed Get Product By ID", logrus.Fields{"status": response.StatusCode})
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	logger.Info("Get Product By ID", logrus.Fields{"status": response.StatusCode})
	context.JSON(response.StatusCode, response.Body)
}

func (handler *ProductHandler) Update(context *gin.Context) {
	headers := context.Request.Header

	id := context.Param("id")
	productID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		logger.Error("Failed Update Product : invalid Product ID", nil)
		helpers.ErrorResponse(context, http.StatusUnauthorized, "invalid Product ID")
		return
	}

	var payload map[string]interface{}
	if err := context.ShouldBindJSON(&payload); err != nil {
		logger.Error("Failed Update Product : Invalid request payload", nil)
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	response, err := handler.productUsecase.Update(headers, payload, productID)
	if err != nil {
		logger.Error("Failed Update Product", logrus.Fields{"status": response.StatusCode})
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	logger.Info("Update Product", logrus.Fields{"status": response.StatusCode})
	context.JSON(response.StatusCode, response.Body)
}

func (handler *ProductHandler) Delete(context *gin.Context) {
	headers := context.Request.Header

	id := context.Param("id")
	productID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		logger.Error("Failed Delete Product : invalid Product ID", nil)
		helpers.ErrorResponse(context, http.StatusUnauthorized, "invalid Product ID")
		return
	}

	response, err := handler.productUsecase.Delete(headers, productID)
	if err != nil {
		logger.Error("Failed Delete Product", logrus.Fields{"status": response.StatusCode})
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	logger.Info("Delete Product", logrus.Fields{"status": response.StatusCode})
	context.JSON(response.StatusCode, response.Body)
}
