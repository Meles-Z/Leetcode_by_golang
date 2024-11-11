package db

import (
	"fmt"

	c "github.com/meles-z/test/configs"
	"github.com/meles-z/test/internal/entities"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init(cfg *c.DatabaseConfigration) (*gorm.DB, error) {
    dns := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
        cfg.Host, cfg.Port, cfg.Username, cfg.Name, cfg.Password)
    conDigDB, err := gorm.Open(postgres.Open(dns), &gorm.Config{})
    if err != nil {
        return nil, err
    }
    DB = conDigDB

    // AutoMigrate entities
    err = conDigDB.AutoMigrate(&entities.User{})
    if err != nil {
        return nil, err
    }

    fmt.Println("Database is connected successfully")
    return conDigDB, nil
}
