package routes

import (
	"go-api/controllers"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Init() *echo.Echo {
	e := echo.New()
	// Enable CORS middleware
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"}, // Replace with your allowed origin(s)
		AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodOptions},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "/family\n/kategori\n/kategori/id\n/hotkeys\n/hotkeys/id\n/options\n/options/id\n/entities\n/entities/id\n/rule\n/rule/id")
	})

	e.GET("/family", controllers.FetchAllFamily)
	e.GET("/kategori", controllers.FetchAllKategori)
	e.GET("/kategori/:id", controllers.FetchKategoriById)
	e.GET("/hotkeys", controllers.FetchAllHotkeys)
	e.GET("/hotkeys/:id", controllers.FetchHotkeysById)
	e.GET("/options", controllers.FetchAllOptions)
	e.GET("/options/:id", controllers.FetchOptionById)
	e.GET("/entities", controllers.FetchAllEntities)
	e.GET("/entities/:id", controllers.FetchEntitiesById)
	e.GET("/rule", controllers.FetchAllRule)
	e.GET("/rule/:id", controllers.FetchRuleById)
	return e
}
