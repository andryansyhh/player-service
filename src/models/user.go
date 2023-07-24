package models

import (
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type User struct {
	Uuid      string      `gorm:"primaryKey;size:36;" json:"uuid"`
	CreatedAt pq.NullTime `gorm:"created_at" json:"created_at"`
	UpdatedAt pq.NullTime `gorm:"updated_at" json:"updated_at"`
	DeletedAt pq.NullTime `gorm:"deleted_at" json:"deleted_at"`
	Username  string      `gorm:"username" json:"username"`
	Password  string      `gorm:"password" json:"password"`
	Email     string      `gorm:"email" json:"email"`
	Phone     string      `gorm:"phone" json:"phone"`

	db *gorm.DB
}

type CreateUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
}

type ResponseUser struct {
	Uuid       string               `gorm:"primaryKey;size:36;" json:"uuid"`
	CreatedAt  pq.NullTime          `gorm:"created_at" json:"created_at"`
	UpdatedAt  pq.NullTime          `gorm:"updated_at" json:"updated_at"`
	DeletedAt  pq.NullTime          `gorm:"deleted_at" json:"deleted_at"`
	Username   string               `gorm:"username" json:"username"`
	Email      string               `gorm:"email" json:"email"`
	Phone      string               `gorm:"phone" json:"phone"`
	UserWallet []UserWalletResponse `json:"user_wallet"`
	// Wallet        float64     `json:"wallet"`
	// AccountNumber string      `json:"account_number"`
	// BankName      string      `json:"bank_name"`
	// AccountName   string      `json:"account_name"`
}

type ResponseUserScanner struct {
	Uuid      string      `gorm:"primaryKey;size:36;" json:"uuid"`
	CreatedAt pq.NullTime `gorm:"created_at" json:"created_at"`
	UpdatedAt pq.NullTime `gorm:"updated_at" json:"updated_at"`
	DeletedAt pq.NullTime `gorm:"deleted_at" json:"deleted_at"`
	Username  string      `gorm:"username" json:"username"`
	Email     string      `gorm:"email" json:"email"`
	Phone     string      `gorm:"phone" json:"phone"`
}

type Login struct {
	Password string `json:"password"`
	Email    string `json:"email"`
}

type LoginResponse struct {
	Token    string `json:"token"`
	UserUuid string `json:"user_uuid"`
}

type Logout struct {
	Token string `json:"token"`
}

func (m *User) TableName() string {
	return "user"
}

func (m *User) SetDb(db *gorm.DB) {
	m.db = db
}
