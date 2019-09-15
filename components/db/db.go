package db

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type DB struct {
	*sqlx.DB
}

func NewDB(cfg DbConfig) (*DB, error) {
	var dsn string

	dsn = cfg.GetDsn()
	db, err := sqlx.Connect(cfg.GetDriver(), dsn)
	if err != nil {
		return nil, err
	}

	return &DB{db}, nil
}

func (db *DB) Begin() (*Tx, error) {
	tx, err := db.Beginx()
	if err != nil {
		return nil, err
	}
	return &Tx{tx}, nil
}
