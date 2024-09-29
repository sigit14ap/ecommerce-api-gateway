package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sigit14ap/api-gateway/helpers"
	"github.com/sigit14ap/api-gateway/internal/usecase"
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
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(response.StatusCode, response.Body)
}

func (handler *ProductHandler) GetAllByShopID(context *gin.Context) {
	headers := context.Request.Header
	response, err := handler.productUsecase.GetAllByShopID(headers)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(response.StatusCode, response.Body)
}

func (handler *ProductHandler) Create(context *gin.Context) {
	headers := context.Request.Header
	var payload map[string]interface{}
	if err := context.ShouldBindJSON(&payload); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	response, err := handler.productUsecase.Create(headers, payload)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(response.StatusCode, response.Body)
}

func (handler *ProductHandler) GetByIDAndShopID(context *gin.Context) {
	headers := context.Request.Header
	id := context.Param("id")
	productID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		helpers.ErrorResponse(context, http.StatusUnauthorized, "invalid Product ID")
		return
	}

	response, err := handler.productUsecase.GetByIDAndShopID(headers, productID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(response.StatusCode, response.Body)
}

func (handler *ProductHandler) Update(context *gin.Context) {
	headers := context.Request.Header

	id := context.Param("id")
	productID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		helpers.ErrorResponse(context, http.StatusUnauthorized, "invalid Product ID")
		return
	}

	var payload map[string]interface{}
	if err := context.ShouldBindJSON(&payload); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	response, err := handler.productUsecase.Update(headers, payload, productID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(response.StatusCode, response.Body)
}

func (handler *ProductHandler) Delete(context *gin.Context) {
	headers := context.Request.Header

	id := context.Param("id")
	productID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		helpers.ErrorResponse(context, http.StatusUnauthorized, "invalid Product ID")
		return
	}

	response, err := handler.productUsecase.Delete(headers, productID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(response.StatusCode, response.Body)
}
