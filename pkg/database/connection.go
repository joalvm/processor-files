package database

import (
	"fmt"

	"github.com/joalvm/processor-medias/pkg/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var formatDns = "host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC"

func New() (*gorm.DB, error) {
	dns := fmt.Sprintf(
		formatDns,
		utils.Env("DB_HOST"),
		utils.Env("DB_USERNAME"),
		utils.Env("DB_PASSWORD"),
		utils.Env("DB_NAME"),
		utils.Env("DB_PORT"),
	)

	return gorm.Open(postgres.Open(dns), &gorm.Config{})
}

func NewWithPostgresDb() (*gorm.DB, error) {
	dns := fmt.Sprintf(
		formatDns,
		utils.Env("DB_HOST"),
		utils.Env("DB_USERNAME"),
		utils.Env("DB_PASSWORD"),
		"postgres",
		utils.Env("DB_PORT"),
	)

	return gorm.Open(postgres.Open(dns), &gorm.Config{})
}

func NewWithDbName(dbName string) (*gorm.DB, error) {
	dns := fmt.Sprintf(
		formatDns,
		utils.Env("DB_HOST"),
		utils.Env("DB_USERNAME"),
		utils.Env("DB_PASSWORD"),
		dbName,
		utils.Env("DB_PORT"),
	)

	return gorm.Open(postgres.Open(dns), &gorm.Config{})
}

func CloseConnection(db *gorm.DB) {
	sqlDB, _ := db.DB()
	sqlDB.Close()
}
