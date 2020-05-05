package user

import (
	"github.com/labstack/echo/v4"
)

type resp struct {
	Message string      `json:"message,omitempty"`
	Error   string      `json:"error,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

// Register node api
func Register(g *echo.Group) {
	g.Use(auth)
	g.Any("/add", addUser)
	g.Any("/disable", disableUser)
	g.Any("/list", listUser)
	g.Any("/get", getUser)
}
