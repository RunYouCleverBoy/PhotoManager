package environment

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"playgrounds.com/utils"
)

const (
	SERVER_PORT          = "SERVER_PORT"
	SERVER_DATABASE_URL  = "SERVER_DATABASE_URL"
	SERVER_DATABASE_NAME = "SERVER_DATABASE_NAME"
	SERVER_JWT_SECRET    = "SERVER_JWT"
	SERVER_TOKEN_EXP     = "SERVER_TOKEN_EXPIRATION"
	SERVER_REFRESH_EXP   = "SERVER_REFRESH_TOKEN_EXPIRATION"
)

func NewFromEnv(defaults *Environment) *Environment {
	var envMap = defaults.toMap()
	for k := range envMap {
		if os.Getenv(k) != "" {
			envMap[k] = os.Getenv(k)
		}
	}

	env := Environment{
		Port:                   envMap[SERVER_PORT],
		DatabaseURL:            envMap[SERVER_DATABASE_URL],
		DatabaseName:           envMap[SERVER_DATABASE_NAME],
		JWTSecret:              []byte(envMap[SERVER_JWT_SECRET]),
		TokenExpiration:        daysOrElse(envMap[SERVER_TOKEN_EXP], defaults.TokenExpiration),
		RefreshTokenExpiration: daysOrElse(envMap[SERVER_REFRESH_EXP], defaults.RefreshTokenExpiration),
	}

	return &env
}

var DefaultEnvironment = Environment{
	Port:                   "8000",
	DatabaseURL:            "mongodb://root:IsThereAnybodyGoingToListenToMyStory@localhost:27017/",
	DatabaseName:           "PhotoManager",
	JWTSecret:              []byte("Sittin' in the stand of the sports arena, waiting for the show to begin Red lights, green lights, strawberry wine, a good friend of mine, follows the stars Venus and Mars are alright tonight"),
	TokenExpiration:        time.Hour * 24,
	RefreshTokenExpiration: time.Hour * 24 * 30,
}

func (e *Environment) toMap() map[string]string {
	return map[string]string{
		SERVER_PORT:          e.Port,
		SERVER_DATABASE_URL:  e.DatabaseURL,
		SERVER_DATABASE_NAME: e.DatabaseName,
		SERVER_JWT_SECRET:    string(e.JWTSecret),
		SERVER_TOKEN_EXP:     e.TokenExpiration.String(),
		SERVER_REFRESH_EXP:   e.RefreshTokenExpiration.String(),
	}
}

func daysOrElse(valDays string, d time.Duration) time.Duration {
	const day = 24 * time.Hour
	if days, err := strconv.Atoi(valDays); err == nil {
		return time.Duration(days) * day
	} else {
		return d
	}
}

func (e *Environment) String() string {
	environmentMap := e.toMap()
	environmentMap[SERVER_JWT_SECRET] = environmentMap[SERVER_JWT_SECRET][:10] + "..."
	entries := utils.MapEntries(&environmentMap)
	maxEntry := utils.MaxBy(entries, func(e *utils.MapEntry[string, string]) int {
		return len(e.Key)
	})
	max := len(maxEntry.Key)

	format := fmt.Sprintf("%%-%ds: %%s", max)

	return utils.JoinToString(entries, "\n", func(e *utils.MapEntry[string, string]) string {
		return fmt.Sprintf(format, e.Key, e.Value)
	})
}
