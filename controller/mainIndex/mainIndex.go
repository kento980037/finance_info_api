package mainIndex

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"finance-info-api/scraping/mainIndex"
)

func GetEntries() echo.HandlerFunc {
	return func(c echo.Context) error {
		response := mainIndex.Scraping(c.Param("code"))

		return c.JSON(http.StatusOK, response)
	}
}
