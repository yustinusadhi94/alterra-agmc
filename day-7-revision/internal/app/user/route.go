package user

import (
	middlewares "day-7-revision/internal/middleware"
	"github.com/labstack/echo/v4"
)

func (h *handler) Route(g *echo.Group) {
	g.GET("", h.GetUsers, middlewares.ValidateToken())
	g.GET("/:id", h.GetUserById, middlewares.ValidateToken())
	g.POST("", h.CreateUser)
	g.PUT("/:id", h.UpdateUser, middlewares.ValidateToken())
	g.DELETE("/:id", h.DeleteUser, middlewares.ValidateToken())
}
