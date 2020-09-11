package controller

import (
	"github.com/andersonlira/group-api/config"
	"github.com/labstack/echo-contrib/prometheus"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

//MapRoutes for the endpoints which the API listens for
func MapRoutes(e *echo.Echo) {
	g := e.Group("/group-api/v1")
	if config.Values.UsePrometheus {
		p := prometheus.NewPrometheus("echo", nil)
		p.Use(e)
	}
	g.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization, echo.HeaderContentType},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.POST, echo.DELETE, echo.OPTIONS},
	}))

	g.GET("/group", GetGroupList)
	g.GET("/group/:id", GetGroupByID)
	g.POST("/group", SaveGroup)
	g.PUT("/group/:id", UpdateGroup)
	g.DELETE("/group/:id", DeleteGroup)
	g.GET("/health", CheckHealth)
	g.GET("/info", GetInfo)
}
