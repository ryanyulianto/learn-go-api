package dto

type CustomerData struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	PhoneNumber int    `json:"phone_number"`
	Status      string `json:"status"`
}
