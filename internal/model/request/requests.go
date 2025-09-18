package model

type ServiceRequest struct {
	Name      string  	 `json:"name"`
	Cost      float64 	 `json:"cost"`
	UserID    string  	 `json:"user_id"`
	StartDate string  	 `json:"start_date"`
	EndDate   *string    `json:"end_date,omitempty"`
}

type AmountResponse struct {
	UserID string  `json:"user_id"`
	Amount float64 `json:"amount"`
}

type CountingRequest struct {
	StartDate 	string			`json:"start_date"`
	EndDate 	*string 		`json:"end_date,omitempty"`
	UserID 		*string 		`json:"user_id,omitempty"`
	ServiceName *string 		`json:"service_name,omitempty"`
}
