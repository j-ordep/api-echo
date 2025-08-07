package db

import (
    "database/sql"
    "fmt"
    "log"
    "os"

    _ "github.com/lib/pq"
)

type Config struct {
    Host     string
    Port     string
    User     string
    Password string
    DBName   string
    SSLMode  string
}

func NewConfig() *Config {
    return &Config{
        Host:     getEnv("DB_HOST", "localhost"),	// Host: "localhost"
        Port:     getEnv("DB_PORT", "5432"), 	  	// Port: "5432"
        User:     getEnv("DB_USER", "postgres"),  	// ...
        Password: getEnv("DB_PASSWORD", "postgres"),
        DBName:   getEnv("DB_NAME", "api_echo"),
        SSLMode:  getEnv("DB_SSLMODE", "disable"),
    }
}

/* EXEMPLO:

	getEnv("DB_HOST", "localhost"),
	- verifica se "DB_HOST" existe, se existir atribui "localhost ao atributo da STRUCT"
*/

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key);
    if value != "" {
        return value
    }
    return defaultValue
}

func (c *Config) ConnectionString() string {
    return fmt.Sprintf(
        "host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
        c.Host, c.Port, c.User, c.Password, c.DBName, c.SSLMode,
    )
}

func Connect() (*sql.DB, error) {
    config := NewConfig()
    
    db, err := sql.Open("postgres", config.ConnectionString())
    if err != nil {
        return nil, fmt.Errorf("erro ao abrir conexão com o banco de dados: %v", err)
    }

    db.SetMaxOpenConns(25)
    db.SetMaxIdleConns(25)

	err = db.Ping()
    if err != nil {
        return nil, fmt.Errorf("erro ao conectar: %v", err)
    }

    log.Println("Conectado ao postgres com sucesso!")
	
    return db, nil
}