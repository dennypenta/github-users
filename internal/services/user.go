package services

import (
	"context"
	"errors"
	"fmt"
	"github/internal/models"
	"sync"
	"time"

	"github.com/google/go-github/github"
		"github/internal/settings"
	"github/internal/mappers"
	"net/http"
)

<<<<<<< HEAD
const (
	defaultTimeout = time.Millisecond*5000
)

=======
>>>>>>> 284b57459ce6bb4c1b17f9407385872888ceeea4
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
	wg := &sync.WaitGroup{}
	inProgres := make(chan int64, s.ConcurrencyLimit)
	defer close(inProgres)

	for _, id := range IDs {
		inProgres <- id
		wg.Add(1)

		go func(id int64) {
			ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
			defer cancel()
			defer wg.Done()
			defer func() {
				<-inProgres
			}()

			u, _, err := s.source.GetByID(ctx, id)
			if err != nil {
				errs = append(errs, &UserFetchError{ID: id, message: err.Error()})
				return
			}

			users = append(users, s.Mapper(u))
		}(id)
	}

	wg.Wait()
	return users, errs
}

func NewUserService() *UserService {
	return &UserService{
		ConcurrencyLimit: settings.Concurrency.Limit,
		Mapper: mappers.MapUser,
		source: github.NewClient(http.DefaultClient).Users,
	}
}
