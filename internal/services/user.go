package services

import (
	"context"
	"errors"
	"fmt"
	"github/internal/models"
	"time"

	"github.com/google/go-github/github"
		"github/internal/settings"
	"github/internal/mappers"
	"net/http"
)

const (
	defaultTimeout = time.Millisecond*5000
)

var (
	ErrTimeout      = errors.New("time is out")
	FetchErrMessage = "Error while fetching the user #%d: %s"
)

type UserMapper func(*github.User) *models.User

type UserFetchError struct {
	ID      int64
	message string
}

func (e *UserFetchError) Error() string {
	return fmt.Sprintf(FetchErrMessage, e.ID, e.message)
}

type UserService struct {
	ConcurrencyLimit int
	Mapper           UserMapper

	source *github.UsersService
}

func (s *UserService) GetUsersByIDs(IDs []int64) ([]*models.User, []error) {
	users := make([]*models.User, 0)
	errs := make([]error, 0)
	usersConsumer := make(chan *models.User, len(IDs))
	errsConsumer := make(chan error, len(IDs))
	sem := make(chan int64, s.ConcurrencyLimit)
	defer close(sem)
	defer close(usersConsumer)
	defer close(errsConsumer)

	for _, id := range IDs {
		sem <- id

		go func(id int64) {
			ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
			defer cancel()
			defer func() {
				<-sem
			}()

			u, _, err := s.source.GetByID(ctx, id)
			if err != nil {
				errsConsumer <- &UserFetchError{ID: id, message: err.Error()}
				return
			}

			usersConsumer <- s.Mapper(u)
		}(id)
	}

	for i := 0; i < len(IDs); i++ {
		select {
		case u := <- usersConsumer:
			users = append(users, u)
		case e := <- errsConsumer:
			errs = append(errs, e)
		}
	}

	return users, errs
}

func NewUserService() *UserService {
	return &UserService{
		ConcurrencyLimit: settings.Concurrency.Limit,
		Mapper: mappers.MapUser,
		source: github.NewClient(http.DefaultClient).Users,
	}
}
