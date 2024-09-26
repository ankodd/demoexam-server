package service

import (
	"dario.cat/mergo"
	"errors"
	"github.com/ankodd/demoexam/core/internal/storage"
	"github.com/ankodd/demoexam/core/internal/utils/errs"
	"github.com/ankodd/demoexam/core/pkg/models"
	"github.com/ankodd/demoexam/core/pkg/validate"
	"time"
)

func OrderAdd(order *models.Order, s *storage.OrderStorage) error {
	if err := validate.OrderStatus(order.Status); err != nil {
		return errors.New(errs.ConflictErr)
	}

	order.UpdatedAt = time.Now()

	if err := s.Add(order); err != nil {
		return errors.New(errs.InternalServerErr)
	}

	return nil
}

func OrderFetch(id int64, s *storage.OrderStorage) (*models.Order, error) {
	orders, err := s.Fetch(id)
	if err != nil {
		return nil, errors.New(errs.NotFoundErr)
	}

	return orders, nil
}

func OrderFetchAll(s *storage.OrderStorage) (*[]models.Order, error) {
	orders, err := s.FetchAll()
	if err != nil {
		return nil, errors.New(errs.InternalServerErr)
	}

	return orders, nil
}

func OrderUpdate(id int64, new *models.Order, s *storage.OrderStorage) error {
	if len(new.Status) != 0 {
		if err := validate.OrderStatus(new.Status); err != nil {
			return err
		}
	}

	order, err := s.Fetch(id)
	if err != nil {
		return errors.New(errs.NotFoundErr)
	}

	if err := mergo.Merge(order, new, mergo.WithOverride); err != nil {
		return errors.New(errs.ConflictErr)
	}

	err = s.Update(id, order)
	if err != nil {
		return errors.New(errs.InternalServerErr)
	}

	return nil
}

func OrderDelete(id int64, s *storage.OrderStorage) error {
	err := s.Delete(id)
	if err != nil {
		return errors.New(errs.InternalServerErr)
	}

	return nil
}
