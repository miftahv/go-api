package controllers

import (
	"go-api/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func FetchAllOptions(c echo.Context) error{
	result, err := models.FetchAllOptions()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()} )
	}

	return c.JSON(http.StatusOK, result)
	

}

func FetchOptionById(c echo.Context) error{
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid ID"})
	}
	result, err := models.FetchOptionById(id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()} )
	}

	return c.JSON(http.StatusOK, result)
	

}