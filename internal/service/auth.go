package service

import (
	"errors"
	"github.com/ankodd/demoexam/core/internal/storage"
	"github.com/ankodd/demoexam/core/internal/utils/errs"
	"github.com/ankodd/demoexam/core/internal/utils/hash"
	"github.com/ankodd/demoexam/core/pkg/models"
	"github.com/ankodd/demoexam/core/pkg/validate"
)

// Login authenticates a user in the system.
//
// Returns errs.NotFoundErr if user not found, errs.AuthorizationFailedErr if passwords don't match.
func Login(user *models.User, s *storage.UserStorage) error {
	// Fetching user in storage
	fetchedUser, err := s.FetchByKey("username", user.Username)
	if err != nil {
		return errors.New(errs.NotFoundErr)
	}

	// Comparable passwords
	if err := hash.VerifyPassword(fetchedUser.Password, user.Password); err != nil {
		return errors.New(errs.AuthorizationFailedErr)
	}

	return nil
}

// Register creates a new user in the system.
//
// Returns errs.ConflictErr if user is invalid, errs.UserIsExistsErr if username exists, errs.InternalServerErr on storage error.
func Register(user *models.User, s *storage.UserStorage) error {
	// Validating user
	if err := validate.User(user); err != nil {
		return errors.New(errs.ConflictErr)
	}

	// Check if username exists in storage
	if err := validate.UserIsExists(user, s); err != nil {
		return errors.New(errs.UserIsExistsErr)
	}

	// Hashing password
	user.Password = hash.Password(user.Password)

	// Adding user to storage
	if err := s.Add(user); err != nil {
		return errors.New(errs.InternalServerErr)
	}

	return nil
}
