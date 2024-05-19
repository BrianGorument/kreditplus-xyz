package handler

import (
	"database/sql"
	"encoding/json"
	"kreditplus-xyz/internal/model"
	"kreditplus-xyz/internal/repository"
	"kreditplus-xyz/internal/service"
	"net/http"
)

type TransactionHandler struct {
	service *service.TransactionService
}

func NewTransactionHandler(db *sql.DB) *TransactionHandler {
	repo := repository.NewTransactionRepository(db)
	return &TransactionHandler{
		service: service.NewTransactionService(repo),
	}
}

func (h *TransactionHandler) Handle(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.GetTransactions(w, r)
	case http.MethodPost:
		h.CreateTransaction(w, r)
	case http.MethodPatch:
		h.UpdateTransactionStatus(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (h *TransactionHandler) GetTransactions(w http.ResponseWriter, r *http.Request) {
	transactions, err := h.service.GetAllTransactions()
	if err != nil {
		jsonResponse(w, http.StatusInternalServerError, false, "Failed to retrieve transactions", 1304, nil)
		return
	}
	jsonResponse(w, http.StatusOK, true, "Successfully retrieved transactions", 0, transactions)
}

func (h *TransactionHandler) CreateTransaction(w http.ResponseWriter, r *http.Request) {
	var transaction model.Transaction
	if err := json.NewDecoder(r.Body).Decode(&transaction); err != nil {
		jsonResponse(w, http.StatusBadRequest, false, "Invalid request payload", 1305, nil)
		return
	}
	if err := h.service.CreateTransaction(transaction); err != nil {
		jsonResponse(w, http.StatusInternalServerError, false, "Failed to create transaction", 1306, nil)
		return
	}
	jsonResponse(w, http.StatusCreated, true, "Transaction created successfully", 0, nil)
}

func (h *TransactionHandler) UpdateTransactionStatus(w http.ResponseWriter, r *http.Request) {
	var req struct {
		ContractNumber string `json:"contract_number"`
		IsPaid         bool   `json:"is_paid"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		jsonResponse(w, http.StatusBadRequest, false, "Invalid request payload", 1307, nil)
		return
	}
	if err := h.service.UpdateTransactionStatus(req.ContractNumber, req.IsPaid); err != nil {
		jsonResponse(w, http.StatusInternalServerError, false, "Failed to update transaction status", 1308, nil)
		return
	}
	jsonResponse(w, http.StatusOK, true, "Transaction status updated successfully", 0, nil)
}

func (h *TransactionHandler) GetDueDateStatus(w http.ResponseWriter, r *http.Request) {
	var req struct {
		NIK string `json:"nik"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		jsonResponse(w, http.StatusBadRequest, false, "Invalid request payload", 1309, nil)
		return
	}
	status, err := h.service.GetDueDateStatus(req.NIK)
	if err != nil {
		jsonResponse(w, http.StatusInternalServerError, false, "Failed to retrieve due date status", 1310, nil)
		return
	}
	jsonResponse(w, http.StatusOK, true, "Successfully retrieved due date status", 0, status)
}
