package inventoryHandler

import "github.com/larb26656/game-shop-tutorial/modules/inventory/inventoryUsecase"

type (
	inventoryGrpcHandler struct {
		inventoryUsecase inventoryUsecase.InventoryUsecaseService
	}
)

func NewInventoryGrpcHandler(inventoryUsecase inventoryUsecase.InventoryUsecaseService) *inventoryGrpcHandler {
	return &inventoryGrpcHandler{
		inventoryUsecase: inventoryUsecase,
	}
}
