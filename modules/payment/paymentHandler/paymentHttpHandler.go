package paymentHandler

import (
	"github.com/larb26656/game-shop-tutorial/config"
	"github.com/larb26656/game-shop-tutorial/modules/payment/paymentUsecase"
)

type (
	PaymentHttpHandlerService interface{}

	paymentHttpHandler struct {
		cfg            *config.Config
		paymentUsecase paymentUsecase.PaymentUsecaseService
	}
)

func NewPaymentHttpHandler(cfg *config.Config, paymentUsecase paymentUsecase.PaymentUsecaseService) PaymentHttpHandlerService {
	return &paymentHttpHandler{
		cfg:            cfg,
		paymentUsecase: paymentUsecase,
	}
}
