package config

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Postgre struct {
	//DB Configuration
	Username string
	Password string
	Port     string
	Address  string
	Database string
}

func GetDatabase(dsnMaster string) *gorm.DB {
	dsn := "host=localhost user=postgres password=myPassword dbname=database_player port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")
	return db
}
