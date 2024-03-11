package environment

import (
	"os"
	"strconv"
	"time"
)

func NewFromEnv(defaults *Environment) *Environment {
	var env = *defaults
	tokenExpirationStr := os.Getenv("TokenExpiration")
	refreshTokenExpirationStr := os.Getenv("RefreshTokenExpiration")
	if t, err := strconv.Atoi(tokenExpirationStr); err == nil {
		env.TokenExpiration = time.Hour * 24 * time.Duration(t)
	}

	if t, err := strconv.Atoi(refreshTokenExpirationStr); err == nil {
		env.RefreshTokenExpiration = time.Hour * 24 * time.Duration(t)
	}

	if port := os.Getenv("ServerPort"); port != "" {
		env.Port = port
	}

	if dbUrl := os.Getenv("MongoUrl"); dbUrl != "" {
		env.DatabaseURL = dbUrl
	}

	if dbName := os.Getenv("MongoDBName"); dbName != "" {
		env.DatabaseName = dbName
	}

	if jwtSecret := os.Getenv("JWTSecret"); jwtSecret != "" {
		env.JWTSecret = []byte(jwtSecret)
	}

	return &env
}

var DefaultEnvironment = Environment{
	Port:                   "8000",
	DatabaseURL:            "mongodb://localhost:27017/",
	DatabaseName:           "PhotoManager",
	JWTSecret:              []byte("Sittin' in the stand of the sports arena, waiting for the show to begin Red lights, green lights, strawberry wine, a good friend of mine, follows the stars Venus and Mars are alright tonight"),
	TokenExpiration:        time.Hour * 24,
	RefreshTokenExpiration: time.Hour * 24 * 30,
}
