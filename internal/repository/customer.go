package repository

import (
	"belajar-go-api/domain"
	"context"
	"database/sql"

	"github.com/doug-martin/goqu/v9"
)

type customerRepository struct {
	db *goqu.Database
}

func NewCustomer(con *sql.DB) domain.CustomerRepository {
	return &customerRepository{
		db: goqu.New("default", con),
	}
}

func (cr customerRepository) GetAll(ctx context.Context) ([]domain.Customer, error) {
	panic("haha")
}
func (cr customerRepository) FindById(ctx context.Context, id string) (domain.Customer, error) {
	panic("haha")
}
func (cr customerRepository) Save(ctx context.Context, customer *domain.Customer) error {
	panic("haha")
}
func (cr customerRepository) Update(ctx context.Context, customer *domain.Customer) error {
	panic("haha")
}
func (cr customerRepository) Delete(ctx context.Context, id string) error {
	panic("haha")
}
