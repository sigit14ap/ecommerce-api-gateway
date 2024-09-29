package usecase

import (
	"net/http"
	"strconv"

	"github.com/sigit14ap/api-gateway/config"
	"github.com/sigit14ap/api-gateway/internal/repository/api"
)

type ProductUsecase struct {
	client *api.ApiCLient
}

func NewProductUsecase(cfg *config.Config) *ProductUsecase {
	return &ProductUsecase{
		client: api.NewClient(cfg.ProductServiceUrl+"/api/v1", cfg.AppSecret),
	}
}

func (s *ProductUsecase) GetAllByShopID(headers http.Header) (*api.ApiClientResponse, error) {
	return s.client.Get("/shop/products", headers)
}

func (s *ProductUsecase) GetByIDAndShopID(headers http.Header, id uint64) (*api.ApiClientResponse, error) {
	return s.client.Get("/shop/products/"+strconv.FormatUint(id, 10), headers)
}

func (s *ProductUsecase) Create(headers http.Header, payload interface{}) (*api.ApiClientResponse, error) {
	return s.client.Post("/shop/products", headers, payload)
}

func (s *ProductUsecase) Update(headers http.Header, payload interface{}, id uint64) (*api.ApiClientResponse, error) {
	return s.client.Put("/shop/products/"+strconv.FormatUint(id, 10), headers, payload)
}

func (s *ProductUsecase) Delete(headers http.Header, id uint64) (*api.ApiClientResponse, error) {
	return s.client.Delete("/shop/products/"+strconv.FormatUint(id, 10), headers)
}

func (s *ProductUsecase) GetAllProductsWithStock(headers http.Header) (*api.ApiClientResponse, error) {
	return s.client.Get("/products", headers)
}
