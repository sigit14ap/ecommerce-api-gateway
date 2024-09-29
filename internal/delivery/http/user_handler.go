package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sigit14ap/api-gateway/internal/usecase"
)

type UserHandler struct {
	userUsecase *usecase.UserUsecase
}

func NewUserHandler(userUsecase *usecase.UserUsecase) *UserHandler {
	return &UserHandler{userUsecase: userUsecase}
}

func (handler *UserHandler) Register(context *gin.Context) {
	headers := context.Request.Header
	var payload map[string]interface{}
	if err := context.ShouldBindJSON(&payload); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	response, err := handler.userUsecase.Register(headers, payload)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(response.StatusCode, response.Body)
}

func (handler *UserHandler) Login(context *gin.Context) {
	headers := context.Request.Header
	var payload map[string]interface{}
	if err := context.ShouldBindJSON(&payload); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	response, err := handler.userUsecase.Login(headers, payload)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(response.StatusCode, response.Body)
}

func (handler *UserHandler) Me(context *gin.Context) {
	headers := context.Request.Header
	response, err := handler.userUsecase.Me(headers)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(response.StatusCode, response.Body)
}
