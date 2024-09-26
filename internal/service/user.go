package service

import (
	"dario.cat/mergo"
	"errors"
	"github.com/ankodd/demoexam/core/internal/storage"
	"github.com/ankodd/demoexam/core/internal/utils/errs"
	"github.com/ankodd/demoexam/core/pkg/models"
	"github.com/ankodd/demoexam/core/pkg/validate"
	"strconv"
)

// UserIsAuthorized check if user authorized
func UserIsAuthorized(chatID int64, s *storage.UserStorage) bool {
	if storage.FieldIsExist(s, "chat_id", strconv.FormatInt(chatID, 10)) {
		return true
	}

	return false
}

func UserFetch(id int64, s *storage.UserStorage) (*models.User, error) {
	users, err := s.Fetch(id)
	if err != nil {
		return nil, errors.New(errs.NotFoundErr)
	}

	return users, nil
}

func UserFetchAll(s *storage.UserStorage) (*[]models.User, error) {
	users, err := s.FetchAll()
	if err != nil {
		return nil, errors.New(errs.InternalServerErr)
	}

	return users, nil
}

func UserUpdate(id int64, new *models.User, s *storage.UserStorage) error {
	if err := validate.UserUpdate(new); err != nil {
		return err
	}

	user, err := s.Fetch(id)
	if err != nil {
		return errors.New(errs.NotFoundErr)
	}

	if err := mergo.Merge(user, new, mergo.WithOverride); err != nil {
		return errors.New(errs.InternalServerErr)
	}

	if err := s.Update(id, user); err != nil {
		return errors.New(errs.ConflictErr)
	}

	return nil
}

func UserDelete(id int64, s *storage.UserStorage) error {
	err := s.Delete(id)
	if err != nil {
		return errors.New(errs.InternalServerErr)
	}

	return nil
}
