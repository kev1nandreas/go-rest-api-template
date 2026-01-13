# Go REST API Template Using PostgreSQL, Redis, and MongoDB Logging🚀

A production-ready **Go REST API template** built with **Gin**, featuring **JWT Authentication**, **Rate Limiting**, **Swagger**, **PostgreSQL**, **Redis**, **MongoDB**, and **database seeders using gofakeit**. Fully Dockerized and ready for local development or deployment.

---

## ✨ Features

* RESTful API (CRUD-ready)
* JWT Authentication
* API Key security
* Rate Limiting (Redis-backed)
* Swagger / OpenAPI documentation
* PostgreSQL with GORM
* Redis cache layer
* MongoDB for logging
* Database migration & seeding
* End-to-End (E2E) tests
* Docker & Docker Compose
* Conventional commits & pre-commit hooks

---

## 📁 Project Structure

```text
go-rest-api-template/
├── .env
├── Dockerfile
├── Makefile
├── README.md
├── cmd
│   └── server
├── docker-compose.yml
├── docs
├── env
│   └── env.go
├── go.mod
├── go.sum
├── pkg
│   ├── api
│   │   ├── books.go
│   │   ├── books_test.go
│   │   ├── router.go
│   │   ├── user.go
│   │   └── user_test.go
│   ├── auth
│   │   ├── auth.go
│   │   └── auth_test.go
│   ├── cache
│   │   ├── cache.go
│   │   └── cache_test.go
│   ├── database
│   │   ├── db.go
│   │   ├── db_test.go
│   │   ├── migration
│   │   │   └── migrate.go
│   │   ├── mongo.go
│   │   └── seeders
│   │       ├── cmd
│   │       │   └── main.go
│   │       ├── seeder.go
│   │       ├── user_seeder.go
│   │       └── book_seeder.go
│   ├── middleware
│   ├── models
│   │   ├── book.go
│   │   └── user.go
│   └── response
│       └── response.go
├── scripts
└── tests
```

---

## ⚙️ Getting Started

### Prerequisites

* Go **1.21+**
* Docker
* Docker Compose

---

### Installation

1. Clone the repository

```bash
git clone https://github.com/kev1nandreas/go-rest-api-template.git
```

2. Enter the project directory

```bash
cd go-rest-api-template
```

3. Create a `.env` file (copy from example or create your own)

```bash
cp .env.example .env
```

4. Run the application

```bash
make up
```

> 📌 See the `Makefile` for additional commands (seed, test, etc).

---

## 🚀 Running the Application For Development

1. Prepare the `.env` file with your configuration.

```bash
cp .env.example .env
```

2. Run the PostgreSQL, Redis, and MongoDB containers (if not already running)

```bash
docker compose up -d db postgres redis mongo
```

3. Build and run the Go application locally

```bash
make run-local
```

or using air

```bash
air
```

---

## 🔧 Makefile Commands

| Command            | Description                                      |
|--------------------|--------------------------------------------------|
| `make setup`       | Install Swagger dependencies and generate docs  |
| `make build-docker`| Build Docker image without cache                |
| `make run-local`   | Start local containers and run the app          |
| `make up`          | Start all containers with Docker Compose        |
| `make down`        | Stop all containers                             |
| `make restart`     | Restart all containers                          |
| `make build`       | Build the Go application                        |
| `make test`        | Run all tests with race detection and coverage  |
| `make seed`        | Run database seeders (up)                       |
| `make seed-clear`  | Clear all seeded data (down)                    |
| `make clean`       | Stop and remove all containers and images       |

---

## 🧬 Database Migration

Migrations run automatically on app startup.

Migration files are located in:

```text
pkg/database/migration
```

---

## 🌱 Database Seeder (gofakeit)

This project includes **database seeders** using [`gofakeit`](https://github.com/brianvoe/gofakeit) to generate realistic fake data.

### Seeder Structure

```text
pkg/database/seeders
├── cmd
│   └── main.go
├── seeder.go
├── user_seeder.go
└── book_seeder.go
```

### Run Seeder (Up)

```bash
make seed
```

This will:

* Seed users (10 by default)
* Seed books (20 by default)

### Clear Seeded Data (Down)

```bash
make seed-clear
```

### Run Manually

You can also run the seeder manually with args:

```bash
# Seed the database
go run pkg/database/seeders/main.go up

# Clear seeded data
go run pkg/database/seeders/main.go down
```

---

## 🔒 Pre Commit and Conventional Commits

This project uses **pre-commit hooks** to enforce code quality and **conventional commits** for commit messages.

This pre-commit script runs automatically before each commit to check code formatting, linting, and tests.

---

## 📘 API Documentation (Swagger)

Once the app is running, open:

```text
http://localhost:8001/swagger/index.html
```

---

## 🔌 API Endpoints

### Auth

* `POST /api/v1/register`
* `POST /api/v1/login`

### Books

* `GET /api/v1/books`
* `GET /api/v1/books/:id`
* `POST /api/v1/books`
* `PUT /api/v1/books/:id`
* `DELETE /api/v1/books/:id`

---

## 🔑 Authentication

Include JWT token in request headers:

```bash
Authorization: Bearer <YOUR_TOKEN>
```

Example:

```bash
curl -H "Authorization: Bearer <TOKEN>" \
http://localhost:8001/api/v1/books
```

---

## 🧪 End-to-End (E2E) Tests

E2E tests are written in **Python (pytest)**.

### Setup

```bash
cd tests
python -m venv venv
source venv/bin/activate
pip install -r requirements.txt
```

### Environment Variables

```bash
export BASE_URL=http://localhost:8001/api/v1
export API_KEY=your-api-key
```

### Run Tests

```bash
pytest e2e.py
```

---