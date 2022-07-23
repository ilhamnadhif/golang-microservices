package main

import (
	"api-gateway/app"
	"api-gateway/config"
	"api-gateway/controller"
	"api-gateway/service"
	"context"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	initConfig := config.InitConfig()
	user, userClose := app.InitUserService(initConfig.Service[config.User])
	defer userClose()
	//auth, authClose := app.InitAuthService(initConfig.Service[config.Auth])
	//defer authClose()
	//product, userClose := app.InitUserService(initConfig.Service[config.Product])
	//defer userClose()

	userService := service.NewUserService(user)
	userController := controller.NewUserController(userService)

	e := echo.New()
	e.HTTPErrorHandler = app.CustomHTTPErrorHandler
	e.Use(middleware.Recover())

	userRouter := e.Group("/users")
	userRouter.GET("", userController.FindAll)
	userRouter.GET("/:userID", userController.FindOneByID)
	userRouter.POST("", userController.Create)
	userRouter.PUT("/:userID", userController.Update)
	userRouter.DELETE("/:userID", userController.Delete)

	go func() {
		if err := e.Start(initConfig.Server.HostPort); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal("shutting down the server")
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with a timeout of 10 seconds.
	// Use a buffered channel to avoid missing signals as recommended for signal.Notify
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}
