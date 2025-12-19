package handler

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/go-playground/validator/v10"

	"user-service/internal/repository"
	"user-service/internal/service"
)

type UserHandler struct {
	Repo     *repository.UserRepository
	Validate *validator.Validate
}

type UserRequest struct {
	Name string `json:"name" validate:"required"`
	Dob  string `json:"dob" validate:"required"`
}

func NewUserHandler(repo *repository.UserRepository) *UserHandler {
	return &UserHandler{
		Repo:     repo,
		Validate: validator.New(),
	}
}

// POST /users
func (h *UserHandler) CreateUser(c *fiber.Ctx) error {
	var req UserRequest
	if err := c.BodyParser(&req); err != nil {
		return fiber.ErrBadRequest
	}
	if err := h.Validate.Struct(req); err != nil {
		return fiber.ErrBadRequest
	}

	id, err := h.Repo.CreateUser(req.Name, req.Dob)
	if err != nil {
		return fiber.ErrInternalServerError
	}

	return c.Status(201).JSON(fiber.Map{
		"id":   id,
		"name": req.Name,
		"dob":  req.Dob,
	})
}

// GET /users/:id
func (h *UserHandler) GetUser(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	user, err := h.Repo.GetUserByID(int32(id))
	if err != nil {
		return fiber.ErrNotFound
	}

	age := service.CalculateAge(user.Dob)

	return c.JSON(fiber.Map{
		"id":   user.ID,
		"name": user.Name,
		"dob":  user.Dob.Format("2006-01-02"),
		"age":  age,
	})
}

// GET /users
func (h *UserHandler) ListUsers(c *fiber.Ctx) error {
	users, err := h.Repo.ListUsers()
	if err != nil {
		return fiber.ErrInternalServerError
	}

	response := []fiber.Map{}
	for _, u := range users {
		response = append(response, fiber.Map{
			"id":   u.ID,
			"name": u.Name,
			"dob":  u.Dob.Format("2006-01-02"),
			"age":  service.CalculateAge(u.Dob),
		})
	}

	return c.JSON(response)
}

// PUT /users/:id
func (h *UserHandler) UpdateUser(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	var req UserRequest
	if err := c.BodyParser(&req); err != nil {
		return fiber.ErrBadRequest
	}
	if err := h.Validate.Struct(req); err != nil {
		return fiber.ErrBadRequest
	}

	user, err := h.Repo.UpdateUser(int32(id), req.Name, req.Dob)
	if err != nil {
		return fiber.ErrInternalServerError
	}

	return c.JSON(fiber.Map{
		"id":   user.ID,
		"name": user.Name,
		"dob":  user.Dob.Format("2006-01-02"),
	})
}

// DELETE /users/:id
func (h *UserHandler) DeleteUser(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	if err := h.Repo.DeleteUser(int32(id)); err != nil {
		return fiber.ErrInternalServerError
	}

	return c.SendStatus(204)
}
