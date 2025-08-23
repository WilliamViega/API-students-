package main

import (
  "github.com/labstack/echo/v4"
  "github.com/labstack/echo/v4/middleware"
  "net/http"
)

func main() {
  // Echo instance
  e := echo.New()

  // Middleware
  e.Use(middleware.Logger())
  e.Use(middleware.Recover())

  // Routes
  e.GET("/students", getStudensts)

  // Start server
   e.Logger.Fatal(e.Start(":8080"))
}

// Handler
func getStudensts(c echo.Context) error {
  return c.String(http.StatusOK, "List of all students")
} 