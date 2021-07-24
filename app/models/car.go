package models

type Car struct {
	ID      int64  `json:"id"`
	Make    string `json:"make"`
	Model   string `json:"model"`
	Year    uint16 `json:"year"`
	Mileage uint32 `json:"mileage"`
	Price   uint32 `json:"price"`
}
