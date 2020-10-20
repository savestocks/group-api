package controller

import (
	"encoding/base64"
	"net/http"
	"os"
	"strings"

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

	g.Use(BasicAuth)

	g.OPTIONS("/group",getDefaultOptions)
	g.GET("/group", GetGroupList)
	g.GET("/group/:id", GetGroupByID)
	g.POST("/group", SaveGroup)
	g.PUT("/group/:id", UpdateGroup)
	g.DELETE("/group/:id", DeleteGroup)
	g.GET("/health", CheckHealth)
	g.GET("/info", GetInfo)
}

func getDefaultOptions(c echo.Context) error {
	return c.JSON(http.StatusNoContent,nil)
}

// BasicAuth is the middleware function to enabled options
func BasicAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if c.Request().Method == echo.OPTIONS || c.Request().Method == echo.HEAD {
			next(c)
			return nil
		}
        auth := strings.SplitN(c.Request().Header.Get("Authorization"), " ", 2)

        if len(auth) != 2 || auth[0] != "Basic" {
            return nil
        }

        payload, _ := base64.StdEncoding.DecodeString(auth[1])
        pair := strings.SplitN(string(payload), ":", 2)

        if len(pair) == 2 && pair[0] == os.Getenv("apikey") && pair[1] == os.Getenv("apisecret") {
			next(c)
            return nil
		}
		return c.JSON(http.StatusUnauthorized, nil)
	}
}
