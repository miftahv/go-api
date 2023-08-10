package controllers

import (
	"go-api/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func FetchAllHotkeys(c echo.Context) error{
	result, err := models.FetchAllHotkeys()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()} )
	}

	return c.JSON(http.StatusOK, result)
	

}
func FetchHotkeysById(c echo.Context) error{
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid ID"})
	}
	result, err := models.FetchHotkeysById(id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()} )
	}

	return c.JSON(http.StatusOK, result)
	

}