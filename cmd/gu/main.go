package main

import (
	"fmt"
	"github/internal/controllers"
	"github/internal/input"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("[WARNING]: Error loading .env file")
	}
}

// example: go run cmd/gu/main.go 18236918 4324516  -123 12
func main() {
	ids := input.ParseArgs()
	controller := controllers.NewUserController()
	controller.ShowUsers(ids)
}
