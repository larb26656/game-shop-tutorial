package middlewareHandler

import (
	"github.com/larb26656/game-shop-tutorial/config"
	"github.com/larb26656/game-shop-tutorial/modules/middleware/middlewareUsecase"
)

type (
	MiddlewareHandlerService interface {}

	middlewareHandler struct {
		cfg *config.Config
		middlewareUsecase middlewareUsecase.MiddlewareUsecaseService
	}
)

func NewMiddlewareHandler(cfg *config.Config, middlewareUsecase middlewareUsecase.MiddlewareUsecaseService) MiddlewareHandlerService {
	return &middlewareHandler{cfg, middlewareUsecase}
}