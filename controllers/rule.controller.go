package controllers

import (
	"go-api/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func FetchAllRule(c echo.Context) error{
	result, err := models.FetchAllRule()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()} )
	}

	return c.JSON(http.StatusOK, result)
	

}
func FetchRuleById(c echo.Context) error{
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid ID"})
	}
	result, err := models.FetchRuleById(id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()} )
	}

	return c.JSON(http.StatusOK, result)
	

}