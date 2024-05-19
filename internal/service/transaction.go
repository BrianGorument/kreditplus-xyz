package service

import (
	"kreditplus-xyz/internal/model"
	"kreditplus-xyz/internal/repository"
	"sync"
)

type TransactionService struct {
	repository *repository.TransactionRepository
	mu         sync.Mutex
}

func NewTransactionService(repo *repository.TransactionRepository) *TransactionService {
	return &TransactionService{
		repository: repo,
	}
}

func (s *TransactionService) GetAllTransactions() ([]model.Transaction, error) {
	return s.repository.GetAll()
}

func (s *TransactionService) CreateTransaction(transaction model.Transaction) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.repository.Create(transaction)
}

func (s *TransactionService) UpdateTransactionStatus(contractNumber string, isPaid bool) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.repository.UpdateTransactionStatus(contractNumber, isPaid)
}

func (s *TransactionService) GetDueDateStatus(nik string) ([]model.Transaction, error) {
	return s.repository.GetDueDateStatus(nik)
}
