package controllers

import (
	"github/internal/models"
)

type UserService interface {
	GetUsersByIDs([]int64) ([]*models.User, []error)
}

type UserOutput interface {
	OutMany([]*models.User)
	ErrMany([]error)
}

type UserController struct {
	service UserService
	output UserOutput
}

func (c *UserController) ShowUsers(ids []int64) {
	users, errs := c.service.GetUsersByIDs(ids)

	c.output.ErrMany(errs)
	c.output.OutMany(users)
}

func NewUserController(serivce UserService, output UserOutput) *UserController {
	return &UserController{
		service: serivce,
		output: output,
	}
}
