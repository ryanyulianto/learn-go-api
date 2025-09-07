package repository

import (
	"belajar-go-api/domain"
	"context"
	"database/sql"
	"time"

	"github.com/doug-martin/goqu/v9"
)

const customerTable = "customers"

type customerRepository struct {
	db *goqu.Database
}

func NewCustomer(con *sql.DB) domain.CustomerRepository {
	return &customerRepository{
		db: goqu.New("default", con),
	}
}

func (cr customerRepository) GetAll(ctx context.Context, status *string) (result []domain.Customer, err error) {
	dataset := cr.db.From(customerTable).Where(goqu.C("deleted_at").IsNull()).Order(goqu.C("created_at").Desc())
	if *status != "" {
		dataset = dataset.Where(goqu.C("status").Eq(*status))
	}

	err = dataset.ScanStructsContext(ctx, &result)
	return
}
func (cr customerRepository) FindById(ctx context.Context, id string) (result domain.Customer, err error) {
	dataset := cr.db.From(customerTable).Where(goqu.C("deleted_at").IsNull(), goqu.C("id").Eq(id))
	_, err = dataset.ScanStructContext(ctx, &result)
	return
}
func (cr customerRepository) FindByEmail(ctx context.Context, email string) (result domain.Customer, err error) {
	dataset := cr.db.From(customerTable).Where(goqu.C("deleted_at").IsNull(), goqu.C("email").Eq(email))
	_, err = dataset.ScanStructContext(ctx, &result)
	return
}
func (cr customerRepository) Save(ctx context.Context, customer *domain.Customer) error {
	executor := cr.db.Insert(customerTable).Rows(customer).Executor()
	_, err := executor.ExecContext(ctx)
	return err
}
func (cr customerRepository) Update(ctx context.Context, customer *domain.Customer) error {
	executor := cr.db.Update(customerTable).Where(goqu.C("id").Eq(customer.ID)).Set(customer).Executor()
	_, err := executor.ExecContext(ctx)
	return err
}
func (cr customerRepository) Delete(ctx context.Context, id string) error {
	executor := cr.db.Update(customerTable).Where(goqu.C("id").Eq(id)).Set(goqu.Record{"deleted_at": sql.NullTime{Valid: true, Time: time.Now()}}).Executor()
	_, err := executor.ExecContext(ctx)
	return err
}
func (cr customerRepository) ForceDelete(ctx context.Context, id string) error {
	executor := cr.db.Delete(customerTable).Where(goqu.C("id").Eq(id)).Executor()
	_, err := executor.ExecContext(ctx)
	return err
}
