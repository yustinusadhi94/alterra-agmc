package http

import (
	"day-7-revision/internal/app/auth"
	"day-7-revision/internal/app/user"
	"day-7-revision/internal/factory"
	"github.com/labstack/echo/v4"
)

func NewHttp(e *echo.Echo, f *factory.Factory) {
	user.NewHandler(f).Route(e.Group("/v1/users"))
	auth.NewHandler(f).Route(e.Group("/v1/login"))
}
