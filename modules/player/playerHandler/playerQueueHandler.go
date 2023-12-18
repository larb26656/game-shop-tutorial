package playerHandler

import (
	"github.com/larb26656/game-shop-tutorial/config"
	"github.com/larb26656/game-shop-tutorial/modules/player/playerUsecase"
)

type (
	PlayerQueueHandlerService interface{}

	playerQueueHandler struct {
		cfg           *config.Config
		playerUsecase playerUsecase.PlayerUsecaseService
	}
)

func NewPlayerQueueHandler(cfg *config.Config, playerUsecase playerUsecase.PlayerUsecaseService) PlayerQueueHandlerService {
	return &playerQueueHandler{playerUsecase: playerUsecase}
}
