package handler

import (
	"api-echo/internal/domain"
	"api-echo/internal/service"
	"database/sql"
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

	err := c.Bind(&newUser)
    if err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{
            "error": domain.ErrInvalidJSON.Error(),
        })
    }

    createdUser, err := h.service.CreateUser(&newUser)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{
            "error": domain.ErrInternalServer.Error(),
        })
    }
    return c.JSON(http.StatusCreated, createdUser)
}

func (h *UserHandler) FindById(c echo.Context) error {
  	id := c.Param("id")
	if id == "" {
        return c.JSON(http.StatusBadRequest, map[string]string {
            "error": domain.ErrUserIDRequired.Error(),
        })
    }

	user, err := h.service.FindById(id)
	if err != nil {
        if err == sql.ErrNoRows {
            return c.JSON(http.StatusNotFound, map[string]string{
                "error": domain.ErrUserNotFound.Error(),
            })
        }

    	return c.JSON(http.StatusInternalServerError, map[string]string{
        	"error": domain.ErrInternalServer.Error(),
        })
    }

	return c.JSON(http.StatusOK, user)
}

func (h *UserHandler) FindAll(c echo.Context) error {

	users, err := h.service.FindAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string {
			"error": domain.ErrInternalServer.Error(),
		})
	}

	return c.JSON(http.StatusOK, users)
}

func (h *UserHandler) UpdateUser(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return c.JSON(http.StatusBadRequest, map[string]string {
			"error": domain.ErrUserIDRequired.Error(),
		})
	}

	var user domain.User

	err := c.Bind(&user)
	if err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{
            "error": domain.ErrInvalidJSON.Error(),
        })
    }

	newUser, err := h.service.UpdateUser(id, &user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string {
			"error": domain.ErrInternalServer.Error(), 
		})
	}

	return c.JSON(http.StatusOK, newUser)
}

func (h *UserHandler) DeleteById(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return c.JSON(http.StatusBadRequest, map[string]string {
			"error": domain.ErrUserIDRequired.Error(),
		})
	}


	err := h.service.DeleteById(id)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{
            "error": domain.ErrInternalServer.Error(),
        })
    }

	return c.JSON(http.StatusOK, map[string]string {
    	"message": domain.MsgUserDeleted,
	})
}