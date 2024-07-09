package storage

import (
	"database/sql"
	"time"
)

type SourcePostgresStorage struct {
	db *sql.DB
}

type dbSource struct {
	ID                   int64     `db:"id"`
	Name                 string    `db:"name"`
	Surname              string    `db:"surname"`
	Patronymic           string    `db:"patronymic"`
	Address              string    `db:"address"`
	PassportSerialNumber string    `db:"passport_serial_number"`
	PassportNumber       string    `db:"passport_number"`
	CreatedAt            time.Time `db:"created_at"`
}
