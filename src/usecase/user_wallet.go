package usecase

import (
	"context"
	"errors"
	"player-service/src/models"
)

type UserWalletUsecase interface {
	CreateUserWallet(req models.CreateUserWallet) error
	TopupUserWallet(ctx context.Context, req models.TopupUserWallet, userUuid string) error
}

func (u *PlayerUsecase) CreateUserWallet(req models.CreateUserWallet) error {
	userWallet, err := u.repo.GetUserWalletByUserUuid(req.UserUuid)
	if err != nil {
		return err
	}
	if req.AccountNumber == userWallet.AccountNumber {
		return errors.New("account number already resgistered")
	}
	req.Wallet = 0
	err = u.repo.CreateUserWallet(req)
	if err != nil {
		return err
	}
	return nil
}

func (u PlayerUsecase) TopupUserWallet(ctx context.Context, req models.TopupUserWallet, userUuid string) error {
	user, err := u.repo.GetUserWalletByUserUuid(userUuid)
	if err != nil {
		return err
	}

	if user == nil {
		return errors.New("user not found")
	}

	req.Amount += user.Wallet
	err = u.repo.UpdateUserWallet(req, userUuid)
	if err != nil {
		return err
	}

	return nil
}
