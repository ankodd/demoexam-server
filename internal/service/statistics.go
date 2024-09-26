package service

import (
	"github.com/ankodd/demoexam/core/internal/storage"
	"github.com/ankodd/demoexam/core/pkg/models"
)

func Statistics(s *storage.OrderStorage) (*models.Statistics, error) {
	countCompletedOrder, err := s.CountCompletedOrders()
	if err != nil {
		return nil, err
	}

	averageTime, err := s.AverageTime()
	if err != nil {
		return nil, err
	}

	countTypesFailures, err := s.CountFailuresByTypes()
	if err != nil {
		return nil, err
	}

	return &models.Statistics{
		CountCompletedOrders: countCompletedOrder,
		AverageTimeInHours:   averageTime,
		CountFailuresByTypes: countTypesFailures,
	}, nil
}
