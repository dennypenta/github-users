package mappers

import (
	"github.com/google/go-github/github"
	"github/internal/models"
)

func MapUser(user *github.User) *models.User {
	return &models.User{
		URL: *user.HTMLURL,
		Meta: models.UserMeta{
			ID:        *user.ID,
			Name:      *user.Name,
			AvatarURL: *user.AvatarURL,
		},
	}
}
