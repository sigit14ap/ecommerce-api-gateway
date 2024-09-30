package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sigit14ap/api-gateway/internal/usecase"
)

type OrderHandler struct {
	orderUsecase *usecase.OrderUsecase
}

func NewOrderHandler(orderUsecase *usecase.OrderUsecase) *OrderHandler {
	return &OrderHandler{orderUsecase: orderUsecase}
}

func (handler *OrderHandler) Checkout(context *gin.Context) {
	headers := context.Request.Header
	var payload map[string]interface{}
	if err := context.ShouldBindJSON(&payload); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	response, err := handler.orderUsecase.Checkout(headers, payload)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(response.StatusCode, response.Body)
}
