package handler

import (
	"net/http"
	"github.com/labstack/echo/v4"
)





func Health_check(c echo.Context) error{
	return c.String(http.StatusOK, "Welcome!")
}