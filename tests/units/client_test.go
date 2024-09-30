package api_test

import (
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/sigit14ap/api-gateway/internal/repository/api"
)

func startMockServer(mockResponse map[string]interface{}, statusCode int) *httptest.Server {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(statusCode)
		json.NewEncoder(w).Encode(mockResponse)
	})
	return httptest.NewServer(handler)
}

func TestApiClient_DoRequest(t *testing.T) {
	if err := godotenv.Load("../../.env"); err != nil {
		log.Fatal("Error loading .env file")
	}

	mockResponse := map[string]interface{}{"message": "created"}
	mockServer := startMockServer(mockResponse, http.StatusOK)
	defer mockServer.Close()

	client := api.NewClient(os.Getenv("WAREHOUSE_SERVICE_BASE_URL"), os.Getenv("APP_SECRET"))

	headers := http.Header{"Content-Type": []string{"application/json"}}

	response, err := client.DoRequest(http.MethodGet, "/api/v1/warehouses", headers, nil)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if response.StatusCode != http.StatusOK {
		t.Fatalf("Expected status code %d, got %d", http.StatusOK, response.StatusCode)
	}
}
