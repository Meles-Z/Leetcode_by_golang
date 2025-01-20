package database

import (
	"fmt"

	config "github.com/meles-z/go-learning/configs"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB(cfg config.DatabaseConfigration) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%d name=%s username=%s password=%s", cfg.Host, cfg.Port, cfg.Name, cfg.Username, cfg.Password)
	testDB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	DB = testDB
	err = testDB.AutoMigrate()
	if err != nil {
		return nil, fmt.Errorf("error to migrating db:%s", err)
	}
	return testDB, nil
}
