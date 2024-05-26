package server

import (
	"my_go_webserver/internal/api/routes"
	"my_go_webserver/pkg/domains/post"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func ImportPostRoute(e *echo.Echo, db *gorm.DB) {
	repo := post.NewRepo(db)
	service := post.NewService(repo)
	post := e.Group("/api")
	routes.DeviceRoutes(post, service)
}
