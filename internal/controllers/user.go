package controllers

import (
	"github.com/dennypenta/github-users/internal/models"
)

type UserService interface {
	GetUsersByIDs([]int64) ([]models.User, error)
}

type UserOutput interface {
	OutMany([]models.User)
	Err(error)
}

type UserController struct {
	service UserService
	output  UserOutput
}

func (c *UserController) ShowUsers(ids []int64) {
	users, err := c.service.GetUsersByIDs(ids)

	c.output.Err(err)
	c.output.OutMany(users)
}

func NewUserController(serivce UserService, output UserOutput) *UserController {
	return &UserController{
		service: serivce,
		output:  output,
	}
}
