package db

import (
	"course_project/configs"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewGormDBByDSN(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, err
	}

	// cek schema db used
	var schema string
	db.Raw("SELECT CURRENT_SCHEMA()").Scan(&schema)
	log.Println("Current schema:", schema)

	return db, nil
}

func NewGormDB(cfg configs.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=%s",
		cfg.Database.Host,
		cfg.Database.User,
		cfg.Database.Pass,
		cfg.Database.Name,
		cfg.Database.Port,
		cfg.Database.SSLMode,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// cek schema db used
	var schema string
	db.Raw("SELECT CURRENT_SCHEMA()").Scan(&schema)
	log.Println("Current schema:", schema)

	return db, nil
}
