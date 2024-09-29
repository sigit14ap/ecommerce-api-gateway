package usecase

import (
	"net/http"

	"github.com/sigit14ap/api-gateway/config"
	"github.com/sigit14ap/api-gateway/internal/repository/api"
)

type UserUsecase struct {
	client *api.ApiCLient
}

func NewUserUsecase(cfg *config.Config) *UserUsecase {
	return &UserUsecase{
		client: api.NewClient(cfg.UserServiceUrl+"/api/v1", cfg.AppSecret),
	}
}

func (s *UserUsecase) Register(headers http.Header, payload interface{}) (*api.ApiClientResponse, error) {
	return s.client.Post("/users/register", headers, payload)
}

func (s *UserUsecase) Login(headers http.Header, payload interface{}) (*api.ApiClientResponse, error) {
	return s.client.Post("/users/login", headers, payload)
}

func (s *UserUsecase) Me(headers http.Header) (*api.ApiClientResponse, error) {
	return s.client.Get("/users/me", headers)
}
