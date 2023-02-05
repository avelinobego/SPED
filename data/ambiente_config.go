package data

import (
	"fmt"

	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		fmt.Printf("Alerta! Arquivo \".env\" nao encontrado!\n")
	}
}
