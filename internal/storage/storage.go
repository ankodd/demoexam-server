package storage

import (
	"database/sql"
)

type Storage interface {
	Close()
	DB() *sql.DB
	Table() string
}
