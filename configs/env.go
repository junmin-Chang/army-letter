package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func EnvMongoURI() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("환경변수 불러오기 실패")
	}
	return os.Getenv("MONGOURI")
}
