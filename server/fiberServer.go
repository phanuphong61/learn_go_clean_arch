package server

import (
	"e-assets/config"
	"e-assets/middleware"

	"fmt"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type fiberServer struct {
	app *fiber.App
	db  *gorm.DB
	cfg *config.Config
}

func NewFiberServer(cfg *config.Config, db *gorm.DB) Server {
	return &fiberServer{
		app: fiber.New(),
		db:  db,
		cfg: cfg,
	}
}

func (s *fiberServer) Start() {

	// initialzie routers here

	fiber := fiber.New()
	fiber.Use(middleware.CorsMiddleware)
	fmt.Print(fiber.Listen(s.cfg.Server.Port))
}
