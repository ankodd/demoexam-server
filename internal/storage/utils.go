package storage

import (
	"database/sql"
	"errors"
	"fmt"
)

func CreateTable(db *sql.DB, tableQuery string) error {
	_, err := db.Exec(tableQuery)

	if err != nil {
		return err
	}

	return nil
}

func FieldIsExist(st Storage, fieldKey, fieldVal string) bool {
	var field string
	err := st.DB().QueryRow(
		fmt.Sprintf("SELECT %s FROM %s WHERE %s = $1", fieldKey, st.Table(), fieldKey),
		fieldVal).Scan(&field)
	if errors.Is(err, sql.ErrNoRows) {
		return false
	}

	return true
}
