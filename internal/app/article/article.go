package article

import (
	"context"

	"github.com/pythinh/go-news/internal/pkg/types"
)

type (
	repository interface {
		FindByID(ctx context.Context, id string) (*types.Article, error)
		FindAll(ctx context.Context) ([]types.Article, error)
		Create(ctx context.Context, article types.Article) (string, error)
		Update(ctx context.Context, article types.Article) error
		Delete(ctx context.Context, id string) error
	}
	service struct {
		repo repository
	}
)

func newService(r repository) *service {
	return &service{r}
}

func (s *service) Get(ctx context.Context, id string) (*types.Article, error) {
	return s.repo.FindByID(ctx, id)
}

func (s *service) GetAll(ctx context.Context) ([]types.Article, error) {
	return s.repo.FindAll(ctx)
}

func (s *service) Create(ctx context.Context, article types.Article) (string, error) {
	return s.repo.Create(ctx, article)
}

func (s *service) Update(ctx context.Context, article types.Article) error {
	_, err := s.repo.FindByID(ctx, article.ID)
	if err != nil {
		return err
	}
	return s.repo.Update(ctx, article)
}

func (s *service) Delete(ctx context.Context, id string) error {
	_, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil
	}
	return s.repo.Delete(ctx, id)
}
