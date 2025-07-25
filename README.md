#  Innoscripta Banking Service

This project is a **Golang backend service** for managing bank accounts and transactions.  
It supports:
- Account creation
- Deposit & withdraw transactions
- Detailed transaction logs (ledger)
- ACID-like consistency (atomic balance updates)
- Horizontal scalability with RabbitMQ & MongoDB

All built with **Go (Gin framework)**, **PostgreSQL**, **MongoDB**, and **RabbitMQ**.

---

## ✨ **Features**

✅ Create new accounts with an initial balance  
✅ Deposit / withdraw funds via asynchronous queue  
✅ Transaction logs stored in MongoDB for fast, scalable queries  
✅ ACID-like consistency to prevent double spending  
✅ Layered architecture: `web` → `service` → `producer/consumer` → `repository`  
✅ Dockerized, ready to run with `docker-compose`  

---

**API Endpoints**

| Method | Path | Purpose |
|--|--|--|
| `POST /accounts` | Create new account |
| `GET /accounts/:id` | Get account details by account number |
| `POST /transactions` | Deposit / withdraw (enqueue transaction) |
| `GET /accounts/:id/transactions?page=1&limit=10` | Paginated transaction log |

> All data is processed asynchronously for performance & reliability.


## 🚀 Getting Started

### Clone the repository
```bash
git clone https://github.com/sbdoddi7/innoscripta.git
cd innoscripta

- go mod tidy
- docker-compose up --build


Mock Generate: 
- mockgen -source=src/model/account.go  -destination=src/account/mocks/account_mock.go -package=mocks

- mockgen -source=src/model/transaction.go  -destination=src/transaction/mocks/account_mock.go -package=mocks

Notes
- Unit tests are not yet fully covered for all service, repository, and handler methods due to time constraints.

Potential Enhancements
- Request validation using go-playground/validator
- Improved error handling & logging
- Auto‑generate API documentation for consumers & front‑end.
- Graceful retry / DLQ for RabbitMQ consumer

- PostMan Collection Link: https://www.postman.com/soma-502420/innoscripata-cs/collection/t27j01n/api-s?action=share&source=copy-link&creator=43671109