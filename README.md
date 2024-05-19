# README

## _PT XYZ MULTIFINANCE_



## Deskripsi
Penyesuaian strutur database beserta code golang untuk menyempurnakan bisnis perusahaan XYZ

## Instalasi
Jika ada langkah-langkah instalasi yang perlu dilakukan, Anda bisa menuliskannya di sini. Misalnya:
1. Clone repositori ini ke lokal Anda.
2. Buka file .sql di dalam folder "database".
3. Jalankan skrip SQL untuk membuat database.

## Structuring
Berikut adalah struktur proyek:


  ```sh
  KreditPlusXYZ/
|-- internal/
|   |-- config/
|   |   |-- config.go    # Contains project configuration settings
|   |-- database/
|   |   |-- database.sql # SQL file for database schema
|   |   |-- database.go  # Manages database connections and queries
|   |-- handler/
|   |   |-- customer_handler_test.go  #Handles HTTP requests customer unit testing
|   |   |-- customer.go  # Handles HTTP requests related to customer operations
|   |   |-- transaction.go # Handles HTTP requests related to transaction operations
|   |   |-- customer_handler_test.go  #Handles HTTP requests customer unit testing
|   |-- model/
|   |   |-- customer.go  # Defines the data structure for customer information
|   |   |-- transaction.go # Defines the data structure for transaction information
|   |-- repository/
|   |   |-- customer_repository_test.go  # Manages database interactions for customer unit testing
|   |   |-- customer.go  # Manages database interactions for customer data
|   |   |-- transaction.go # Manages database interactions for transaction data
|   |   |-- transaction_repository_test.go  # Manages database interactions for transaction unit testing
|   |-- service/
|   |   |-- customer_service_test.go  # Implements business logic and interacts with the repository for customer operations
|   |   |-- customer.go  # Implements business logic and interacts with the repository for customer operations
|   |   |-- transaction.go # Implements business logic and interacts with the repository for transaction operations
|-- Dockerfile           # Contains instructions to build a Docker image for the project
|-- go.mod               # Module definition file for managing dependencies
|-- go.sum               # Sum file containing expected cryptographic hashes for module dependencies
|-- main.go              # Entry point of the application

  ```

# Security Measures
- SQL Injection Prevention: 
Penggunaan kueri berparameter di lapisan repositori untuk mencegah serangan injeksi SQL.

- Cross-Site Scripting (XSS) Prevention: 
Pengkodean output yang tepat untuk mencegah serangan XSS.

- Cross-Site Request Forgery (CSRF) Protection: 
Penggunaan token CSRF dalam formulir untuk mencegah serangan CSRF.


## Entity Relationship Diagram

![ERD-Diagram-XYZ](https://github.com/BrianGorument/kreditplus-xyz/assets/95519233/054f6c28-48af-4e59-9733-df2c42a594a4)



## Features

- CRUD Operations: Kode ini mengimplementasikan operasi CRUD (Create, Read, Update, Delete) untuk entitas seperti pelanggan dan transaksi.
- Validasi Data: Terdapat validasi data di beberapa titik, seperti validasi gaji pelanggan yang tidak boleh negatif.
- Koneksi Database: Membuat dan menggunakan koneksi ke database SQL untuk menyimpan dan mengambil data.
- Unit Testing: Setiap lapisan (repository, service, handler) diuji dengan unit testing untuk memastikan fungsi-fungsi dasar berjalan dengan baik.
- HTTP Handler: Menggunakan HTTP handler untuk menangani permintaan masuk dari klien, seperti membuat transaksi baru atau mendapatkan daftar transaksi.
- Error Handling: Mengelola error dengan baik, seperti memberikan respon yang tepat dan informatif saat terjadi kesalahan.
- Pemisahan Kode: Memisahkan kode ke dalam paket-paket terpisah seperti handler, service, dan repository untuk memudahkan pengembangan dan pemeliharaan.
Query SQL Dinamis: Menggunakan query SQL dinamis dengan parameter untuk memperoleh data dari database.


## Tech Stack
### RESTful-API
- [Go](https://go.dev/)
- [Router](https://pkg.go.dev/net/http) - Go Basic Library
- [SQLmock](https://pkg.go.dev/github.com/data-dog/go-sqlmock) - Unit Testing
- [testify](https://pkg.go.dev/github.com/stretchr/testify@v1.9.0#section-readme) - Unit Testing
- [MySQL](https://www.mysql.com/) - SQL Database
