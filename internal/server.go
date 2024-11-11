package internal

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	c "github.com/meles-z/test/configs"
	db "github.com/meles-z/test/internal/db_utils"
	"gorm.io/gorm"
)

type IServer interface {
	Start() error
}

type Server struct {
	DB *gorm.DB

	configs c.Config
}

func NewServer(cfg c.Config) IServer {
	mainDb, err := db.Init(&cfg.Db)
	if err != nil {
		log.Fatalf("DB INIT ERROR: %v", err.Error())
	}
	return &Server{DB: mainDb, configs: cfg}
}

func (s *Server) Start() error {
	e:=echo.New()
	fmt.Printf("Starting server... exposed at: %d", s.configs.Server.Port)
	return e.Start(fmt.Sprintf("%s:%d", "0.0.0.0", 80))
	
}
