package sqlparse

import (
	"database/sql"
)

type RowsParser[T any] interface {
	ParseRows(rows *sql.Rows) (*[]T, error)
	ParseRow(row *sql.Row) (*T, error)
}
