package db

type User struct {
	UUID string `json:"uuid"`
	Name string `json:"name"`
}

type UserPurchase struct {
	User 	  User 	  `json:"user"`
	Service   Service `json:"service"`
	StartDate string  `json:"start_date"`
	EndDate   string  `json:"end_date"`
}

type Service struct {
	ID 		int 	`json:"id"`
	Name 	string  `json:"name"`
	Amount 	float64 `json:"amount"`
}

type CountingResponse struct {
	UserID 	  string  	`json:"user_id"`
	Amount 	  float64 	`json:"amount"`
	Services  []Service `json:"services"`
	StartDate string 	`json:"start_date"`
	EndDate   string 	`json:"end_date"`
}
