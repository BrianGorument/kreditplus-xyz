package handler

import (
	"bytes"
	"encoding/json"
	"kreditplus-xyz/internal/model"
	"kreditplus-xyz/internal/repository"
	"kreditplus-xyz/internal/service"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestGetTransactionsHandler(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := repository.NewTransactionRepository(db)
	svc := service.NewTransactionService(repo)
	handler := TransactionHandler{service: svc}

	rows := sqlmock.NewRows([]string{"contract_number", "nik", "otr", "admin_fee", "installment", "interest", "asset_name", "loan_date", "due_date", "is_paid"}).
		AddRow("TRX001", "1234567890123456", 15000, 500, 1000, 5, "Bike", "2024-06-01", "2024-07-01", false)

	mock.ExpectQuery("SELECT \\* FROM transactions").WillReturnRows(rows)

	req, err := http.NewRequest(http.MethodGet, "/transactions", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	handler.GetTransactions(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	mock.ExpectationsWereMet()
}

func TestCreateTransactionHandler(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := repository.NewTransactionRepository(db)
	svc := service.NewTransactionService(repo)
	handler := TransactionHandler{service: svc}

	transaction := model.Transaction{
		ContractNumber: "TRX003",
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

	mock.ExpectExec("INSERT INTO transactions").WithArgs(
		transaction.ContractNumber,
		transaction.NIK,
		transaction.OTR,
		transaction.AdminFee,
		transaction.Installment,
		transaction.Interest,
		transaction.AssetName,
		sqlmock.AnyArg(),
		sqlmock.AnyArg(),
		false,
	).WillReturnResult(sqlmock.NewResult(1, 1))

	body, err := json.Marshal(transaction)
	assert.NoError(t, err)

	req, err := http.NewRequest(http.MethodPost, "/transactions", bytes.NewBuffer(body))
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	handler.CreateTransaction(rr, req)

	assert.Equal(t, http.StatusCreated, rr.Code)
	mock.ExpectationsWereMet()
}

func TestUpdateTransactionStatusHandler(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := repository.NewTransactionRepository(db)
	svc := service.NewTransactionService(repo)
	handler := TransactionHandler{service: svc}

	mock.ExpectExec("UPDATE transactions SET is_paid = \\? WHERE contract_number = \\?").WithArgs(true, "TRX003").
		WillReturnResult(sqlmock.NewResult(1, 1))

	body, err := json.Marshal(map[string]interface{}{
		"contract_number": "TRX003",
		"is_paid":         true,
	})
	assert.NoError(t, err)

	req, err := http.NewRequest(http.MethodPost, "/transactions/update-status", bytes.NewBuffer(body))
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	handler.UpdateTransactionStatus(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	mock.ExpectationsWereMet()
}

func TestGetDueDateStatusHandler(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := repository.NewTransactionRepository(db)
	svc := service.NewTransactionService(repo)
	handler := TransactionHandler{service: svc}

	rows := sqlmock.NewRows([]string{"contract_number", "nik", "otr", "admin_fee", "installment", "interest", "asset_name", "loan_date", "due_date", "is_paid"}).
		AddRow("TRX001", "1234567890123456", 15000, 500, 1000, 5, "Bike", "2024-06-01", "2024-07-01", false)

	mock.ExpectQuery("SELECT \\* FROM transactions WHERE nik = \\? AND due_date > NOW()").WithArgs("1234567890123456").WillReturnRows(rows)

	body, err := json.Marshal(map[string]interface{}{
		"nik": "1234567890123456",
	})
	assert.NoError(t, err)

	req, err := http.NewRequest(http.MethodPost, "/transactions/due-date-status", bytes.NewBuffer(body))
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	handler.GetDueDateStatus(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	mock.ExpectationsWereMet()
}
