package outputs

import (
	"github/internal/models"
	"fmt"
	)

type UserOutput struct {}

func (o *UserOutput) OutMany(users []*models.User) {
	for _, u := range users {
		fmt.Println("#", u.Meta.ID, " user")
		fmt.Println("name: ", u.Meta.Name)
		fmt.Println("avatar url: ", u.URL)
		fmt.Println("url: ", u.URL)
	}
}

func (o *UserOutput) ErrMany(errs []error) {
	for _, err := range errs {
		fmt.Println(err.Error())
	}
}

func NewUserOutput() *UserOutput {
	return &UserOutput{}
}
