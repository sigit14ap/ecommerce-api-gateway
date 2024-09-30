package usecase

import (
	"net/http"
	"strconv"

	"github.com/sigit14ap/api-gateway/config"
	"github.com/sigit14ap/api-gateway/internal/repository/api"
)

type WarehouseUsecase struct {
	client *api.ApiCLient
}

func NewWarehouseUsecase(cfg *config.Config) *WarehouseUsecase {
	return &WarehouseUsecase{
		client: api.NewClient(cfg.WarehouseServiceUrl+"/api/v1", cfg.AppSecret),
	}
}

func (s *WarehouseUsecase) GetAllWarehouses(headers http.Header) (*api.ApiClientResponse, error) {
	return s.client.Get("/warehouses", headers)
}

func (s *WarehouseUsecase) SetWarehouseStatus(headers http.Header, id uint64) (*api.ApiClientResponse, error) {
	return s.client.Patch("/warehouses/"+strconv.FormatUint(id, 10)+"/status", headers, nil)
}

func (s *WarehouseUsecase) GetStocksByWarehouseID(headers http.Header, id uint64) (*api.ApiClientResponse, error) {
	return s.client.Get("/stocks/warehouse/"+strconv.FormatUint(id, 10), headers)
}

func (s *WarehouseUsecase) SendStock(headers http.Header, payload interface{}) (*api.ApiClientResponse, error) {
	return s.client.Post("/stocks/send-stock", headers, payload)
}

func (s *WarehouseUsecase) TransferStock(headers http.Header, payload interface{}) (*api.ApiClientResponse, error) {
	return s.client.Post("/stocks/transfer-stock", headers, payload)
}
