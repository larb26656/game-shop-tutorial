package inventoryUsecase

import "github.com/larb26656/game-shop-tutorial/modules/inventory/inventoryRepository"

type (
	InventoryUsecaseService interface{}

	inventoryUsecase struct {
		inventoryRepository inventoryRepository.InventoryRepositoryService
	}
)

func NewInventoryUsecase(inventoryRepository inventoryRepository.InventoryRepositoryService) InventoryUsecaseService {
	return &inventoryUsecase{
		inventoryRepository: inventoryRepository,
	}
}
