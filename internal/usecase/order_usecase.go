package usecase

import (
	"net/http"

	"github.com/sigit14ap/api-gateway/config"
	"github.com/sigit14ap/api-gateway/internal/repository/api"
)

type OrderUsecase struct {
	client *api.ApiCLient
}

func NewOrderUsecase(cfg *config.Config) *OrderUsecase {
	return &OrderUsecase{
		client: api.NewClient(cfg.OrderServiceUrl+"/api/v1", cfg.AppSecret),
	}
}

func (s *OrderUsecase) Checkout(headers http.Header, payload interface{}) (*api.ApiClientResponse, error) {
	return s.client.Post("/order/checkout", headers, payload)
}
