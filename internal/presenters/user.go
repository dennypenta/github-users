package presenters

import (
	"fmt"
	"github.com/dennypenta/github-users/internal/models"
)

type UserPresenter struct{}

func (o *UserPresenter) OutMany(users []models.User) {
	for _, u := range users {
		fmt.Println("#", u.Meta.ID, " user")
		fmt.Println("name: ", u.Meta.Name)
		fmt.Println("avatar url: ", u.URL)
		fmt.Println("url: ", u.URL)
	}
}

func (o *UserPresenter) Err(err error) {
	fmt.Println(err.Error())
}

func NewUserPresenter() *UserPresenter {
	return &UserPresenter{}
}
