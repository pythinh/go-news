package user

import (
	"context"
	"errors"

	"github.com/pythinh/go-news/internal/pkg/types"
	"golang.org/x/crypto/bcrypt"
)

type (
	repository interface {
		FindByUsername(ctx context.Context, username string) (*types.User, error)
		CheckByUsername(ctx context.Context, username string) bool
		FindAll(ctx context.Context) ([]types.User, error)
		Create(ctx context.Context, user *types.User) (string, error)
		Update(ctx context.Context, user *types.User) error
		Delete(ctx context.Context, id string) error
	}
	service struct {
		repo repository
	}
)

func newService(r repository) *service {
	return &service{r}
}

func (s *service) Get(ctx context.Context, username string) (*types.User, error) {
	return s.repo.FindByUsername(ctx, username)
}

func (s *service) Check(ctx context.Context, username string) bool {
	return s.repo.CheckByUsername(ctx, username)
}

func (s *service) GetAll(ctx context.Context) ([]types.User, error) {
	return s.repo.FindAll(ctx)
}

func (s *service) Create(ctx context.Context, user *types.User) (string, error) {
	if user.Password == "" {
		return s.repo.Create(ctx, user)
	}
	hashPass, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	user.Password = string(hashPass)
	return s.repo.Create(ctx, user)
}

func (s *service) CreatePass(ctx context.Context, username, password string) error {
	ok := s.repo.CheckByUsername(ctx, username)
	if !ok {
		return errors.New("user is not exits")
	}
	user, err := s.repo.FindByUsername(ctx, username)
	if err != nil {
		return err
	}
	hashPass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashPass)
	return s.repo.Update(ctx, user)
}

func (s *service) Update(ctx context.Context, user *types.User) error {
	ok := s.repo.CheckByUsername(ctx, user.Username)
	if !ok {
		return errors.New("user is not exits")
	}
	return s.repo.Update(ctx, user)
}

func (s *service) UpdatePass(ctx context.Context, username, oldPassword, newPassword string) error {
	user, err := s.repo.FindByUsername(ctx, username)
	if err != nil {
		return err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(oldPassword))
	if err != nil {
		return err
	}
	hashPass, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashPass)
	return s.repo.Update(ctx, user)
}

func (s *service) Delete(ctx context.Context, id string) error {
	return s.repo.Delete(ctx, id)
}
