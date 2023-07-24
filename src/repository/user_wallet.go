package repository

import (
	"player-service/src/models"

	"github.com/google/uuid"
)

type UserWalletRepo interface {
	CreateUserWallet(req models.CreateUserWallet) error
	GetUserWalletByUserUuid(user_uuid string) (*models.UserWallet, error)
	UpdateUserWallet(req models.TopupUserWallet, userUuid string) error
	GetallUserWalletsByUserUuid(user_uuid string) ([]models.UserWalletResponse, error)
}

func (u *PlayerRepository) CreateUserWallet(req models.CreateUserWallet) error {
	tx := u.db.Begin()
	defer tx.Rollback()

	createUserWallet := &models.UserWallet{
		AccountNumber: req.AccountNumber,
		Wallet:        req.Wallet,
		UserUuid:      req.UserUuid,
		BankName:      req.BankName,
		AccountName:   req.AccountName,
	}

	createUserWallet.Uuid = uuid.New().String()
	if err := tx.Debug().Create(&createUserWallet).Error; err != nil {
		// return any error will rollback
		return err
	}
	// res.Uuid = createUser.Uuid
	err := tx.Commit().Error
	if err != nil {
		return err
	}

	return nil
}

func (u *PlayerRepository) UpdateUserWallet(req models.TopupUserWallet, userUuid string) error {
	err := u.db.Debug().Model(models.UserWallet{}).Where("user_uuid =?", userUuid).Update("wallet", req.Amount).Error
	if err != nil {
		return err
	}
	return nil
}

// func (u *User) UpdateUserByID(ID string, req RequestCreateUser) error {
// 	tx := u.db.Begin()
// 	defer tx.Rollback()

// 	err := tx.Debug().Table("user").Where("uuid = ?", ID).Updates(map[string]interface{}{
// 		"username":         req.Username,
// 		"fullname":         req.Fullname,
// 		"business_name":    req.BusinessName,
// 		"business_type_id": req.BusinessTypeId,
// 		"postal_code_id":   req.PostalCodeId,
// 		"updated_at":       time.Now(),
// 	}).Error
// 	if err != nil {
// 		return err
// 	}

// 	err = tx.Commit().Error
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

func (r *PlayerRepository) GetUserWalletByUserUuid(user_uuid string) (*models.UserWallet, error) {
	var res models.UserWallet
	if err := r.db.Debug().
		Model(models.UserWallet{}).
		Where("user_wallet.user_uuid = ?", user_uuid).
		Scan(&res).Error; err != nil {
		return nil, err
	}

	return &res, nil
}

func (r *PlayerRepository) GetallUserWalletsByUserUuid(user_uuid string) ([]models.UserWalletResponse, error) {
	var res []models.UserWalletResponse
	if err := r.db.Debug().
		Model(models.UserWallet{}).
		Where("user_wallet.user_uuid = ?", user_uuid).
		Scan(&res).Error; err != nil {
		return nil, err
	}

	return res, nil
}
