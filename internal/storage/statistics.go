package storage

import (
	"database/sql"
	"github.com/ankodd/demoexam/core/internal/storage/dbquery"
)

func (s *OrderStorage) CountCompletedOrders() (int64, error) {
	var count int64

	row := s.db.QueryRow(dbquery.CountCompletedOrders)
	if err := row.Scan(&count); err != nil {
		return 0, err
	}

	return count, nil
}

func (s *OrderStorage) AverageTime() (float64, error) {
	var averageTime sql.NullFloat64

	row := s.db.QueryRow(dbquery.AverageTime)
	if err := row.Scan(&averageTime); err != nil {
		return 0, err
	}

	if !averageTime.Valid {
		return 0.0, nil
	}

	return averageTime.Float64, nil
}

func (s *OrderStorage) CountFailuresByTypes() (map[string]int64, error) {
	var countMap = map[string]int64{}

	rows, err := s.db.Query(dbquery.CountTypesFailures)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var (
			typeFailure string
			count       int64
		)

		if err := rows.Scan(&typeFailure, &count); err != nil {
			return nil, err
		}

		countMap[typeFailure] = count
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return countMap, nil
}
