package usecase

import (
	"context"
	"errors"
	"player-service/src/config"
	"player-service/src/helpers"
	"player-service/src/middleware"
	"player-service/src/models"
	"strings"

	"github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

type UserUsecase interface {
	CreateUser(req models.CreateUser) error
	Login(ctx context.Context, req models.Login) (*models.LoginResponse, error)
	Logout(ctx context.Context, req models.Logout) error
	GetUserByUuid(userUuid string) (*models.ResponseUser, error)
	GetAllUsers(req models.ListRequest) (*models.JsonResponse, error)
}

func (u *PlayerUsecase) CreateUser(req models.CreateUser) error {

	if req.Email == "" || req.Password == "" {
		return errors.New("email or password cannot be empty")
	}

	if req.Username == "" {
		return errors.New("fulllname cannot be empty")
	}

	check, err := u.repo.Login(req.Email)
	if err != nil {
		return err
	}
	if req.Email == check.Email {
		return errors.New("email already registered")
	}

	if req.Username == check.Username {
		return errors.New("username already registered")
	}

	if req.Phone == check.Phone {
		return errors.New("phone already registered")
	}

	isMatch := helpers.CheckPassword(req.Password)
	if !isMatch {

		return errors.New("error input password")
	}

	hash, err := helpers.NewHashPassword(req.Password)
	if err != nil {
		return errors.New("failed hashing password")
	}

	req.Password = hash
	err = u.repo.CreateUser(req)
	if err != nil {
		return err
	}
	return nil
}

func (u PlayerUsecase) Login(ctx context.Context, req models.Login) (*models.LoginResponse, error) {
	if !strings.Contains(req.Email, "@") {
		return nil, errors.New("must login with valid email")
	}

	user, err := u.repo.Login(req.Email)
	if err != nil {
		return nil, err
	}
	if user.Uuid == "" {
		return nil, errors.New("user not found")
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return nil, errors.New("invalid password")
	}

	// Generate JWT token
	token, err := middleware.GenerateJWT(user.Uuid)
	if err != nil {
		return nil, errors.New("failed to generate JWT token")
	}

	// save to redis
	config.SetRedisData(token, user.Uuid)

	return &models.LoginResponse{
		Token:    token,
		UserUuid: user.Uuid,
	}, nil
}

func (u PlayerUsecase) Logout(ctx context.Context, req models.Logout) error {
	config.DeleteRedisData(req.Token)
	return nil
}

func (u PlayerUsecase) GetUserByUuid(userUuid string) (*models.ResponseUser, error) {
	resUser, err := u.repo.GetUserByUuid(userUuid)
	if err != nil {
		return nil, errors.New("error get user data")
	}

	resUserWallet, err := u.repo.GetUserWalletByUserUuid(userUuid)
	if err != nil {
		return nil, errors.New("error get wallet data")
	}

	return &models.ResponseUser{
		Uuid:          resUser.Uuid,
		Email:         resUser.Email,
		Username:      resUser.Username,
		Phone:         resUser.Phone,
		CreatedAt:     pq.NullTime{Time: resUser.CreatedAt.Time},
		UpdatedAt:     pq.NullTime{Time: resUser.UpdatedAt.Time},
		DeletedAt:     pq.NullTime{Time: resUser.DeletedAt.Time},
		Wallet:        resUserWallet.Wallet,
		AccountNumber: resUserWallet.AccountNumber,
		BankName:      resUserWallet.BankName,
		AccountName:   resUserWallet.AccountName,
	}, nil
}

func (u PlayerUsecase) GetAllUsers(req models.ListRequest) (*models.JsonResponse, error) {
	paging, res, err := u.repo.FindAllTopup(req)
	if err != nil {
		return nil, err
	}

	data := []models.ResponseUser{}
	for _, res := range res {

		userData := models.ResponseUser{
			Uuid:          res.Uuid,
			Email:         res.Email,
			Username:      res.Username,
			Phone:         res.Phone,
			CreatedAt:     pq.NullTime{Time: res.CreatedAt.Time},
			UpdatedAt:     pq.NullTime{Time: res.UpdatedAt.Time},
			DeletedAt:     pq.NullTime{Time: res.DeletedAt.Time},
			Wallet:        res.Wallet,
			AccountNumber: res.AccountNumber,
			BankName:      res.BankName,
			AccountName:   res.AccountName,
		}

		data = append(data, userData)
	}

	resp := &models.JsonResponse{
		Page:      paging.Page,
		TotalPage: paging.TotalPage,
		TotalObjs: int64(paging.TotalObjs),
		PerPage:   paging.PerPage,
		Objs:      data,
	}
	return resp, nil
}
