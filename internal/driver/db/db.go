package db

import (
	"errors"
	"fmt"
	constant "gh5-backend/pkg/constants"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	dbConnections map[string]*gorm.DB
)

func Init() {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		os.Getenv(constant.DB_HOST),
		os.Getenv(constant.DB_USER),
		os.Getenv(constant.DB_PASS),
		os.Getenv(constant.DB_NAME),
		os.Getenv(constant.DB_PORT),
		os.Getenv(constant.DB_SSLMODE),
		os.Getenv(constant.DB_TZ),
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic("failed init db")
	}

	dbConnections = make(map[string]*gorm.DB)
	dbConnections[constant.DB_GH5_BACKEND] = db
}

func GetConnection(name string) (*gorm.DB, error) {
	if dbConnections[name] == nil {
		return nil, errors.New("connection is undefined")
	}
	return dbConnections[name], nil
}
