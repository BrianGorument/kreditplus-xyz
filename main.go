package main

import (
	"kreditplus-xyz/internal/config"
	"kreditplus-xyz/internal/database"
	"kreditplus-xyz/internal/handler"
	"log"
	"net/http"
)

func main() {
	cfg := config.LoadConfig()

	db := database.Connect(cfg.DB)

	customerHandler := handler.NewCustomerHandler(db)
	transactionHandler := handler.NewTransactionHandler(db)

	http.HandleFunc("/customers", customerHandler.Handle)
	http.HandleFunc("/transaction", transactionHandler.Handle)

	log.Println("Starting server on port", cfg.ServerPort)
	if err := http.ListenAndServe(":"+cfg.ServerPort, nil); err != nil {
		log.Fatalf("Could not start server: %s\n", err)
	}
}
