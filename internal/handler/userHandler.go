package handler

import (
	"api-echo/internal/domain"
	"api-echo/internal/dto"
	"api-echo/internal/service"
	"database/sql"
	"log"
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

	log.Println("Handler: recebida requisição para criar usuário")

	var inputDto dto.InputDto

	err := c.Bind(&inputDto)
    if err != nil {
		log.Println("Handler: erro ao fazer bind do JSON")
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

		log.Println("handler: erro interno ao criar usuário")
        return c.JSON(http.StatusInternalServerError, map[string]string{
            "error": domain.ErrInternalServer.Error(),
        })
    }
	log.Println("handler: usuário criado com sucesso")

	outputDto := dto.ToResponse(createdUser)
    return c.JSON(http.StatusCreated, outputDto)
}

func (h *UserHandler) FindById(c echo.Context) error {
	log.Println("Handler: recebida requisição para buscar usuário por ID")

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
		log.Println("handler: erro interno ao buscar usuário")
    	return c.JSON(http.StatusInternalServerError, map[string]string{
        	"error": domain.ErrInternalServer.Error(),
        })
    }
	log.Println("Handler: usuário encontrado")

	output := dto.ToResponse(user) 
	return c.JSON(http.StatusOK, output)
}

func (h *UserHandler) FindAll(c echo.Context) error {
	log.Println("handler: recebida requisição para buscar todos os usuários")

	users, err := h.service.FindAll()
	if err != nil {
		log.Println("handler: erro ao buscar todos os usuarios")
		return c.JSON(http.StatusInternalServerError, map[string]string {
			"error": domain.ErrInternalServer.Error(),
		})
	}

	log.Println("handler: usuários encontrados")
	var outputDtos []*dto.OutputDto
    for _, user := range users {
        outputDtos = append(outputDtos, dto.ToResponse(user))
    }
    
    return c.JSON(http.StatusOK, outputDtos)
}

func (h *UserHandler) UpdateUser(c echo.Context) error {
	log.Println("handler: recebida requisição para atualizar usuário")

	id := c.Param("id")
	if id == "" {
		return c.JSON(http.StatusBadRequest, map[string]string {
			"error": domain.ErrUserIDRequired.Error(),
		})
	}

	var inputDto dto.InputDto

	err := c.Bind(&inputDto)
	if err != nil {
		log.Println("handler: erro ao fazer bind do JSON")
        return c.JSON(http.StatusBadRequest, map[string]string{
            "error": domain.ErrInvalidJSON.Error(),
        })
    }

	user := dto.ToDomain(inputDto)

	domainUser, err := h.service.UpdateUser(id, user)
	if err != nil {
		log.Println("handler: erro ao atualizar usuario")
		return c.JSON(http.StatusInternalServerError, map[string]string {
			"error": domain.ErrInternalServer.Error(), 
		})
	}

	log.Println("handler: usuário atualizado com sucesso")

	outputDto := dto.ToResponse(domainUser)
	return c.JSON(http.StatusOK, outputDto)
}

func (h *UserHandler) DeleteById(c echo.Context) error {
	log.Println("handler: recebida requisição para deletar usuário")

	id := c.Param("id")
	if id == "" {
		return c.JSON(http.StatusBadRequest, map[string]string {
			"error": domain.ErrUserIDRequired.Error(),
		})
	}

	err := h.service.DeleteById(id)
    if err != nil {
		log.Println("handler: erro ao deletar usuário")
        return c.JSON(http.StatusInternalServerError, map[string]string{
            "error": domain.ErrInternalServer.Error(),
        })
    }

	log.Println("handler: usuário deletado com sucesso")

	return c.JSON(http.StatusOK, map[string]string {
    	"message": domain.MsgUserDeleted,
	})
}