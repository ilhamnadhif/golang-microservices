package controller

import "github.com/labstack/echo/v4"

type UserController interface {
	FindOneByID(c echo.Context) error
	FindAll(c echo.Context) error
	Create(c echo.Context) error
	Update(c echo.Context) error
	Delete(c echo.Context) error
}
