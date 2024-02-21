package main

import (
	"e-assets/config"
	"e-assets/database"
	"e-assets/server"
)

func main() {

	cfg := config.GetConfig()

	db := database.NewMysqlDatabase(&cfg)

	server.NewFiberServer(&cfg, db.GetDb()).Start()

	// gin -p 9000 -a 9001 run main.go
}
