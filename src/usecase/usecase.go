package usecase

import "player-service/src/repository"

type PlayerUsecase struct {
	repo repository.RepoInterface
}

type UsecaseInterface interface {
	UserUsecase
	UserWalletUsecase
}

func NewUsecase(repo repository.RepoInterface) UsecaseInterface {
	return &PlayerUsecase{
		repo: repo,
	}
}
