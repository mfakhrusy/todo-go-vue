package handlers

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

type H map[string]interface{}

// get tasks endpoint
func GetTask(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, "taaaasks")
	}
}

// put tasks endpoint
func PutTask(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusCreated, H{
			"created": 123,
		})
	}
}

// delete tasks endpoint
func DeleteTask(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))
		return c.JSON(http.StatusOK, H{
			"deleted": id,
		})
	}
}
