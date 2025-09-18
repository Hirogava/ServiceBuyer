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

	ErrNoRecordsFound = errors.New("no records found")

	ErrZeroStartDate = errors.New("start_date is required")
	ErrEndDateBeforeStartDate = errors.New("end_date cannot be before start_date")
	ErrInvalidStartDate = errors.New("invalid start_date format")
	ErrInvalidEndDate = errors.New("invalid end_date format")
)
