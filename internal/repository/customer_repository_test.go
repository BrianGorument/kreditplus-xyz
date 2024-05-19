package repository

import (
	"kreditplus-xyz/internal/model"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestGetAllCustomers(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := NewCustomerRepository(db)

	rows := sqlmock.NewRows([]string{"nik", "full_name", "legal_name", "birth_place", "birth_date", "salary", "photo_ktp", "photo_selfie", "one_month_limit", "two_month_limit", "three_month_limit", "four_month_limit"}).
		AddRow("1234567890123456", "John Doe", "John Doe", "New York", "1980-01-01", 5000, "ktp_photo.jpg", "selfie_photo.jpg", 1000, 2000, 3000, 4000)

	mock.ExpectQuery("SELECT \\* FROM customers").WillReturnRows(rows)

	customers, err := repo.GetAll()
	assert.NoError(t, err)
	assert.Len(t, customers, 1)
	assert.Equal(t, "John Doe", customers[0].FullName)
}

func TestCreateCustomer(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := NewCustomerRepository(db)

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

	err = repo.Create(customer)
	assert.NoError(t, err)
	mock.ExpectationsWereMet()
}
