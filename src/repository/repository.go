package repository

import (
	"gorm.io/gorm"
)

type PlayerRepository struct {
	db *gorm.DB
}

type RepoInterface interface {
	UserRepo
	UserWalletRepo
}

func NewRepo(db *gorm.DB) *PlayerRepository {
	return &PlayerRepository{db: db}
}
