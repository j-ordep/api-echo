package db

import (
    "fmt"
    "log"
    "os"

    "gorm.io/driver/postgres"
    "gorm.io/gorm"
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

func Connect() (*gorm.DB, error) {
    config := NewConfig()
    dsn := config.ConnectionString()

    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        return nil, fmt.Errorf("erro ao conectar com o banco de dados: %v", err)
    }

    log.Println("Conectado ao postgres com GORM com sucesso!")
    return db, nil
}