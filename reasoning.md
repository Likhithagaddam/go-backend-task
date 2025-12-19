# Reasoning and Design Decisions

## Understanding of the Problem

The goal of this task was to build a RESTful backend service in Go to manage users with their name and date of birth (DOB).  
A key requirement was to **store DOB in the database** while **calculating the user’s age dynamically** when fetching user data, instead of persisting age.

The solution needed to strictly use the given technology stack (Go, PostgreSQL, SQLC, Fiber, Zap, Validator) and follow clean backend development practices.

---

## Overall Approach

I approached the problem by first making the API functional end-to-end and then progressively improving the structure and design.  
The final solution follows a **clean, layered architecture**, which makes the codebase easy to read, maintain, and extend.

The application is divided into the following layers:

- **Handler Layer** – Handles HTTP requests and responses
- **Service Layer** – Contains business logic (age calculation)
- **Repository Layer** – Handles database operations
- **Routes** – Centralized API route registration
- **Middleware** – Cross-cutting concerns like request ID and logging
- **Logger** – Centralized structured logging using Zap

Each layer has a single responsibility, reducing coupling between components.

---

## Database and Age Calculation Design

The database schema stores only the following fields for a user:
- `id`
- `name`
- `dob`

The **age is not stored** in the database because it is a derived value that changes over time.  
Instead, age is calculated dynamically using Go’s `time` package whenever user data is fetched.

This avoids data inconsistency and follows best practices for handling derived attributes.

---

## Use of SQLC

SQLC is used to generate the database access layer from raw SQL queries.

Key reasons for choosing SQLC:
- Compile-time safety for SQL queries
- Clear separation between SQL and application logic
- Reduced runtime errors
- Improved maintainability

All CRUD operations (create, read, update, delete, list) are handled through SQLC-generated code, and no raw SQL is written in handlers or services.

---

## Validation and Error Handling

Input validation is implemented using `go-playground/validator` to ensure required fields (`name`, `dob`) are present and valid.

The API uses clean and meaningful HTTP status codes:
- `201 Created` for successful user creation
- `200 OK` for successful reads and updates
- `204 No Content` for successful deletion
- `400 Bad Request` for invalid input
- `404 Not Found` when a user does not exist

This makes the API predictable and easy to consume.

---

## Logging and Middleware

Uber Zap is used for structured, production-ready logging.

Two custom middlewares are implemented:
- **Request ID Middleware** – Injects a unique request ID into every response
- **Request Duration Middleware** – Logs execution time, HTTP method, path, status code, and request ID

These features improve observability and debugging, similar to real-world backend systems.

---

## Maintainability and Extensibility

The chosen architecture makes it easy to:
- Add pagination to the list users API
- Add new endpoints or features
- Introduce unit tests (e.g., for age calculation)
- Add Docker support if required

Optional enhancements like Dockerization, pagination, and unit tests were intentionally kept out as they were marked optional, but the current structure supports them easily.

---

## Conclusion

This solution focuses on correctness, clarity, and maintainability.  
It strictly follows the required technology stack, implements all mandatory features, and includes production-style logging and middleware.

The codebase is designed to be clean, understandable, and extensible, making it suitable for real-world backend development.
