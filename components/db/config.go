package db

type DbConfig interface {
	GetDsn() string
	GetDriver() string
}