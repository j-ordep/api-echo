# API Echo - Gerenciador de Usuários

Uma API REST completa para gerenciamento de usuários implementada em Go, utilizando Clean Architecture, Echo framework e PostgreSQL.

## Descrição do Projeto

Esta API fornece operações CRUD (Create, Read, Update, Delete) para gerenciamento de usuários. O projeto foi desenvolvido seguindo os princípios da Clean Architecture, garantindo separação de responsabilidades, testabilidade e manutenibilidade do código.

### Características Principais

- **Clean Architecture**: Separação clara entre camadas (Domain, Repository, Service, Handler)
- **Echo Framework**: Framework web rápido e minimalista para Go
- **PostgreSQL**: Banco de dados relacional robusto
- **Gerenciamento de ponteiros**: Uso otimizado de ponteiros para performance
- **Tratamento de erros**: Sistema consistente de tratamento de erros
- **API RESTful**: Endpoints seguindo padrões REST

## Arquitetura

O projeto segue a Clean Architecture com as seguintes camadas:

```
api-echo/
├── internal/
│   ├── domain/          # Entidades de domínio e interfaces
│   │   ├── user.go      # Entidade User
│   │   └── repository.go # Interface UserRepository
│   ├── repository/      # Implementação do acesso a dados
│   │   └── userRepository.go
│   ├── service/         # Lógica de negócio
│   │   └── userService.go
│   └── handler/         # Controladores HTTP
│       └── userHandler.go
└── main.go             # Ponto de entrada da aplicação
```

### Camadas da Arquitetura

1. **Domain**: Define as entidades principais (User) e interfaces (UserRepository)
2. **Repository**: Implementa o acesso aos dados e persistência no PostgreSQL
3. **Service**: Contém a lógica de negócio e regras da aplicação
4. **Handler**: Gerencia as requisições HTTP e respostas

## Endpoints da API

| Método | Endpoint     | Descrição                    |
|--------|--------------|------------------------------|
| POST   | /user        | Criar novo usuário           |
| GET    | /user/:id    | Buscar usuário por ID        |
| GET    | /users       | Listar todos os usuários     |
| PUT    | /user/:id    | Atualizar usuário existente  |
| DELETE | /user/:id    | Remover usuário              |

## Estrutura da Entidade User

```go
type User struct {
    Id        string    `json:"id"`
    Name      string    `json:"name"`
    Email     string    `json:"email"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}
```

## Tecnologias Utilizadas

- **Go**: Linguagem de programação principal
- **Echo v4**: Framework web para APIs REST
- **PostgreSQL**: Sistema de gerenciamento de banco de dados
- **database/sql**: Driver padrão do Go para SQL
- **UUID**: Para geração de identificadores únicos

## Configuração e Execução

### Pré-requisitos

- Go 1.19 ou superior
- PostgreSQL
- Git

### Instalação

1. Clone o repositório:
```bash
git clone <url-do-repositorio>
cd api-echo
```

2. Instale as dependências:
```bash
go mod tidy
```

3. Configure o banco de dados PostgreSQL e atualize a string de conexão em `main.go`

4. Execute a aplicação:
```bash
go run main.go
```

A API estará disponível em `http://localhost:8080`

## Exemplos de Uso

### Criar Usuário
```bash
curl -X POST http://localhost:8080/user \
  -H "Content-Type: application/json" \
  -d '{"name": "João Silva", "email": "joao@email.com"}'
```

### Buscar Usuário
```bash
curl http://localhost:8080/user/{id}
```

### Listar Usuários
```bash
curl http://localhost:8080/users
```

### Atualizar Usuário
```bash
curl -X PUT http://localhost:8080/user/{id} \
  -H "Content-Type: application/json" \
  -d '{"name": "João Santos", "email": "joao.santos@email.com"}'
```

### Deletar Usuário
```bash
curl -X DELETE http://localhost:8080/user/{id}
```

## Padrões de Resposta

### Sucesso
```json
{
  "id": "uuid-do-usuario",
  "name": "João Silva",
  "email": "joao@email.com",
  "created_at": "2023-01-01T00:00:00Z",
  "updated_at": "2023-01-01T00:00:00Z"
}
```

### Erro
```json
{
  "error": "Descrição do erro"
}
```

## Funcionalidades Implementadas

- CRUD completo de usuários
- Validação de dados de entrada
- Tratamento consistente de erros
- Geração automática de timestamps
- Identificadores únicos (UUID)
- Middleware de logging
- Arquitetura modular e testável

## Estrutura do Projeto

O projeto implementa injeção de dependências, onde cada camada depende apenas de interfaces, não de implementações concretas. Isso garante baixo acoplamento e alta testabilidade.

### Fluxo de Dados

1. **Handler** recebe requisição HTTP
2. **Handler** valida dados e chama **Service**
3. **Service** aplica regras de negócio e chama **Repository**
4. **Repository** persiste/recupera dados do PostgreSQL
5. Resposta retorna pelo caminho inverso

## Próximos Passos

- Implementação de autenticação e autorização
- Adição de testes unitários e de integração
- Documentação com Swagger
- Containerização com Docker
- Pipeline de CI/CD