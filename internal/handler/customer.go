package handler

import (
	"database/sql"
	"encoding/json"
	"kreditplus-xyz/internal/model"
	"kreditplus-xyz/internal/repository"
	"kreditplus-xyz/internal/service"
	"net/http"
)

type CustomerHandler struct {
	service service.CustomerServiceInt
}

func NewCustomerHandler(db *sql.DB) *CustomerHandler {
	repo := repository.NewCustomerRepository(db)
	svc := service.NewCustomerService(repo)
	return &CustomerHandler{service: svc}
}

func (h *CustomerHandler) Handle(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.GetCustomers(w, r)
	case http.MethodPost:
		h.CreateCustomer(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (h *CustomerHandler) GetCustomers(w http.ResponseWriter, r *http.Request) {
	customers, err := h.service.GetAllCustomers()
	if err != nil {
		jsonResponse(w, http.StatusInternalServerError, false, "Failed to retrieve customers", 1301, nil)
		return
	}
	jsonResponse(w, http.StatusOK, true, "Successfully retrieved customers", 0, customers)
}

func (h *CustomerHandler) CreateCustomer(w http.ResponseWriter, r *http.Request) {
	var customer model.Customer
	if err := json.NewDecoder(r.Body).Decode(&customer); err != nil {
		jsonResponse(w, http.StatusBadRequest, false, "Invalid request payload", 1302, nil)
		return
	}
	if err := h.service.CreateCustomer(customer); err != nil {
		jsonResponse(w, http.StatusInternalServerError, false, "Failed to create customer", 1303, nil)
		return
	}
	jsonResponse(w, http.StatusCreated, true, "Customer created successfully", 0, nil)
}

func jsonResponse(w http.ResponseWriter, status int, success bool, message string, errorCode int, data interface{}) {
	response := struct {
		Success   bool        `json:"success"`
		Message   string      `json:"message"`
		ErrorCode int         `json:"error_code,omitempty"`
		Data      interface{} `json:"data"`
	}{
		Success:   success,
		Message:   message,
		ErrorCode: errorCode,
		Data:      data,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(response)
}
