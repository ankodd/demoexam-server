package sqlparse

import (
	"database/sql"
	"github.com/ankodd/demoexam/core/pkg/models"
)

type UserParser struct{}

func User() *UserParser {
	return &UserParser{}
}

func (p *UserParser) ParseRows(rows *sql.Rows) (*[]models.User, error) {
	out := make([]models.User, 0)
	for rows.Next() {
		var user models.User
		if err := rows.Scan(
			&user.Id,
			&user.CreatedAt,
			&user.Username,
			&user.Password,
			&user.Phone,
			&user.TgChatId,
			&user.Type); err != nil {
			return nil, err
		}
		out = append(out, user)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &out, nil
}

func (p *UserParser) ParseRow(row *sql.Row) (*models.User, error) {
	var user models.User
	if err := row.Scan(
		&user.Id,
		&user.CreatedAt,
		&user.Username,
		&user.Password,
		&user.Phone,
		&user.TgChatId,
		&user.Type); err != nil {
		return nil, err
	}

	return &user, nil
}
