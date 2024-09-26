package validate

import (
	"errors"
	"github.com/ankodd/demoexam/core/internal/storage"
	"github.com/ankodd/demoexam/core/pkg/models"
)

func UserType(userType models.Type) error {
	switch userType {
	case "executor":
		return nil
	case "client":
		return nil
	default:
		return errors.New("invalid user type")
	}
}

func UserIsExists(user *models.User, st *storage.UserStorage) error {
	if storage.FieldIsExist(st, "username", user.Username) {
		return errors.New("username already exist")
	}

	return nil
}

func Password(password string) error {
	if len(password) < 8 {
		return errors.New("password is too short")
	}

	return nil
}

func Phone(phone string) error {
	if len(phone) < 12 {
		return errors.New("len phone is must be 12")
	}

	if phone[0] != '+' {
		return errors.New("phone must start with '+'")
	}

	return nil
}

func UserUpdate(user *models.User) error {
	if len(user.Type) != 0 {
		if err := UserType(user.Type); err != nil {
			return err
		}
	}

	if len(user.Phone) != 0 {
		if err := Phone(user.Phone); err != nil {
			return err
		}
	}

	if len(user.Password) != 0 {
		if err := Password(user.Password); err != nil {
			return err
		}
	}

	return nil
}

func User(user *models.User) error {
	if err := Password(user.Password); err != nil {
		return errors.New("invalid password")
	}

	if err := Phone(user.Phone); err != nil {
		return errors.New("invalid phone")
	}

	if err := UserType(user.Type); err != nil {
		return err
	}

	return nil
}
