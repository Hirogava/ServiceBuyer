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
		parsedEndDate, err := time.Parse("2006-01-02", *request.EndDate)
		if err != nil {
			return errors.ErrInvalidDateFormat
		}
		strDate := parsedEndDate.Format("2006-01-02")
		request.EndDate = &strDate
	}

	return nil
}

func ParseCountingRequest(req *model.CountingRequest) error {
	if req.StartDate == "" {
		return errors.ErrZeroStartDate
	}

	startDate, err := time.Parse("2006-01-02", req.StartDate)
	if err != nil {
		return errors.ErrInvalidStartDate
	}

	endDate := time.Now()
	if req.EndDate != nil && *req.EndDate != "" {
		endDate, err = time.Parse("2006-01-02", *req.EndDate)
		if err != nil {
			return errors.ErrInvalidEndDate
		}
	}

	if endDate.Before(startDate) {
		return errors.ErrEndDateBeforeStartDate
	}

	return nil
}
