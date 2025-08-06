package main

import (
	"api-echo/internal/handler"
	"api-echo/internal/repository"
	"api-echo/internal/service"
	"database/sql"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	db, err := sql.Open("postgres", "conexão")
    if err != nil {
        panic(err)
    }
    defer db.Close()

	r := repository.NewUserRepository(db)
	s := service.NewService(r)
	h := handler.NewUserHandler(s)

	e := echo.New()

	// e.Use(middleware.Logger())

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}, time=${latency_human}\n",
	}))


	// testando rotas

	e.POST("/user", h.CreateUser)

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"message": "Hello, World!",
		})
	})

	e.GET("/home", func(c echo.Context) error {
		return c.String(http.StatusOK, "Pagina home")
	})

	
	e.Logger.Fatal(e.Start(":8080"))
}
