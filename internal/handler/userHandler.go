package handler

import (
	"api-echo/internal/domain"
	"api-echo/internal/dto"
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

	var inputDto dto.InputDto

	err := c.Bind(&inputDto)
    if err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{
            "error": domain.ErrInvalidJSON.Error(),
        })
    }

	domainUser := dto.ToDomain(inputDto)

    createdUser, err := h.service.CreateUser(domainUser)
    if err != nil {
		if err == domain.ErrUserAlreadyExists {
			return c.JSON(http.StatusConflict, map[string]string{
           	 	"error": domain.ErrUserAlreadyExists.Error(),
       		 })	
		}

        return c.JSON(http.StatusInternalServerError, map[string]string{
            "error": domain.ErrInternalServer.Error(),
        })
    }

	outputDto := dto.ToResponse(createdUser)

    return c.JSON(http.StatusCreated, outputDto)
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

	output := dto.ToResponse(user) 

	return c.JSON(http.StatusOK, output)
}

func (h *UserHandler) FindAll(c echo.Context) error {

	users, err := h.service.FindAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string {
			"error": domain.ErrInternalServer.Error(),
		})
	}

	var outputDtos []*dto.OutputDto
    for _, user := range users {
        outputDtos = append(outputDtos, dto.ToResponse(user))
    }
    
    return c.JSON(http.StatusOK, outputDtos)
}

func (h *UserHandler) UpdateUser(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return c.JSON(http.StatusBadRequest, map[string]string {
			"error": domain.ErrUserIDRequired.Error(),
		})
	}

	var inputDto dto.InputDto

	err := c.Bind(&inputDto)
	if err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{
            "error": domain.ErrInvalidJSON.Error(),
        })
    }

	user := dto.ToDomain(inputDto)

	domainUser, err := h.service.UpdateUser(id, user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string {
			"error": domain.ErrInternalServer.Error(), 
		})
	}

	outputDto := dto.ToResponse(domainUser)

	return c.JSON(http.StatusOK, outputDto)
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