package lib

import (
	"fmt"

	"github.com/joho/godotenv"
)

type Config struct {
	DBHost     string
	DBPort     string
	DBUsername string
	DBPassword string
	DBName     string
	Migrate    uint8
}

func (config *Config) ReadFromEnv() {
	GetStringEnv(
		"DB_HOST",
		&config.DBHost,
	)
	GetStringEnv(
		"DB_PORT",
		&config.DBPort,
	)
	GetStringEnv(
		"DB_USERNAME",
		&config.DBUsername,
	)
	GetStringEnv(
		"DB_PASSWORD",
		&config.DBPassword,
	)
	GetStringEnv(
		"DB_NAME",
		&config.DBName,
	)
}

func (config *Config) GetPGConnString() string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.DBHost,
		config.DBPort,
		config.DBUsername,
		config.DBPassword,
		config.DBName,
	)
}

func LoadEnv() {
	godotenv.Load("development.env")
}

func InitConfig() *Config {
	config := Config{}
	config.ReadFromEnv()
	return &config
}

func GetTestConfig() Config {
	return Config{
		DBHost:     "localhost",
		DBPort:     "5432",
		DBUsername: "heymat",
		DBPassword: "yukberhemat",
		DBName:     "heymattest",
	}
}
