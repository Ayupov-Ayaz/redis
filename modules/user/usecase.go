package user

import (
	"context"
	"github.com/ayupov-ayaz/redis/models"
)

type UserUsecase interface {
	Create(ctx context.Context, user *models.User) error
	Update(ctx context.Context, user *models.User) error
	Get(ctx context.Context, id uint64) (*models.User, error)
	Delete(ctx context.Context, id uint64) error
}