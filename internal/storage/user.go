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

type UserStorage struct {
	db     *sql.DB
	logger *slog.Logger
	table  string
}

func NewUserStorage(logger *slog.Logger, storagePath string) (*UserStorage, error) {
	db, err := sql.Open("sqlite3", storagePath)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	logger.Info("user storage initialize", slog.String("path", storagePath))

	err = CreateTable(db, dbquery.CreateUserTable)
	if err != nil {
		return nil, err
	}

	return &UserStorage{db: db, logger: logger, table: "users"}, nil
}

func (s *UserStorage) Add(item *models.User) error {
	_, err := s.db.Exec(
		dbquery.InsertToUsers,
		item.Username, item.Password, item.Phone, item.Type)
	if err != nil {
		return err
	}

	s.logger.Debug("query", slog.String("query to storage", dbquery.InsertToUsers))
	return nil
}

func (s *UserStorage) FetchAll() (*[]models.User, error) {
	rows, err := s.db.Query(fmt.Sprintf(dbquery.SelectAll, s.table))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	out, err := sqlparse.User().ParseRows(rows)
	if err != nil {
		return nil, err
	}

	s.logger.Debug("query", slog.String("query to storage", dbquery.SelectAll))
	return out, nil
}

func (s *UserStorage) Fetch(id int64) (*models.User, error) {
	row := s.db.QueryRow(fmt.Sprintf(dbquery.SelectID, s.table), id)
	out, err := sqlparse.User().ParseRow(row)
	if err != nil {
		return nil, err
	}

	s.logger.Debug("query", slog.String("query to storage", dbquery.SelectID))
	return out, nil
}

func (s *UserStorage) FetchByKey(key, val string) (*models.User, error) {
	row := s.db.QueryRow(fmt.Sprintf(dbquery.SelectKey, s.table, key), val)
	out, err := sqlparse.User().ParseRow(row)
	if err != nil {
		return nil, err
	}

	s.logger.Debug("query", slog.String("query to storage", dbquery.SelectKey))
	return out, nil
}

func (s *UserStorage) Update(id int64, new *models.User) error {
	_, err := s.db.Exec(dbquery.UpdateUser, new.Username, new.Phone, new.Type, id)
	if err != nil {
		return err
	}

	s.logger.Debug("query", slog.String("query to storage", dbquery.UpdateUser))
	return nil
}

func (s *UserStorage) Delete(id int64) error {
	_, err := s.db.Exec(fmt.Sprintf(dbquery.Delete, s.table), id)
	if err != nil {
		return err
	}

	s.logger.Debug("query", slog.String("query to storage", dbquery.Delete))
	return nil
}

func (s *UserStorage) Close() {
	s.logger.Debug("user storage", slog.String("storage", "closed"))
	s.db.Close()
}

func (s *UserStorage) DB() *sql.DB {
	return s.db
}

func (s *UserStorage) Table() string {
	return s.table
}
