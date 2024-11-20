package runtime

import (
	"os"

	"github.com/joho/godotenv"
)

const (
	envFile  = ".env"
	keyToken = "TOKEN"
)

type env struct {
	Token string
}

func NewEnv() env {
	err := godotenv.Load(envFile)
	if err != nil {
		panic(err)
	}

	return env{
		Token: os.Getenv(keyToken),
	}
}
