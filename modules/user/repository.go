package user

import (
	"context"
	"github.com/ayupov-ayaz/redis/models"
)

//go:generate mockgen -destination=../../mocks/user_repository_mock.go -package=mocks github.com/ayupov-ayaz/redis/modules/user UserRepository
type UserRepository interface {
	Create(ctx context.Context, user *models.User) error
	Update(ctx context.Context, user *models.User) error
	Get(ctx context.Context, id uint64) (*models.User, error)
	GetByEmail(ctx context.Context, email string) (*models.User, error)
	Delete(ctx context.Context, id uint64) error
}
