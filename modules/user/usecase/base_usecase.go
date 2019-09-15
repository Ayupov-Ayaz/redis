package usecase

import (
	"context"
	"github.com/ayupov-ayaz/redis/models"
	"github.com/ayupov-ayaz/redis/modules/user"
	"github.com/ayupov-ayaz/redis/services"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type baseUserUsecase struct {
	timeOut time.Duration
	userRepo user.UserRepository
	bcryptSrv services.BcryptService
	logger *zap.Logger
}

func NewBaseUserUsecase(userRepository user.UserRepository, timeOut time.Duration) user.UserUsecase {
	return &baseUserUsecase{
		timeOut:  timeOut,
		userRepo: userRepository,
		bcryptSrv: services.NewBcryptService(bcrypt.DefaultCost),
		logger:   zap.L().Named("user_usecase"),
	}
}

func (uc *baseUserUsecase) Create(ctx context.Context, user *models.User) error {
	c, cancel := context.WithTimeout(ctx, uc.timeOut)
	defer cancel()

	u, err := uc.userRepo.GetByEmail(c, user.Email)
	if err != nil && err != models.UserNotFound {
		uc.logger.Error("get user by email failed", zap.Error(err))
		return err
	}
	if u != nil {
		uc.logger.Error("user with email already exist")
		return models.AlreadyExist
	}

	hashPass, err := uc.bcryptSrv.HashPassword(user.Password)
	if err != nil {
		uc.logger.Error("hashing password failed", zap.Error(err))
		return err
	}
	user.Password = hashPass

	if err := uc.userRepo.Create(c, user); err != nil {
		uc.logger.Error("create user failed", zap.Error(err))
		return err
	}
	return nil
}

func (uc *baseUserUsecase) Update(ctx context.Context, user *models.User) error {
	c, cancel := context.WithTimeout(ctx, uc.timeOut)
	defer cancel()

	if err := uc.userRepo.Update(c, user); err != nil {
		uc.logger.Error("update user failed", zap.Error(err))
		return err
	}
	return nil
}

func (uc *baseUserUsecase) Get(ctx context.Context, id uint64) (*models.User, error) {
	c, cancel := context.WithTimeout(ctx, uc.timeOut)
	defer cancel()

	user, err := uc.userRepo.Get(c, id)
	if err != nil {
		uc.logger.Error("get user by id failed", zap.Error(err))
		return nil, err
	}
	return user, err
}

func (uc *baseUserUsecase) Delete(ctx context.Context, id uint64) error {
	c, cancel := context.WithTimeout(ctx, uc.timeOut)
	defer cancel()

	if err := uc.userRepo.Delete(c, id); err != nil {
		uc.logger.Error("delete user failed", zap.Error(err))
		return err
	}
	return nil
}



