package environment

import (
	"time"
)

type Environment struct {
	Port                   string
	DatabaseURL            string
	DatabaseName           string
	JWTSecret              []byte
	TokenExpiration        time.Duration
	RefreshTokenExpiration time.Duration
}
