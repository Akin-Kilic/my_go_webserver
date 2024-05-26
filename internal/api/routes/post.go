package routes

import (
	"my_go_webserver/pkg/domains/post"
	"my_go_webserver/pkg/models"

	"github.com/labstack/echo/v4"
)

func DeviceRoutes(e *echo.Group, service post.Service) {
	e.GET("/posts", GetAllPosts(service))
	e.POST("/post", CreatePost(service))
	e.GET("/post/:id", GetPost(service))
	e.PUT("/post/:id", UpdatePost(service))
	e.DELETE("/post/:id", DeletePost(service))

}

func GetAllPosts(s post.Service) func(c echo.Context) error {
	return func(c echo.Context) error {
		posts, err := s.GetAllPosts()
		if err != nil {
			return c.JSON(400, err.Error())
		}
		return c.JSON(200, posts)
	}
}

func GetPost(s post.Service) func(c echo.Context) error {
	return func(c echo.Context) error {
		id := c.Param("id")
		post, err := s.GetPost(id)
		if err != nil {
			return c.JSON(400, err.Error())
		}
		return c.JSON(200, post)
	}
}

func CreatePost(s post.Service) func(c echo.Context) error {
	return func(c echo.Context) error {
		var post models.Post
		err := c.Bind(&post)
		if err != nil {
			return c.JSON(400, err.Error())
		}
		err = s.CreatePost(post)
		if err != nil {
			return c.JSON(400, err.Error())
		}
		return c.JSON(200, "post created")
	}
}

func UpdatePost(s post.Service) func(c echo.Context) error {
	return func(c echo.Context) error {
		var post models.Post
		err := c.Bind(&post)
		if err != nil {
			return c.JSON(400, err.Error())
		}
		id := c.Param("id")
		err = s.UpdatePost(id, post)
		if err != nil {
			return c.JSON(400, err.Error())
		}
		return c.JSON(200, "posts updated")
	}
}

func DeletePost(s post.Service) func(c echo.Context) error {
	return func(c echo.Context) error {
		id := c.Param("id")
		err := s.DeletePost(id)
		if err != nil {
			return c.JSON(400, err.Error())
		}
		return c.JSON(200, "post deleted")
	}
}
