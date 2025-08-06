package handler

import (
	"api-echo/internal/domain"
	"api-echo/internal/service"
	"database/sql"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	service *service.UserService
}

func NewUserHandler(service *service.UserService) *UserHandler {
	return &UserHandler{service: service}
}

func (h *UserHandler) CreateUser(c echo.Context) error {
    var newUser domain.User

	err := c.Bind(&newUser) // le a requisção e atribui a newUser
    if err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{
            "error": "JSON inválido",
        })
    }

    createdUser, err := h.service.CreateUser(&newUser)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{
            "error": "Erro interno no servidor ao criar usuário",
        })
    }
    return c.JSON(http.StatusCreated, createdUser)
}

func (h *UserHandler) FindById(c echo.Context) error {
  	id := c.Param("id")
	if id == "" {
        return c.JSON(http.StatusBadRequest, map[string]string {
            "error": "id necessário",
        })
    }

	user, err := h.service.FindById(id)
	if err != sql.ErrNoRows {
        return c.JSON(http.StatusNotFound, map[string]string {
            "error": fmt.Sprintf("usuario com id: %s não encontrado", id),
        })
    }
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, user)
}

func (h *UserHandler) FindAll(c echo.Context) error {

	users, err := h.service.FindAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string {
			"error": "Erro interno no servidor",
		})
	}

	return c.JSON(http.StatusOK, users)
}

func (h *UserHandler) UpdateUser(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return c.JSON(http.StatusBadRequest, map[string]string {
			"error": "id necessário",
		})
	}

	var user domain.User

	err := c.Bind(&user)
	if err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{
            "error": "JSON inválido",
        })
    }

	newUser, err := h.service.UpdateUser(id, &user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string {
			"error": "Erro na atualizalão do usuário", 
		})
	}

	return c.JSON(http.StatusOK, newUser)
}

func (h *UserHandler) DeleteById(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return c.JSON(http.StatusBadRequest, map[string]string {
			"error": "id necessário",
		})
	}


	err := h.service.DeleteById(id)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{
            "error": "Erro ao deletar usuário",
        })
    }

	return c.JSON(http.StatusOK, map[string]string {
    	"message": "Usuário deletado com sucesso",
	})
}