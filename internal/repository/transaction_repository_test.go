package repository

import (
	"kreditplus-xyz/internal/model"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestGetAllTransactions(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := NewTransactionRepository(db)

	rows := sqlmock.NewRows([]string{"contract_number", "nik", "otr", "admin_fee", "installment", "interest", "asset_name", "loan_date", "due_date", "is_paid"}).
		AddRow("TRX001", "1234567890123456", 15000, 500, 1000, 5, "Bike", "2024-06-01", "2024-07-01", false)

	mock.ExpectQuery("SELECT \\* FROM transactions").WillReturnRows(rows)

	transactions, err := repo.GetAll()
	assert.NoError(t, err)
	assert.Len(t, transactions, 1)
}

func TestCreateTransaction(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := NewTransactionRepository(db)

	transaction := model.Transaction{
		ContractNumber: "TRX001",
		NIK:            "1234567890123456",
		OTR:            15000,
		AdminFee:       500,
		Installment:    1000,
		Interest:       5,
		AssetName:      "Bike",
		LoanDate:       "2024-06-01",
		DueDate:        "2024-07-01",
		IsPaid:         false,
	}

	mock.ExpectExec("INSERT INTO transactions").WithArgs(transaction.ContractNumber, transaction.NIK, transaction.OTR, transaction.AdminFee, transaction.Installment, transaction.Interest, transaction.AssetName, transaction.LoanDate, transaction.DueDate, transaction.IsPaid).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = repo.Create(transaction)
	assert.NoError(t, err)
}

func TestUpdateTransactionStatus(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := NewTransactionRepository(db)

	mock.ExpectExec("UPDATE transactions SET is_paid = \\? WHERE contract_number = \\?").WithArgs(true, "TRX001").
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = repo.UpdateTransactionStatus("TRX001", true)
	assert.NoError(t, err)
}

func TestGetDueDateStatus(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := NewTransactionRepository(db)

	rows := sqlmock.NewRows([]string{"contract_number", "nik", "otr", "admin_fee", "installment", "interest", "asset_name", "loan_date", "due_date", "is_paid"}).
		AddRow("TRX001", "1234567890123456", 15000, 500, 1000, 5, "Bike", "2024-06-01", "2024-07-01", false)

	mock.ExpectQuery("SELECT \\* FROM transactions WHERE nik = \\? AND due_date > NOW()").WithArgs("1234567890123456").WillReturnRows(rows)

	transactions, err := repo.GetDueDateStatus("1234567890123456")
	assert.NoError(t, err)
	assert.Len(t, transactions, 1)
}
