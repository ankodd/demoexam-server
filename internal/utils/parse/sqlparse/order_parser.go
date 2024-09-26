package sqlparse

import (
	"database/sql"
	"github.com/ankodd/demoexam/core/pkg/models"
)

type OrderParser struct{}

func Order() *OrderParser {
	return &OrderParser{}
}

func (p *OrderParser) ParseRows(rows *sql.Rows) (*[]models.Order, error) {
	out := make([]models.Order, 0)
	for rows.Next() {
		var order models.Order
		if err := rows.Scan(
			&order.Id,
			&order.CreatedAt,
			&order.UpdatedAt,
			&order.Hardware,
			&order.TypeFailure,
			&order.Description,
			&order.ClientId,
			&order.ExecutorId,
			&order.Status); err != nil {
			return nil, err
		}
		out = append(out, order)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &out, nil
}

func (p *OrderParser) ParseRow(row *sql.Row) (*models.Order, error) {
	var order models.Order
	if err := row.Scan(
		&order.Id,
		&order.CreatedAt,
		&order.UpdatedAt,
		&order.Hardware,
		&order.TypeFailure,
		&order.Description,
		&order.ClientId,
		&order.ExecutorId,
		&order.Status); err != nil {
		return nil, err
	}

	return &order, nil
}
