package handler

import (
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/sigit14ap/api-gateway/internal/repository/api"
	"github.com/sigit14ap/api-gateway/internal/usecase"
	"github.com/sigit14ap/api-gateway/pkg/logger"
	"github.com/sirupsen/logrus"
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

	responseChan := make(chan *api.ApiClientResponse)
	errChan := make(chan error)

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		response, err := handler.orderUsecase.Checkout(headers, payload)
		if err != nil {
			errChan <- err
			return
		}
		responseChan <- response
	}()

	go func() {
		wg.Wait()
		close(responseChan)
		close(errChan)
	}()

	select {
	case response := <-responseChan:
		logger.Info("Checkout Order", logrus.Fields{"status": response.StatusCode})
		context.JSON(response.StatusCode, response.Body)

	case err := <-errChan:
		logger.Error("Failed Checkout Order", nil)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
}
