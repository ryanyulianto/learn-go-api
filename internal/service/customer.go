package service

import (
	"belajar-go-api/domain"
	"belajar-go-api/dto"
	"context"
	"fmt"
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
	var customerData []dto.CustomerData
	for _, customer := range customers {
		customerData = append(customerData, dto.CustomerData{
			ID:          customer.ID,
			Name:        customer.Name,
			Email:       customer.Email,
			PhoneNumber: customer.PhoneNumber,
			Status:      customer.Status,
		})
	}
	fmt.Println(customerData)
	return customerData, nil
}
