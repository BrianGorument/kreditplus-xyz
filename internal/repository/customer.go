package repository

import (
	"database/sql"
	"kreditplus-xyz/internal/model"
)

type CustomerRepository struct {
	db *sql.DB
}

func NewCustomerRepository(db *sql.DB) *CustomerRepository {
	return &CustomerRepository{db: db}
}

func (r *CustomerRepository) GetAll() ([]model.Customer, error) {
	rows, err := r.db.Query("SELECT * FROM customers")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	customers := []model.Customer{}
	for rows.Next() {
		var customer model.Customer
		if err := rows.Scan(&customer.NIK, &customer.FullName, &customer.LegalName, &customer.BirthPlace, &customer.BirthDate, &customer.Salary, &customer.PhotoKTP, &customer.PhotoSelfie, &customer.OneMonthLimit, &customer.TwoMonthLimit, &customer.ThreeMonthLimit, &customer.FourMonthLimit); err != nil {
			return nil, err
		}
		customers = append(customers, customer)
	}

	return customers, nil
}

func (r *CustomerRepository) Create(customer model.Customer) error {
	_, err := r.db.Exec("INSERT INTO customers (nik, full_name, legal_name, birth_place, birth_date, salary, photo_ktp, photo_selfie, one_month_limit, two_month_limit, three_month_limit, four_month_limit) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)",
		customer.NIK, customer.FullName, customer.LegalName, customer.BirthPlace, customer.BirthDate, customer.Salary, customer.PhotoKTP, customer.PhotoSelfie, customer.OneMonthLimit, customer.TwoMonthLimit, customer.ThreeMonthLimit, customer.FourMonthLimit)
	return err
}
