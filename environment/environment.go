package environment

import (
	"os"
	"time"
)

type Environment struct {
	DatabaseURL            string
	DatabaseName           string
	JWTSecret              []byte
	TokenExpiration        time.Duration
	RefreshTokenExpiration time.Duration
}

func NewFromEnv() *Environment {
	return &Environment{
		DatabaseURL:            os.Getenv("MongoUrl"),
		DatabaseName:           os.Getenv("MongoDBName"),
		JWTSecret:              []byte(os.Getenv("JWTSecret")),
		TokenExpiration:        time.Hour * 24,
		RefreshTokenExpiration: time.Hour * 24 * 30,
	}
}

func ApplyEnvironment(env *Environment) {
	os.Setenv("MongoUrl", env.DatabaseURL)
	os.Setenv("MongoDBName", env.DatabaseName)
	os.Setenv("JWTSecret", string(env.JWTSecret))
}
