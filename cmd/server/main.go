package main

import (
	"database/sql"

	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"

	"user-service/internal/handler"
	"user-service/internal/repository"
	"user-service/internal/routes"
        "user-service/db/sqlc"
        "go.uber.org/zap"
        "user-service/internal/logger"
        "user-service/internal/middleware"

)

func main() {
        logg := logger.NewLogger()
        defer logg.Sync()
	dbConn, err := sql.Open(
		"postgres",
		"postgres://postgres:Likhi@1234@localhost:5432/user_service?sslmode=disable",
	)
	if err != nil {
		logg.Fatal("Failed to connect to database", zap.Error(err))

	}

	queries := db.New(dbConn)
        repo := repository.NewUserRepository(queries)

	userHandler := handler.NewUserHandler(repo)

	app := fiber.New()
        app.Use(middleware.RequestID())
        app.Use(middleware.RequestLogger(logg))

	// ✅ register routes
	routes.Register(app, userHandler)

	logg.Info("Server started", zap.String("port", "3000"))

	app.Listen(":3000")
}
