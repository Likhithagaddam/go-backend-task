package repository

import (
	"context"
	"time"

	db "user-service/db/sqlc"
)

type UserRepository struct {
	Queries *db.Queries
}

func NewUserRepository(q *db.Queries) *UserRepository {
	return &UserRepository{Queries: q}
}

func (r *UserRepository) CreateUser(name string, dob string) (int32, error) {
	parsedDob, err := time.Parse("2006-01-02", dob)
	if err != nil {
		return 0, err
	}

	user, err := r.Queries.CreateUser(
		context.Background(),
		db.CreateUserParams{
			Name: name,
			Dob:  parsedDob,
		},
	)
	if err != nil {
		return 0, err
	}

	return user.ID, nil
}

func (r *UserRepository) GetUserByID(id int32) (db.User, error) {
	return r.Queries.GetUserByID(context.Background(), id)
}

func (r *UserRepository) ListUsers() ([]db.User, error) {
	return r.Queries.ListUsers(context.Background())
}

func (r *UserRepository) UpdateUser(id int32, name string, dob string) (db.User, error) {
	parsedDob, err := time.Parse("2006-01-02", dob)
	if err != nil {
		return db.User{}, err
	}

	return r.Queries.UpdateUser(
		context.Background(),
		db.UpdateUserParams{
			ID:   id,
			Name: name,
			Dob:  parsedDob,
		},
	)
}

func (r *UserRepository) DeleteUser(id int32) error {
	return r.Queries.DeleteUser(context.Background(), id)
}
