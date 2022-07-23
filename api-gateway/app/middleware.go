package app

import (
	"api-gateway/model/dto"
	"github.com/labstack/echo/v4"
	"net/http"
)

func CustomHTTPErrorHandler(err error, c echo.Context) {
	data, ok := err.(*echo.HTTPError)
	if !ok {
		data = echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	c.JSON(data.Code, dto.WebResponse{
		Code:   data.Code,
		Status: http.StatusText(data.Code),
		Data:   data.Message,
	})

}
