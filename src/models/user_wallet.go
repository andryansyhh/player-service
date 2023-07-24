package models

import (
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type UserWallet struct {
	Uuid          string      `gorm:"primaryKey;size:36;" json:"uuid"`
	CreatedAt     pq.NullTime `gorm:"created_at" json:"created_at"`
	UpdatedAt     pq.NullTime `gorm:"updated_at" json:"updated_at"`
	DeletedAt     pq.NullTime `gorm:"deleted_at" json:"deleted_at"`
	User          User        `gorm:"foreignKey:UserUuid"`
	UserUuid      string      `gorm:"user_uuid" json:"user_uuid"`
	Wallet        float64     `gorm:"wallet" json:"wallet"`
	AccountNumber string      `gorm:"account_number" json:"account_number"`
	BankName      string      `gorm:"bank_name" json:"bank_name"`
	AccountName   string      `gorm:"account_name" json:"account_name"`

	db *gorm.DB
}

type CreateUserWallet struct {
	UserUuid      string  `json:"user_uuid"`
	Wallet        float64 `json:"wallet"`
	AccountNumber string  `json:"account_number"`
	BankName      string  `json:"bank_name"`
	AccountName   string  `json:"account_name"`
}

type TopupUserWallet struct {
	Amount float64 `json:"amount"`
}

type UserWalletResponse struct {
	Uuid          string  `gorm:"primaryKey;size:36;" json:"uuid"`
	UserUuid      string  `gorm:"user_uuid" json:"user_uuid"`
	Wallet        float64 `gorm:"wallet" json:"wallet"`
	AccountNumber string  `gorm:"account_number" json:"account_number"`
	BankName      string  `gorm:"bank_name" json:"bank_name"`
	AccountName   string  `gorm:"account_name" json:"account_name"`
}

func (m *UserWallet) TableName() string {
	return "user_wallet"
}

func (m *UserWallet) SetDb(db *gorm.DB) {
	m.db = db
}
