package repository

import (
	"database/sql"
	"kreditplus-xyz/internal/model"
)

type TransactionRepository struct {
	db *sql.DB
}

func NewTransactionRepository(db *sql.DB) *TransactionRepository {
	return &TransactionRepository{db: db}
}

func (r *TransactionRepository) GetAll() ([]model.Transaction, error) {
	rows, err := r.db.Query("SELECT * FROM transactions")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var transactions []model.Transaction
	for rows.Next() {
		var transaction model.Transaction
		if err := rows.Scan(&transaction.ContractNumber, &transaction.NIK, &transaction.OTR, &transaction.AdminFee, &transaction.Installment, &transaction.Interest, &transaction.AssetName, &transaction.LoanDate, &transaction.DueDate, &transaction.IsPaid); err != nil {
			return nil, err
		}
		transactions = append(transactions, transaction)
	}
	return transactions, nil
}

func (r *TransactionRepository) Create(transaction model.Transaction) error {
	_, err := r.db.Exec("INSERT INTO transactions (contract_number, nik, otr, admin_fee, installment, interest, asset_name, loan_date, due_date, is_paid) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)",
		transaction.ContractNumber, transaction.NIK, transaction.OTR, transaction.AdminFee, transaction.Installment, transaction.Interest, transaction.AssetName, transaction.LoanDate, transaction.DueDate, transaction.IsPaid)
	return err
}

func (r *TransactionRepository) UpdateTransactionStatus(contractNumber string, isPaid bool) error {
	_, err := r.db.Exec("UPDATE transactions SET is_paid = ? WHERE contract_number = ?", isPaid, contractNumber)
	return err
}

func (r *TransactionRepository) GetDueDateStatus(nik string) ([]model.Transaction, error) {
	rows, err := r.db.Query("SELECT * FROM transactions WHERE nik = ? AND due_date > NOW()", nik)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var transactions []model.Transaction
	for rows.Next() {
		var transaction model.Transaction
		if err := rows.Scan(&transaction.ContractNumber, &transaction.NIK, &transaction.OTR, &transaction.AdminFee, &transaction.Installment, &transaction.Interest, &transaction.AssetName, &transaction.LoanDate, &transaction.DueDate, &transaction.IsPaid); err != nil {
			return nil, err
		}
		transactions = append(transactions, transaction)
	}
	return transactions, nil
}
