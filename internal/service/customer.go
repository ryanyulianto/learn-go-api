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
func (cs customerService) Index(ctx context.Context, queries map[string]string) ([]dto.CustomerData, error) {
	status := queries["status"]
	customers, err := cs.customerRepository.GetAll(ctx, &status)
	if err != nil {
		return nil, err
	}

	customerData := make([]dto.CustomerData, len(customers))
	for i, customer := range customers {
		customerData[i] = dto.CustomerData{
			ID:          customer.ID,
			Name:        customer.Name,
			Email:       customer.Email,
			PhoneNumber: int(customer.PhoneNumber),
			Status:      customer.Status,
		}
	}
	return customerData, nil
}
func (cs customerService) Create(ctx context.Context, req dto.CreateCustomerRequest) error {
	emailIsUsed, err := cs.emailExists(ctx, req.Email, nil)
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

func (cs customerService) Update(ctx context.Context, req dto.UpdateCustomerRequest) error {
	_, err := cs.customerRepository.FindById(ctx, req.ID)
	if err != nil {
		return errors.New("customer not found")
	}
	emailIsUsed, err := cs.emailExists(ctx, req.Email, &req.ID)
	if err != nil {
		return errors.New(err.Error())
	}
	if emailIsUsed {
		return errors.New("email already exists")
	}
	req_data := domain.Customer{
		ID:          req.ID,
		Name:        req.Name,
		Email:       req.Email,
		PhoneNumber: req.PhoneNumber,
		Status:      req.Status,
	}
	return cs.customerRepository.Update(ctx, &req_data)
}

func (cs customerService) emailExists(ctx context.Context, email string, where_not_id *string) (bool, error) {
	customer, err := cs.customerRepository.FindByEmail(ctx, email)
	if err != nil {
		return false, err
	}
	if *where_not_id == "" {
		return false, errors.New("where_not_id is required")
	}
	if customer.ID == *where_not_id {
		return false, errors.New("email already exists")
	}
	return customer.ID != "", nil
}
