package services

import (
	"context"
	"errors"
	"fmt"
	"github.com/dennypenta/github-users/internal/models"
	"strings"
	"time"

	"github.com/google/go-github/github"
)

const (
	defaultTimeout = time.Millisecond * 5000
)

var (
	ErrTimeout = errors.New("time is out")

	FetchErrMessage = "Error while fetching the user #%d: %s"
)

type FetchErrorCollection struct {
	errs []error
}

func(e *FetchErrorCollection) Error() string {
	msgs := make([]string, 0, len(e.errs))

	for _, err := range e.errs {
		msgs = append(msgs, err.Error())
	}

	return strings.Join(msgs, ";")
}

func NewFetchErrorCollection(errs ...error) *FetchErrorCollection {
	return &FetchErrorCollection{errs: errs}
}

type UserMapper interface {
	Map(*github.User) models.User
}

type UserRepo interface {
	GetByID(context.Context, int64) (*github.User, *github.Response, error)
}

type UserFetchError struct {
	ID      int64
	message string
}

func (e *UserFetchError) Error() string {
	return fmt.Sprintf(FetchErrMessage, e.ID, e.message)
}

type UserService struct {
	concurrencyLimit int
	mapper           UserMapper

	repo UserRepo
}

func (s *UserService) GetUsersByIDs(IDs []int64) ([]models.User, error) {
	users := make([]models.User, 0)
	errs := make([]error, 0)
	usersConsumer := make(chan models.User, len(IDs))
	errsConsumer := make(chan error, len(IDs))
	sem := make(chan int64, s.concurrencyLimit)
	defer close(sem)
	defer close(usersConsumer)
	defer close(errsConsumer)

	// ErrGroup is solving this task well too
	for _, id := range IDs {
		sem <- id

		go func(id int64) {
			ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
			defer cancel()
			defer func() {
				<-sem
			}()

			u, _, err := s.repo.GetByID(ctx, id)
			if err != nil {
				errsConsumer <- &UserFetchError{ID: id, message: err.Error()}
				return
			}

			usersConsumer <- s.mapper.Map(u)
		}(id)
	}

	for i := 0; i < len(IDs); i++ {
		select {
		case u := <-usersConsumer:
			users = append(users, u)
		case e := <-errsConsumer:
			errs = append(errs, e)
		}
	}

	return users, NewFetchErrorCollection(errs...)
}

func NewUserService(mapper UserMapper, repo UserRepo, concurrencyLimit int) *UserService {
	return &UserService{
		mapper:           mapper,
		repo:             repo,
		concurrencyLimit: concurrencyLimit,
	}
}
