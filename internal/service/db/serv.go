package db

import (
	"time"

	errors "github.com/Hirogava/ServiceBuyer/internal/errors/db"
	model "github.com/Hirogava/ServiceBuyer/internal/model/request"

	"github.com/google/uuid"
)

func ParseRequest(request *model.ServiceRequest) error {
	if request.Name == "" {
		return errors.ErrServiceNameNull
	}
	if request.Cost <= 0 {
		return errors.ErrInvalidServiceCost
	}
	if _, err := uuid.Parse(request.UserID); err != nil {
		return errors.ErrInvalidUUIDFormat
	}
	if request.EndDate != nil {
		parsedEndDate, err := time.Parse("2006-01-02", request.EndDate.String())
		if err != nil {
			return errors.ErrInvalidDateFormat
		}
		request.EndDate = &parsedEndDate
	}

	return nil
}

func ParseCountingRequest(req *model.CountingRequest) error {
	if req.StartDate.IsZero() {
		return errors.ErrZeroStartDate
	}

	endDate := time.Now()
	if req.EndDate != nil {
		endDate = *req.EndDate
	}
	if endDate.Before(req.StartDate) {
		return errors.ErrEndDateBeforeStartDate
	}

	return nil
}
