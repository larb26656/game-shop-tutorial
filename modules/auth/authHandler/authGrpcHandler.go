package authHandler

import "github.com/larb26656/game-shop-tutorial/modules/auth/authUsecase"

type (
	authGrpcHandler struct {
		authUsecase authUsecase.AuthUsecaseService
	}
)

func NewAuthGrpcHandler(authUsecase authUsecase.AuthUsecaseService) authUsecase.AuthUsecaseService {
	return &authGrpcHandler{authUsecase}
}