package main

import (
	"log"
	"net/http"

	"github.com/dennypenta/github-users/internal/config"
	"github.com/dennypenta/github-users/internal/controllers"
	"github.com/dennypenta/github-users/internal/input"
	"github.com/dennypenta/github-users/internal/mappers"
	"github.com/dennypenta/github-users/internal/presenters"
	"github.com/dennypenta/github-users/internal/services"

	"github.com/google/go-github/github"
)

func main() {
	conf, err := config.New()
	if err != nil {
		log.Fatal(err)
	}

	userMapper := mappers.NewUserMapper()
	userRepo := github.NewClient(http.DefaultClient).Users
	userService := services.NewUserService(userMapper, userRepo, conf.ConcurrenyLimit)
	userOutput := presenters.NewUserPresenter()

	ids := input.ParseArgs()
	controller := controllers.NewUserController(userService, userOutput)
	controller.ShowUsers(ids)
}
