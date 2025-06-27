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
- Testing: `go test`, `httptest`

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

🔐 Authentication
Login:
POST /login
Body: {"username": "admin", "password": "payroll123"}

Authorization:
Add header to all requests:
```bash
Authorization: Bearer <your_token>
```
