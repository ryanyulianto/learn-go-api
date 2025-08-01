package dto

type CustomerData struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	PhoneNumber int    `json:"phone_number"`
	Status      string `json:"status"`
}
type CreateCustomerRequest struct {
	Name        string `json:"name" validate:"required"`
	Email       string `json:"email" validate:"required,email"`
	PhoneNumber int64  `json:"phone_number" validate:"required"`
	Status      string `json:"status" validate:"required,oneof=active inactive"`
}
type UpdateCustomerRequest struct {
	ID          string `json:"id" validate:"required"`
	Name        string `json:"name" validate:"required"`
	Email       string `json:"email" validate:"required,email"`
	PhoneNumber int64  `json:"phone_number" validate:"required"`
	Status      string `json:"status" validate:"required,oneof=active inactive"`
}
