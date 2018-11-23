package mappers

import (
	"github.com/google/go-github/github"
	"github/internal/models"
)

type UserMapper struct {

}

func (m *UserMapper) Map(user *github.User) *models.User {
	return &models.User{
		URL: *user.HTMLURL,
		Meta: models.UserMeta{
			ID:        *user.ID,
			Name:      *user.Name,
			AvatarURL: *user.AvatarURL,
		},
	}
}

func NewUserMapper() *UserMapper {
	return &UserMapper{}
}
