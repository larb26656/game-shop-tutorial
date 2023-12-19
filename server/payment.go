package server

import (
	"github.com/larb26656/game-shop-tutorial/modules/payment/paymentHandler"
	"github.com/larb26656/game-shop-tutorial/modules/payment/paymentRepository"
	"github.com/larb26656/game-shop-tutorial/modules/payment/paymentUsecase"
)

func (s *server) paymentService() {
	repo := paymentRepository.NewPaymentRepository(s.db)
	usecase := paymentUsecase.NewPaymentUsecase(repo)
	httpHandler := paymentHandler.NewPaymentHttpHandler(s.cfg, usecase)
	queueHandler := paymentHandler.NewPaymentQueueHandler(s.cfg, usecase)

	_ = httpHandler
	_ = queueHandler

	payment := s.app.Group("/payment_v1")

	// Health Check
	payment.GET("", s.healthCheckService)
}
