package db

import (
	"errors"
)

var (
	ErrServiceNameNull = errors.New("service name is null")
	ErrInvalidServiceCost = errors.New("service cost is invalid")
	ErrInvalidUUIDFormat = errors.New("invalid user_id format")
	ErrInvalidDateFormat = errors.New("invalid date format")

	ErrUserNotFound = errors.New("user not found")
	ErrServiceAlreadyExists = errors.New("service already exists")
)
