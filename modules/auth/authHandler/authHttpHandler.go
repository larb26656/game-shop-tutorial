package authHandler

import (
	"github.com/larb26656/game-shop-tutorial/config"
	"github.com/larb26656/game-shop-tutorial/modules/auth/authUsecase"
)

type (
	AuthHttpandlerService interface {}

	authHttpHandler struct {
		cfg *config.Config
		authUsecase authUsecase.AuthUsecaseService
	}
)

func NewAuthHttpHandler(cfg *config.Config, authUsecase authUsecase.AuthUsecaseService) AuthHttpandlerService {
	return &authHttpHandler{cfg, authUsecase}
}