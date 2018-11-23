package main

import (
	"fmt"
	"github/internal/controllers"
	"github/internal/input"

	"github.com/joho/godotenv"
	"github/internal/settings"
	"github/internal/mappers"
	"github.com/google/go-github/github"
	"net/http"
	"github/internal/services"
	"github/internal/outputs"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("[WARNING]: Error loading .env file")
	}
}

// example: go run cmd/gu/main.go 18236918 4324516  -123 12
func main() {
	concurrencyLimit := settings.Concurrency.Limit
	userMapper := mappers.NewUserMapper()
	userRepo := github.NewClient(http.DefaultClient).Users
	userService := services.NewUserService(userMapper, userRepo, concurrencyLimit)
	userOutput := outputs.NewUserOutput()

	ids := input.ParseArgs()
	controller := controllers.NewUserController(userService, userOutput)
	controller.ShowUsers(ids)
}
