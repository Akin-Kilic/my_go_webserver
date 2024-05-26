package main

import (
	"log"
	"my_go_webserver/pkg/config"
	"my_go_webserver/pkg/db"
	"my_go_webserver/pkg/infrastructer/server"

	"github.com/labstack/echo/v4"
)

func main() {
	conf := config.ReadValue()
	db.GetConnection(conf.Database)
	e := echo.New()
	server.ImportPostRoute(e, db.Client())
	log.Println("Listening on Port:" + conf.Port)
	log.Fatal(e.Start(":" + conf.Port))
}
