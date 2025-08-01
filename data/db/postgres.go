package db

import (
	"fmt"
	"time"

	"github.com/aliblue2/khodro45/configs"
	"github.com/aliblue2/khodro45/pkg/logging"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dbClient *gorm.DB

func InitiateDb(cfg *configs.Config) error {
	var logger = logging.NewLogger(cfg)

	cnnDb := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", cfg.Postgres.Host, cfg.Postgres.Port, cfg.Postgres.User, cfg.Postgres.Password, cfg.Postgres.DbName, cfg.Postgres.SSLMode)

	dbClient, err := gorm.Open(postgres.Open(cnnDb))
	if err != nil {
		logger.Fatal(logging.Postgres, logging.Startup, err.Error(), nil)
		return err
	}

	db, _ := dbClient.DB()

	err = db.Ping()

	if err != nil {
		logger.Fatal(logging.Postgres, logging.Startup, err.Error(), nil)
		return err
	}

	db.SetMaxIdleConns(cfg.Postgres.MaxIdleConns)
	db.SetMaxOpenConns(cfg.Postgres.MaxOpenConns)
	db.SetConnMaxLifetime(cfg.Postgres.SetConnMaxLifetime * time.Minute)
	logger.Info(logging.Postgres, logging.Startup, "postgres client successfully connected and establisher", nil)
	return nil
}

func GetDb() *gorm.DB {
	return dbClient
}
func CloseDb() {
	db, _ := dbClient.DB()
	db.Close()
}
