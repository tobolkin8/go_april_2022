package main

import (
	"errors"
	"flag"
	"os"

	helpers "github.com/tobolkin8/helpers"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {

	var (
		listen = flag.String("listen", ":3000", "listen address of the application")
		env    = flag.String("env", "dev", "application environment")
	)
	flag.Parse()

	// set the flags into env vars
	os.Setenv("APP_ENV", *env)
	// application configuration
	router := echo.New()
	router.HTTPErrorHandler = errorHandler

	router.Use(middleware.Logger(), middleware.Recover())
	registerAuthHandlers(router, helpers.AuthService{})

	// V1 API endpoints
	//v1 := router.Group("/v1")
	//registerUserHandlers(v1, handler.UserService{repo})

	router.Start(*listen)
}

func registerAuthHandlers(router *echo.Echo, authService helpers.AuthService) {
	router.GET("/test", authService.Show)
}

// User related endpoints
// func registerUserHandlers(router *echo.Group, userService handler.UserService) {
// 	router.Get("/users/:id", userService.Show)
// 	router.Get("/users", userService.List)
// 	router.Post("/users", userService.Create)
// }

// errorHandler will catch all errors returned from handlers and convert them
// into a JSON response. If the error is not from the type HTTPError an internal
// server error will follow.
func errorHandler(err error, c echo.Context) {
	if httpError, ok := err.(helpers.HTTPError); ok {
		c.JSON(httpError.Code, httpError)
		return
	}
	c.Error(errors.New("internal server error"))
}
