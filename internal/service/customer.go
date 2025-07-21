package service

import (
	"belajar-go-api/domain"
	"belajar-go-api/dto"
	"context"
	"errors"

	"github.com/google/uuid"
)

type customerService struct {
	customerRepository domain.CustomerRepository
}

func NewCustomer(customerRepository domain.CustomerRepository) domain.CustomerService {
	return &customerService{
		customerRepository: customerRepository,
	}
}
func (cs customerService) Index(ctx context.Context) ([]dto.CustomerData, error) {
	customers, err := cs.customerRepository.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	customerData := make([]dto.CustomerData, len(customers))
	for i, customer := range customers {
		customerData[i] = dto.CustomerData{
			ID:          customer.ID,
			Name:        customer.Name,
			Email:       customer.Email,
			PhoneNumber: customer.PhoneNumber,
			Status:      customer.Status,
		}
	}
	return customerData, nil
}
func (cs customerService) Create(ctx context.Context, req dto.CreateCustomerRequest) error {
	emailIsUsed, err := cs.emailExists(ctx, req.Email)
	if err != nil {
		return err
	}
	if emailIsUsed {
		return errors.New("email already exists")
	}

	customer := domain.Customer{
		ID:          uuid.NewString(),
		Name:        req.Name,
		Email:       req.Email,
		PhoneNumber: req.PhoneNumber,
		Status:      req.Status,
	}
	return cs.customerRepository.Save(ctx, &customer)
}

func (cs customerService) emailExists(ctx context.Context, email string) (bool, error) {
	customer, err := cs.customerRepository.FindByEmail(ctx, email)
	if err != nil {
		return false, err
	}
	return customer.ID != "", nil
}
