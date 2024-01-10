package controllers

import (
	"go-api/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func FetchAllFamily(c echo.Context) error{
	result, err := models.FetchAllFamily()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()} )
	}

	return c.JSON(http.StatusOK, result)
}
