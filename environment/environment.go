package environment

import "os"

type Environment struct {
	DatabaseURL  string
	DatabaseName string
	JWTSecret    string
}

func NewFromEnv() *Environment {
	return &Environment{
		DatabaseURL:  os.Getenv("MongoUrl"),
		DatabaseName: os.Getenv("MongoDBName"),
		JWTSecret:    os.Getenv("JWTSecret"),
	}
}

func ApplyEnvironment(env *Environment) {
	os.Setenv("MongoUrl", env.DatabaseURL)
	os.Setenv("MongoDBName", env.DatabaseName)
	os.Setenv("JWTSecret", env.JWTSecret)
}
