package domain

import (
	"belajar-go-api/dto"
	"context"
	"database/sql"
)

type Customer struct {
	ID          string       `db:"id"`
	Name        string       `db:"name"`
	Email       string       `db:"email"`
	PhoneNumber int64        `db:"phone_number"`
	Status      string       `db:"status"`
	CreatedAt   sql.NullTime `db:"created_at"`
	UpdatedAt   sql.NullTime `db:"updated_at"`
	DeletedAt   sql.NullTime `db:"deleted_at"`
}

type CustomerRepository interface {
	GetAll(ctx context.Context, status *string) ([]Customer, error)
	FindById(ctx context.Context, id string) (Customer, error)
	FindByEmail(ctx context.Context, email string) (Customer, error)
	Save(ctx context.Context, customer *Customer) error
	Update(ctx context.Context, customer *Customer) error
	Delete(ctx context.Context, id string) error
}
type CustomerService interface {
	Index(ctx context.Context, queries map[string]string) ([]dto.CustomerData, error)
	Create(ctx context.Context, req dto.CreateCustomerRequest) error
	Update(ctx context.Context, req dto.UpdateCustomerRequest) error
}
