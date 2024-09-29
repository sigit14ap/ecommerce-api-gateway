package usecase

import (
	"net/http"

	"github.com/sigit14ap/api-gateway/config"
	"github.com/sigit14ap/api-gateway/internal/repository/api"
)

type ShopUsecase struct {
	client *api.ApiCLient
}

func NewShopUsecase(cfg *config.Config) *ShopUsecase {
	return &ShopUsecase{
		client: api.NewClient(cfg.ShopServiceUrl+"/api/v1", cfg.AppSecret),
	}
}

func (s *ShopUsecase) Register(headers http.Header, payload interface{}) (*api.ApiClientResponse, error) {
	return s.client.Post("/shop/register", headers, payload)
}

func (s *ShopUsecase) Login(headers http.Header, payload interface{}) (*api.ApiClientResponse, error) {
	return s.client.Post("/shop/login", headers, payload)
}

func (s *ShopUsecase) Me(headers http.Header) (*api.ApiClientResponse, error) {
	return s.client.Get("/shop/me", headers)
}
