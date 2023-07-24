package repository

import (
	"fmt"
	"player-service/src/helpers"
	"player-service/src/models"
	"strconv"
	"strings"

	"github.com/google/uuid"
)

type UserRepo interface {
	CreateUser(req models.CreateUser) error
	Login(username string) (*models.User, error)
	GetUserByUuid(uuid string) (*models.User, error)
	FindAllUser(req models.ListRequest) (paging models.JsonResponse, res []models.ResponseUserScanner, err error)
}

func (u *PlayerRepository) CreateUser(req models.CreateUser) error {
	tx := u.db.Begin()
	defer tx.Rollback()

	createUser := &models.User{
		Email:    req.Email,
		Username: req.Username,
		Password: req.Password,
		Phone:    req.Phone,
	}

	createUser.Uuid = uuid.New().String()
	if err := tx.Debug().Create(&createUser).Error; err != nil {
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

func (r *PlayerRepository) Login(email string) (*models.User, error) {
	var result models.User
	if err := r.db.Debug().Model(&models.User{}).
		Where("email = ? and deleted_at is null", email).
		First(&result).
		Error; err != nil && err.Error() != "record not found" {
		return nil, err
	}
	return &result, nil
}

func (r *PlayerRepository) GetUserByUuid(uuid string) (*models.User, error) {
	var res models.User
	if err := r.db.Debug().
		Model(models.User{}).
		Where("\"user\".uuid = ?", uuid).
		Scan(&res).Error; err != nil {
		return nil, err
	}

	return &res, nil
}

func (r *PlayerRepository) GetAllUsers() ([]models.ResponseUser, error) {
	var res []models.ResponseUser
	if err := r.db.Debug().
		Model(models.User{}).
		Scan(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

func (r *PlayerRepository) FindAllUser(req models.ListRequest) (paging models.JsonResponse, res []models.ResponseUserScanner, err error) {
	query := `"user".uuid, "user".username, "user".email, "user".phone, "user".created_at, "user".updated_at, "user".deleted_at,
	user_wallet.wallet, user_wallet.bank_name, user_wallet.account_number, user_wallet.account_name`

	q := r.db.Debug().
		Model(models.User{}).
		Select(query).
		Joins(`LEFT JOIN user_wallet on user_wallet.user_uuid = "user".uuid`)

	for _, v := range req.Search {
		switch v.Field {
		case "created_at", "updated_at", "deleted_at":
			from_to := strings.Split(v.Value, "&&")
			from, _ := strconv.ParseInt(from_to[0], 10, 64)
			to, _ := strconv.ParseInt(from_to[1], 10, 64)

			q = q.Where("topup"+"."+v.Field+" >= ? AND topup"+"."+v.Field+" <= ?", from, to)
		case "username", "email":
			q = q.Where("\"user\"."+v.Field+" ILIKE ?", v.Value+"%")
		case "account_name", "bank_name":
			q = q.Where("user_wallet."+v.Field+"  = ?", v.Value)
		case "amount":
			q = q.Where("user_wallet.wallet = ?", v.Value)
		}
	}

	sortBy := []string{}

	switch true {
	case helpers.InArray(req.Sort.Field, "uuid", "register_at", "updated_at", "amount") && helpers.InArray(req.Sort.Value, "desc", "asc"):
		if req.Sort.Field == "amount" {
			req.Sort.Field = "user_wallet.wallet"
		}
		if req.Sort.Field == "register_at" {
			req.Sort.Field = "\"user\".created_at"
		}
		sortBy = append(sortBy, fmt.Sprintf(" %s %s", req.Sort.Field, req.Sort.Value))
	}

	if len(sortBy) != 0 {
		q = q.Order(strings.Join(sortBy, ", "))
	} else {
		q = q.Order("\"user\".created_at DESC")
	}
	if req.Page == 0 {
		req.Page = 1
	}
	if req.Limit == 0 {
		req.Limit = 20
	}
	offset := (req.Page - 1) * req.Limit
	if offset < 0 {
		offset = 0
	}

	q = q.Order(sortBy)
	q.Count(&paging.TotalObjs)
	q = q.Offset(offset)
	q = q.Limit(req.Limit)
	if err := q.Scan(&res).Error; err != nil {
		return paging, nil, err
	}

	paging.Page = req.Page
	paging.PerPage = req.Limit
	paging.TotalPage = int(paging.TotalObjs+int64(req.Limit)-1) / req.Limit
	if paging.TotalPage <= 0 {
		paging.TotalPage = 1
	}

	return paging, res, nil
}
