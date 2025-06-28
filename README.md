# 🧾 Payroll Management System

A backend system to manage employee payroll (payslip), attendance, overtime, and reimbursements — built with Go, Echo, and PostgreSQL.

---

## 🚀 Features

- 🔐 JWT-based authentication
- 👥 Admin & employee roles
- 📆 Attendance period management
- 🕒 Submit attendance & overtime (max 3 hours/day)
- 💵 Submit reimbursement requests
- 💼 Payroll processing per period
- 📄 Employee payslip generation
- 📊 Admin summary of all payslips
- 🧾 Audit log with IP tracking
- 🛡 Request ID tracing & custom logging middleware

---

## ⚙️ Tech Stack

- Language: Go 1.21+
- Framework: [Echo](https://echo.labstack.com)
- Database: PostgreSQL
- ORM: GORM
- Auth: JWT
- Migrations: Raw SQL or `./bin/go-payroll-service migrate up`
- Logging: Echo middleware + request ID
- Testing: `bash test_usecase.sh`

---

## 🛠 Installation & Setup

### 1. Clone the Repository

```bash
git clone https://github.com/alimasyhur/go-payroll-service.git
cd go-payroll-service
```

2. Create env File
```bash
Copy resources/config.example.json to resources/config.json
Sample
{
  "app": {
    "name": "Payroll Service",
    "version": "v1.0.0",
    "httpPort": "9000",
    "jwtSecret": "verysecret"
  },
  "logger": {
    "isEnable": true
  },
  "db": {
    "debugMode": true,
    "driver": "postgres",
    "host": "localhost",
    "maxIdleConnections": 120,
    "maxOpenConnections": 100,
    "name": "dealls-payroll",
    "password": "password",
    "port": 5432,
    "readTimeout": "240s",
    "sslmode": "disable",
    "timeout": "240s",
    "username": "postgres",
    "writeTimeout": "240s"
  }
}
```

3. Run Database Migrations
```bash
run make build
run ./bin/go-payroll-service migrate up
```

4. Seed Initial Data
```bash
run ./bin/go-payroll-service seeder user-seed
```

5. Run the Server
```bash
run ./bin/go-payroll-service
```

6. Test Usecase
```bash
run bash test_usecase.sh
```

## 📘 API Reference

### 🔐 Authentication

| Method | Endpoint       | Description             | Auth Required | Role    |
|--------|----------------|-------------------------|----------------|---------|
| POST   | `/login`       | User login, returns JWT | ❌             | Public  |

---

### 🗓️ Attendance

| Method | Endpoint             | Description                          | Auth | Role     |
|--------|----------------------|--------------------------------------|------|----------|
| POST   | `/attendances`       | Submit check-in / check-out          | ✅   | Employee |

---

### 📅 Attendance Period (Admin)

| Method | Endpoint                             | Description                           | Auth | Role  |
|--------|--------------------------------------|---------------------------------------|------|-------|
| POST   | `/attendance-periods`                | Create a new attendance period        | ✅   | Admin |

---

### ⏱️ Overtime

| Method | Endpoint           | Description                           | Auth | Role     |
|--------|---------------------|---------------------------------------|------|----------|
| POST   | `/overtimes`      | Submit overtime (max 3 hrs/day)       | ✅   | Employee |

---

### 💵 Reimbursement

| Method | Endpoint                | Description                    | Auth | Role     |
|--------|-------------------------|--------------------------------|------|----------|
| POST   | `/reimbursements`       | Submit a reimbursement claim   | ✅   | Employee |

---

### 📄 Payslip

| Method | Endpoint                     | Description                                   | Auth | Role     |
|--------|------------------------------|-----------------------------------------------|------|----------|
| GET    | `/payslips/:payroll_uuid`      | View payslip breakdown for selected payroll   | ✅   | Employee |

---

### 🧮 Payroll (Admin)

| Method | Endpoint                          | Description                                    | Auth | Role  |
|--------|-----------------------------------|------------------------------------------------|------|-------|
| POST   | `/payrolls/run`                   | Generate payslips for active attendance period | ✅   | Admin |
| GET    | `/payrolls/:payroll_uuid/summary` | Get total take-home pay for all employees      | ✅   | Admin |

---

### 🩺 Health Check

| Method | Endpoint          | Description          | Auth | Role  |
|--------|-------------------|----------------------|------|-------|
| GET    | `/health-check`   | Check service status | ❌   | Public |

---

### 🔐 Required Headers (for protected routes)

```http
Authorization: Bearer <jwt_token>
X-Request-ID: <uuid>
Content-Type: application/json
```
