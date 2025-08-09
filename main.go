package main

import (
	"api-echo/internal/db"
	"api-echo/internal/handler"
	"api-echo/internal/repository"
	"api-echo/internal/service"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

    db, err := db.Connect()
    if err != nil {
        log.Fatal("Erro ao conectar com o banco:", err)
    }

    // Não precisa de defer dbConn.Close() com GORM, pois ele gerencia a conexão

    userRepo := repository.NewUserRepository(db)
    userService := service.NewUserService(userRepo)
    userHandler := handler.NewUserHandler(userService)

    e := echo.New()

    e.Use(middleware.Logger())
    e.Use(middleware.Recover())
    e.Use(middleware.CORS())

    setupRoutes(e, userHandler)

    e.Logger.Fatal(e.Start(":8080"))
}

func setupRoutes(e *echo.Echo, userHandler *handler.UserHandler) {

    api := e.Group("/api/v1")

    api.POST("/user", userHandler.CreateUser)
    api.GET("/user/:id", userHandler.FindById)
    api.GET("/users", userHandler.FindAll)
    api.PUT("/user/:id", userHandler.UpdateUser)
    api.DELETE("/user/:id", userHandler.DeleteById)

    e.GET("/teste", func(c echo.Context) error {
        return c.JSON(http.StatusOK, map[string]string{
            "status": "ok",
            "message": "API Echo funcionando!",
        })
    })

    for _, route := range e.Routes() {
        log.Printf("  %s %s", route.Method, route.Path)
    }
}