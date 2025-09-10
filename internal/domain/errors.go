package domain

import "errors"

// User errors
var (
	ErrUserNotFound = errors.New("usuário não encontrado")
	ErrUserAlreadyExists = errors.New("usuário já existe")
	ErrInvalidUserData  = errors.New("dados do usuário inválidos")
	ErrUserIDRequired   = errors.New("ID do usuário é obrigatório")
)

// General errors
var (
    ErrInvalidJSON = errors.New("JSON inválido")
    ErrInternalServer = errors.New("erro interno do servidor")
)

// HTTP
var (
	MsgUserCreated = "Usuário criado com sucesso"
    MsgUserUpdated = "Usuário atualizado com sucesso"
    MsgUserDeleted = "Usuário deletado com sucesso"
)