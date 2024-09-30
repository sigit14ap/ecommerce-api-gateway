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

type WarehouseHandler struct {
	warehouseUsecase *usecase.WarehouseUsecase
}

func NewWarehouseHandler(warehouseUsecase *usecase.WarehouseUsecase) *WarehouseHandler {
	return &WarehouseHandler{warehouseUsecase: warehouseUsecase}
}

func (handler *WarehouseHandler) GetAllWarehouses(context *gin.Context) {
	headers := context.Request.Header
	response, err := handler.warehouseUsecase.GetAllWarehouses(headers)
	if err != nil {
		logger.Error("Failed fetch warehouses", logrus.Fields{"status": response.StatusCode})
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	logger.Info("fetch warehouses", logrus.Fields{"status": response.StatusCode})
	context.JSON(response.StatusCode, response.Body)
}

func (handler *WarehouseHandler) SetWarehouseStatus(context *gin.Context) {
	headers := context.Request.Header
	id := context.Param("id")
	warehouseID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		logger.Error("invalid Warehouse ID", nil)
		helpers.ErrorResponse(context, http.StatusUnauthorized, "invalid Warehouse ID")
		return
	}

	response, err := handler.warehouseUsecase.SetWarehouseStatus(headers, warehouseID)
	if err != nil {
		logger.Error("Failed Set Warehouse Status", logrus.Fields{"status": response.StatusCode})
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	logger.Info("Set Warehouse Status", logrus.Fields{"status": response.StatusCode})
	context.JSON(response.StatusCode, response.Body)
}

func (handler *WarehouseHandler) GetStocksByWarehouseID(context *gin.Context) {
	headers := context.Request.Header
	id := context.Param("id")
	warehouseID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		logger.Error("invalid Warehouse ID", nil)
		helpers.ErrorResponse(context, http.StatusUnauthorized, "invalid Warehouse ID")
		return
	}

	response, err := handler.warehouseUsecase.GetStocksByWarehouseID(headers, warehouseID)
	if err != nil {
		logger.Error("Failed Get Stock Warehouse", logrus.Fields{"status": response.StatusCode})
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	logger.Info("Get Stock Warehouse", logrus.Fields{"status": response.StatusCode})
	context.JSON(response.StatusCode, response.Body)
}

func (handler *WarehouseHandler) SendStock(context *gin.Context) {
	headers := context.Request.Header
	var payload map[string]interface{}
	if err := context.ShouldBindJSON(&payload); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	response, err := handler.warehouseUsecase.SendStock(headers, payload)
	if err != nil {
		logger.Error("Failed Send Stock", logrus.Fields{"status": response.StatusCode})
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	logger.Info("Send Stock", logrus.Fields{"status": response.StatusCode})
	context.JSON(response.StatusCode, response.Body)
}

func (handler *WarehouseHandler) TransferStock(context *gin.Context) {
	headers := context.Request.Header
	var payload map[string]interface{}
	if err := context.ShouldBindJSON(&payload); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	response, err := handler.warehouseUsecase.TransferStock(headers, payload)
	if err != nil {
		logger.Error("Failed Transfer Stock", logrus.Fields{"status": response.StatusCode})
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	logger.Info("Transfer Stock", logrus.Fields{"status": response.StatusCode})
	context.JSON(response.StatusCode, response.Body)
}
