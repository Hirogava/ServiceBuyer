package db

type User struct {
	UUID string `json:"uuid"`
	Name string `json:"name"`
}

type Service struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Amount int `json:"amount"`
}

type UserPurchase struct {
	User User `json:"user"`
	Service Service `json:"service"`
	StartDate string `json:"start_date"`
	EndDate string `json:"end_date"`
}
