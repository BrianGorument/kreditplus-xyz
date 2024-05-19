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

func TestCreateCustomer(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := repository.NewCustomerRepository(db)
	svc := service.NewCustomerService(repo)
	handler := CustomerHandler{service: svc}

	customer := model.Customer{
		NIK:             "1234567890123456",
		FullName:        "John Doe",
		LegalName:       "John Doe",
		BirthPlace:      "New York",
		BirthDate:       "1980-01-01",
		Salary:          5000,
		PhotoKTP:        "ktp_photo.jpg",
		PhotoSelfie:     "selfie_photo.jpg",
		OneMonthLimit:   1000,
		TwoMonthLimit:   2000,
		ThreeMonthLimit: 3000,
		FourMonthLimit:  4000,
	}

	mock.ExpectExec("INSERT INTO customers").WithArgs(
		customer.NIK,
		customer.FullName,
		customer.LegalName,
		customer.BirthPlace,
		customer.BirthDate,
		customer.Salary,
		customer.PhotoKTP,
		customer.PhotoSelfie,
		customer.OneMonthLimit,
		customer.TwoMonthLimit,
		customer.ThreeMonthLimit,
		customer.FourMonthLimit,
	).WillReturnResult(sqlmock.NewResult(1, 1))

	body, err := json.Marshal(customer)
	assert.NoError(t, err)

	req, err := http.NewRequest(http.MethodPost, "/customers", bytes.NewBuffer(body))
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	handler.CreateCustomer(rr, req)

	assert.Equal(t, http.StatusCreated, rr.Code)
	mock.ExpectationsWereMet()
}

func TestGetAllCustomers(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := repository.NewCustomerRepository(db)
	svc := service.NewCustomerService(repo)
	handler := CustomerHandler{service: svc}

	rows := sqlmock.NewRows([]string{"nik", "full_name", "legal_name", "birth_place", "birth_date", "salary", "photo_ktp", "photo_selfie", "one_month_limit", "two_month_limit", "three_month_limit", "four_month_limit"}).
		AddRow("1234567890123456", "John Doe", "John Doe", "New York", "1980-01-01", 5000, "ktp_photo.jpg", "selfie_photo.jpg", 1000, 2000, 3000, 4000)

	mock.ExpectQuery("SELECT \\* FROM customers").WillReturnRows(rows)

	req, err := http.NewRequest(http.MethodGet, "/customers", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	handler.GetCustomers(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	mock.ExpectationsWereMet()
}
