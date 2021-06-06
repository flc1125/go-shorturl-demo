package env

import (
	"log"

	"github.com/joho/godotenv"
)

// 载入 ENV 配置
func Env() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
