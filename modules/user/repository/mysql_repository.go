package repository

import (
	"context"
	"database/sql"
	"github.com/ayupov-ayaz/redis/components/db"
	"github.com/ayupov-ayaz/redis/models"
	"github.com/ayupov-ayaz/redis/modules/user"
	"go.uber.org/zap"
)

type mysqlUserRepository struct {
	logger *zap.Logger
	tx     *db.TransactionManager
}

func (rep *mysqlUserRepository) Create(ctx context.Context, user *models.User) error {
	query := `
		INSERT INTO users (name, email, password) VALUES (:name, :email, :password)
	`

	return rep.tx.Tx(func(tx *db.Tx) error {
		result, err := tx.Exec(query, user)
		if err != nil {
			rep.logger.Error("create user failed", zap.Error(err))
			return err
		}

		id, err := result.LastInsertId()
		if err != nil {
			rep.logger.Error("get last id failed", zap.Error(err))
			return err
		}

		user.ID = id
		return nil
	})

}

func (rep *mysqlUserRepository) Update(ctx context.Context, user *models.User) error {
	query := `UPDATE users
				SET name = :name, email = :email, password = :password 
			WHERE id = :id`

	return rep.tx.Tx(func(tx *db.Tx) error {
		if _, err := tx.Exec(query, user); err != nil {
			rep.logger.Error("update user failed", zap.Error(err))
			return err
		}
		return nil
	})
}

func (rep *mysqlUserRepository) Get(ctx context.Context, id uint64) (*models.User, error) {
	query := `
		SELECT * FROM users WHERE id = :id
	`

	arg := map[string]interface{}{
		"id": id,
	}

	user, err := rep.getUserByQuery(query, arg)
	if err != nil {
		rep.logger.Error("get user by id failed", zap.Error(err))
		return nil, err
	}

	return user, nil
}

func (rep *mysqlUserRepository) getUserByQuery(query string, args interface{}) (*models.User, error) {
	user := new(models.User)

	err := rep.tx.Tx(func(tx *db.Tx) error {
		return tx.Get(user, query, args)
	})

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, models.UserNotFound
		}
		return nil, err
	}

	return user, err
}

func (rep *mysqlUserRepository) GetByEmail(ctx context.Context, email string) (*models.User, error) {
	query := `
		SELECT * FROM users WHERE email = :email
	`
	arg := map[string]interface{}{
		"email": email,
	}

	user, err := rep.getUserByQuery(query, arg)
	if err != nil {
		rep.logger.Error("get user by email failed", zap.Error(err))
		return nil, err
	}

	return user, nil
}

func (rep *mysqlUserRepository) Delete(ctx context.Context, id uint64) error {
	query := `
		DELETE FROM users WHERE id = :id
	`

	arg := map[string]interface{}{
		"id": id,
	}

	return rep.tx.Tx(func(tx *db.Tx) error {
		if _, err := tx.Exec(query, arg); err != nil {
			rep.logger.Error("delete user failed", zap.Error(err))
			return err
		}
		return nil
	})
}

func NewMysqlUserRepository(tx *db.TransactionManager) user.UserRepository {
	return &mysqlUserRepository{
		logger: zap.L().Named("db.user_repository"),
		tx:     tx,
	}
}
