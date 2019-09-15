package user

import (
	"context"
	"github.com/ayupov-ayaz/redis/models"
)

type UserRepository interface {
	Create(ctx context.Context, user *models.User) error
	Update(ctx context.Context, user *models.User) error
	Get(ctx context.Context, id uint) (*models.User, error)
	GetByEmail(ctx context.Context, email string) (*models.User, error)
	Delete(ctx context.Context, id uint) error
}
