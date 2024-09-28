package storage

import (
	"database/sql"
	"fmt"
	"github.com/ankodd/demoexam/core/internal/storage/dbquery"
	"github.com/ankodd/demoexam/core/internal/utils/parse/sqlparse"
	"github.com/ankodd/demoexam/core/pkg/models"
	_ "github.com/mattn/go-sqlite3"
	"log/slog"
)

type OrderStorage struct {
	db     *sql.DB
	logger *slog.Logger
	table  string
}

func NewOrderStorage(logger *slog.Logger, storagePath string) (*OrderStorage, error) {
	db, err := sql.Open("sqlite3", storagePath)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	logger.Info("order storage initialize", slog.String("path", storagePath))

	err = CreateTable(db, dbquery.CreateOrderTable)
	if err != nil {
		return nil, err
	}

	return &OrderStorage{db: db, logger: logger, table: "orders"}, nil
}

func (s *OrderStorage) Add(item *models.Order) error {
	_, err := s.db.Exec(
		dbquery.InsertToOrders,
		item.Hardware,
		item.TypeFailure,
		item.Description,
		item.ClientId,
		item.ExecutorId,
		item.Status)
	if err != nil {
		return err
	}

	s.logger.Debug("query", slog.String("query to storage", dbquery.InsertToOrders))
	return nil
}

func (s *OrderStorage) FetchAll() (*[]models.Order, error) {
	rows, err := s.db.Query(fmt.Sprintf(dbquery.SelectAll, s.table))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	out, err := sqlparse.Order().ParseRows(rows)
	if err != nil {
		return nil, err
	}

	s.logger.Debug("query", slog.String("query to storage", dbquery.SelectAll))
	return out, nil
}

func (s *OrderStorage) Fetch(id int64) (*models.Order, error) {
	row := s.db.QueryRow(fmt.Sprintf(dbquery.SelectID, s.table), id)
	out, err := sqlparse.Order().ParseRow(row)
	if err != nil {
		return nil, err
	}

	s.logger.Debug("query", slog.String("query to storage", dbquery.SelectID))
	return out, nil
}

func (s *OrderStorage) FetchByKey(key, val string) (*[]models.Order, error) {
	rows, err := s.db.Query(fmt.Sprintf(dbquery.SelectKey, s.table, key), val)
	if err != nil {
		return nil, err
	}

	out, err := sqlparse.Order().ParseRows(rows)
	if err != nil {
		return nil, err
	}

	s.logger.Debug("query", slog.String("query to storage", dbquery.SelectKey))
	return out, nil
}

func (s *OrderStorage) Update(id int64, new *models.Order) error {
	_, err := s.db.Exec(
		dbquery.UpdateOrder,
		new.Hardware,
		new.TypeFailure,
		new.Description,
		new.ExecutorId,
		new.Status,
		id)
	if err != nil {
		return err
	}

	s.logger.Debug("query", slog.String("query to storage", dbquery.UpdateOrder))
	return nil
}

func (s *OrderStorage) Delete(id int64) error {
	_, err := s.db.Exec(fmt.Sprintf(dbquery.Delete, s.table), id)
	if err != nil {
		return err
	}

	s.logger.Debug("query", slog.String("query to storage", dbquery.Delete))
	return nil
}

func (s *OrderStorage) Close() {
	s.logger.Debug("order storage", slog.String("storage", "closed"))
	s.db.Close()
}

func (s *OrderStorage) DB() *sql.DB {
	return s.db
}

func (s *OrderStorage) Table() string {
	return s.table
}
