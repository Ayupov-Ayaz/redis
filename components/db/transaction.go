package db

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

type TransactionManager struct {
	db 		 *DB
	logger   *zap.Logger
}

type Tx struct {
	tx *sqlx.Tx
}

func (tx *Tx) prepareQuery(query string, arg interface{}) (string, []interface{}, error) {
	query, args, err := sqlx.Named(query, arg)
	if err != nil {
		return query, args, err
	}
	query, args, err = sqlx.In(query, args...)
	if err != nil {
		return query, args, err
	}
	query = tx.tx.Rebind(query)
	return query, args, nil
}

func (tx *Tx) Commit() error {
	return tx.tx.Commit()
}

func (tx *Tx) Rollback() error {
	return tx.tx.Rollback()
}

func (tx *Tx) Exec(query string, arg interface{}) (sql.Result, error) {
	query, args, err := sqlx.Named(query, arg)
	if err != nil {
		return nil, err
	}
	query = tx.tx.Rebind(query)
	result, err := tx.tx.Exec(query, args...)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (tx *Tx) Get(dest interface{}, query string, arg interface{}) error {
	if arg == nil {
		return tx.tx.Get(dest, query)
	}
	query, args, err := tx.prepareQuery(query, arg)
	if err != nil {
		return err
	}
	return tx.tx.Get(dest, query, args...)
}

func (tx *Tx) Select(dest interface{}, query string, arg interface{}) error {
	if arg == nil {
		return tx.tx.Select(dest, query)
	}
	query, args, err := tx.prepareQuery(query, arg)
	if err != nil {
		return err
	}
	return tx.tx.Select(dest, query, args...)
}

func (tx *Tx) Rows(query string, arg interface{}) (*sqlx.Rows, error) {
	if arg == nil {
		return tx.tx.Queryx(query)
	}
	query, args, err := tx.prepareQuery(query, arg)
	if err != nil {
		return nil, err
	}
	return tx.tx.Queryx(query, args)
}

func NewTransactionManager(db *DB) *TransactionManager {
	return &TransactionManager{
		db: db,
		logger: zap.L().Named("transaction_manager"),
	}
}

func (t *TransactionManager) Tx(handler func(tx *Tx) error) error {
	logger := t.logger.Named("Tx")

	tx, err := t.db.Begin()
	if err != nil {
		logger.Error("transaction begin failed", zap.Error(err))
		return err
	}
	if err := handler(tx); err != nil {
		if err := tx.Rollback(); err != nil {
			logger.Error("transaction rollback failed", zap.Error(err))
		}
		return err
	}
	if err := tx.Commit(); err != nil {
		logger.Error("transaction commit failed", zap.Error(err))
		return err
	}
	return nil
}