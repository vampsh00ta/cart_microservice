package config

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"net/http"

	"log"
	"os"
	"time"
)

type Config struct {
	Env        string `yaml:"env" env-default:"local"`
	Secret     string `yaml:"secret"  env-required:"true"`
	HTTPServer `yaml:"http_server"`
	Redis      `yaml:"redis"`
}

type JwtCustomClaim struct {
	Id       uuid.UUID `json:"id"`
	Username string    `json:"name"`
	Admin    bool      `json:"admin"`
	jwt.RegisteredClaims
}

type HTTPServer struct {
	Address     string        `yaml:"address" env-default:"localhost:8000"`
	Timeout     time.Duration `yaml:"timeout" env-default:"4s"`
	IdleTimeout time.Duration `yaml:"idle_timeout" env-default:"60s"`
}

type Redis struct {
	Address  string `yaml:"address"`
	Password string `yaml:"password"`
	Db       int    `yaml:"db"`
}

func MustLoad() *Config {
	configPath := "/Users/vladislavtrofimov/GolandProjects/cart_mircoservice/config/config.yaml"
	println(configPath)
	if configPath == "" {
		log.Fatal("CONFIG_PATH is not set")
	}
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatal("config file does not exist")
	}
	var cfg Config

	return &cfg
}

var (
	ErrInvalidToken = errors.New("Invalid token")
	AlreadyInCart   = errors.New("Already in cart")
	ValidationError = errors.New("Validation error")
)

func CodeFrom(err error) int {
	switch err {

	case ErrInvalidToken:
		return http.StatusUnauthorized
	case ValidationError:
		return http.StatusBadRequest
	case AlreadyInCart:
		return http.StatusForbidden
	default:
		return http.StatusInternalServerError
	}
}
