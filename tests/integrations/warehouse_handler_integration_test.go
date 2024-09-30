package handler_test

import (
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/sigit14ap/api-gateway/config"
	delivery "github.com/sigit14ap/api-gateway/internal/delivery/http"
	"github.com/sigit14ap/api-gateway/internal/usecase"
	"github.com/stretchr/testify/assert"
)

var (
	cfg              *config.Config
	warehouseHandler *delivery.WarehouseHandler
	testServerURL    string
)

func setup() {
	if err := godotenv.Load("../../.env"); err != nil {
		log.Fatal("Error loading .env file")
	}

	cfg = &config.Config{
		WarehouseServiceUrl: os.Getenv("WAREHOUSE_SERVICE_BASE_URL"),
		AppSecret:           os.Getenv("APP_SECRET"),
	}

	warehouseUsecase := usecase.NewWarehouseUsecase(cfg)
	warehouseHandler = delivery.NewWarehouseHandler(warehouseUsecase)

	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.GET("/api/v1/warehouses", warehouseHandler.GetAllWarehouses)

	testServer := httptest.NewServer(router)
	testServerURL = testServer.URL
}

func TestGetAllWarehouses_Integration(t *testing.T) {
	setup()

	req, _ := http.NewRequest(http.MethodGet, testServerURL+"/api/v1/warehouses", nil)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var responseBody map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&responseBody); err != nil {
		t.Fatalf("Failed to decode response: %v", err)
	}

	assert.NotNil(t, responseBody)
}
