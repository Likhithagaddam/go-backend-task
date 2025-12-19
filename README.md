# Go Backend – User Management Service 
Author: Sai Likhitha Gaddam

A RESTful Go backend service for user management with PostgreSQL, SQLC, and dynamic age calculation. Built with clean architecture, validation, and structured logging.

📌 Overview
------------

This project is a RESTful backend API built using Go (Golang) to manage users with their name and date of birth (DOB).
The service stores DOB in PostgreSQL and calculates age dynamically at runtime using Go’s time package.

The application follows a clean, layered architecture and uses SQLC for type-safe database access.

🛠️ Technology Stack
--------------------

Language: Go (Golang)

Web Framework: GoFiber

Database: PostgreSQL

Database Access Layer: SQLC

Logging: Uber Zap

Validation: go-playground/validator

Middleware: Custom (Request ID & Request duration logging)

📂 Project Structure
---------------------

<img width="323" height="521" alt="image" src="https://github.com/user-attachments/assets/7a0fa235-47a4-4679-82a1-0549e8932fd9" />


🗄️ Database Schema
-------------------

users table
-----------

CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    dob DATE NOT NULL
);

- dob is stored in the database

- age is not stored (calculated dynamically)

🔄 API Endpoints
------------------

- Create User

POST /users

Request

{
  "name": "Alice",
  "dob": "1990-05-10"
}

Response

{
  "id": 1,
  "name": "Alice",
  "dob": "1990-05-10"
}

- Get User by ID

GET /users/{id}

Response
{
  "id": 1,
  "name": "Alice",
  "dob": "1990-05-10",
  "age": 35
}

- List All Users

GET /users

Response
[
  {
    "id": 1,
    "name": "Alice",
    "dob": "1990-05-10",
    "age": 35
  }
]

- Update User

PUT /users/{id}

Request
{
  "name": "Alice Updated",
  "dob": "1991-03-15"
}

Response
{
  "id": 1,
  "name": "Alice Updated",
  "dob": "1991-03-15"
}

- Delete User

DELETE /users/{id}

Response

HTTP 204 No Content

✅ Key Features & Deliverables
-------------------------------

- RESTful API using GoFiber

- PostgreSQL database integration

- SQLC for type-safe query generation

- Clean architecture (Handler, Service, Repository)

- Input validation using go-playground/validator

- DOB stored, age calculated dynamically

- Structured logging using Uber Zap

- Request ID added to every response

- Request duration logging middleware

- Proper HTTP status codes and error handling

🧠 Architecture Overview
-------------------------

Handler Layer - Handles HTTP requests and responses

Service Layer - Contains business logic (age calculation)

Repository Layer - Handles database operations using SQLC

Routes - Centralized API route registration

Middleware - Request ID generation, Request execution time logging

Logger - Centralized Zap logger configuration

🚀 Setup & Run Instructions
----------------------------

1.  Prerequisites

Go (v1.20+ recommended)

PostgreSQL

SQLC installed

2️.  Clone the Repository

git clone <your-github-repo-url>

cd go-backend-task

3️.  Create Database

CREATE DATABASE user_service;

4️.  Run Migration

psql -U postgres -d user_service -f db/migrations/001_create_users.sql

5️.  Generate SQLC Code

cd db/sqlc

sqlc generate

cd ../..

6️.  Install Dependencies

go mod tidy

7️.  Configure Database Connection

Update the database connection string in:

cmd/server/main.go

postgres://postgres:<PASSWORD>@localhost:5432/user_service?sslmode=disable

8️.  Run the Server

go run cmd/server/main.go

Server starts on:
http://localhost:3000

🧪 Testing the API (Windows PowerShell)
---------------------------------------

Invoke-RestMethod `
  -Uri http://127.0.0.1:3000/users `
  -Method Post `
  -ContentType "application/json" `
  -Body '{"name":"Alice","dob":"1990-05-10"}'

Invoke-RestMethod http://127.0.0.1:3000/users/1

📊 Logging & Middleware
-------------------------

Each request includes:

  - Unique X-Request-ID header

  - Request duration logging

Logs are structured using Uber Zap

🏁 Conclusion
-----------------

This project fully satisfies all the mandatory requirements of the task and demonstrates:

- Clean backend architecture

- Type-safe database interaction using SQLC

- Proper use of Go’s time package

- Production-style logging and middleware

- Maintainable and scalable code structure

