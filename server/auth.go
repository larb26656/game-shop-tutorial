package server

import (
	"github.com/larb26656/game-shop-tutorial/modules/auth/authHandler"
	"github.com/larb26656/game-shop-tutorial/modules/auth/authRepository"
	"github.com/larb26656/game-shop-tutorial/modules/auth/authUsecase"
)

func (s *server) authService() {
	repo := authRepository.NewAuthRepository(s.db)
	usecase := authUsecase.NewAuthUsecase(repo)
	httpHandler := authHandler.NewAuthHttpHandler(s.cfg, usecase)
	grpcHandler := authHandler.NewAuthGrpcHandler(usecase)

	_ = httpHandler
	_ = grpcHandler

	auth := s.app.Group("/auth_v1")

	// Health Check
	auth.GET("", s.healthCheckService)
}
