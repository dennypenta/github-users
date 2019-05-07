package mappers

import (
	"github.com/dennypenta/github-users/internal/models"
	"github.com/google/go-github/github"
)

type UserMapper struct {
}

func (m *UserMapper) Map(user *github.User) models.User {
	return models.User{
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
