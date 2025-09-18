package model

import "time"

type ServiceRequest struct {
	Name      string  `json:"name"`
	Cost      float64 `json:"cost"`
	UserID    string  `json:"user_id"`
	StartDate time.Time `json:"start_date"`
	EndDate   *time.Time `json:"end_date,omitempty"`
}

type AmountResponse struct {
	UserID string  `json:"user_id"`
	Amount float64 `json:"amount"`
}
