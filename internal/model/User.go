package model

type User struct {
	ID      int32   `json:"id"`
	FName   string  `json:"fname"`
	City    string  `json:"city"`
	Phone   int64   `json:"phone"` // Using int64 for phone numbers to avoid overflow
	Height  float64 `json:"height"`
	Married bool    `json:"Married"`
}
