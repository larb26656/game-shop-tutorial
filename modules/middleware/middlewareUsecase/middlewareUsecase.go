package middlewareUsecase

import "github.com/larb26656/game-shop-tutorial/modules/middleware/middlewareRepository"

type (
	MiddlewareUsecaseService interface {}

	middlewareUsecase struct {
		middlewareRepository middlewareRepository.MiddlewareRepositoryService
	}
)

func NewMiddlewareUsecase(middlewareRepository middlewareRepository.MiddlewareRepositoryService) MiddlewareUsecaseService {
	return &middlewareUsecase{middlewareRepository}
}