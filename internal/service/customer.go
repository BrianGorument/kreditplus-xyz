package service

import (
	"kreditplus-xyz/internal/model"
	"kreditplus-xyz/internal/repository"
)

type CustomerServiceInt interface {
	GetAllCustomers() ([]model.Customer, error)
	CreateCustomer(customer model.Customer) error
}

type CustomerService struct {
	repository *repository.CustomerRepository
}

func NewCustomerService(repo *repository.CustomerRepository) *CustomerService {
	return &CustomerService{
		repository: repo,
	}
}

func (s *CustomerService) GetAllCustomers() ([]model.Customer, error) {
	return s.repository.GetAll()
}

func (s *CustomerService) CreateCustomer(customer model.Customer) error {
	return s.repository.Create(customer)
}
