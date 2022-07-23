package controller

import (
	"api-gateway/model/dto"
	"api-gateway/service"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type userControllerImpl struct {
	UserService service.UserService
}

func NewUserController(userService service.UserService) UserController {
	return &userControllerImpl{
		UserService: userService,
	}
}

func (controller *userControllerImpl) FindOneByID(c echo.Context) error {
	ctx := c.Request().Context()
	userID, err := strconv.Atoi(c.Param("userID"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	user, err := controller.UserService.FindOneByID(ctx, userID)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, dto.WebResponseSuccess(user))
}

func (controller *userControllerImpl) FindAll(c echo.Context) error {
	ctx := c.Request().Context()
	users, err := controller.UserService.FindAll(ctx)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, dto.WebResponseSuccess(users))
}

func (controller *userControllerImpl) Create(c echo.Context) error {
	ctx := c.Request().Context()
	var req dto.UserCreateReq
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	user, err := controller.UserService.Create(ctx, req)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, dto.WebResponseSuccess(user))
}

func (controller *userControllerImpl) Update(c echo.Context) error {
	ctx := c.Request().Context()
	userID, err := strconv.Atoi(c.Param("userID"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	var req dto.UserUpdateReq
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	req.ID = userID
	user, err := controller.UserService.Update(ctx, req)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, dto.WebResponseSuccess(user))
}

func (controller *userControllerImpl) Delete(c echo.Context) error {
	ctx := c.Request().Context()
	userID, err := strconv.Atoi(c.Param("userID"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	err = controller.UserService.Delete(ctx, userID)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, dto.WebResponseSuccess(nil))
}
