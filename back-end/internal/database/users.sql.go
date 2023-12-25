// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0
// source: users.sql

package database

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users (id, created_at, updated_at, access_token_updated_at, username, password, email, access_token)
VALUES ($1, $2, $3, $4, $5, $6, $7, encode(sha256(random()::text::bytea), 'hex'))
RETURNING id, created_at, updated_at, access_token_updated_at, username, password, email, access_token
`

type CreateUserParams struct {
	ID                   uuid.UUID
	CreatedAt            time.Time
	UpdatedAt            time.Time
	AccessTokenUpdatedAt time.Time
	Username             string
	Password             string
	Email                string
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, createUser,
		arg.ID,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.AccessTokenUpdatedAt,
		arg.Username,
		arg.Password,
		arg.Email,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.AccessTokenUpdatedAt,
		&i.Username,
		&i.Password,
		&i.Email,
		&i.AccessToken,
	)
	return i, err
}

const getUserByAPIKey = `-- name: GetUserByAPIKey :one
SELECT id, created_at, updated_at, access_token_updated_at, username, password, email, access_token FROM users WHERE access_token = $1
`

func (q *Queries) GetUserByAPIKey(ctx context.Context, accessToken string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserByAPIKey, accessToken)
	var i User
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.AccessTokenUpdatedAt,
		&i.Username,
		&i.Password,
		&i.Email,
		&i.AccessToken,
	)
	return i, err
}

const getUserByEmail = `-- name: GetUserByEmail :one
SELECT id, created_at, updated_at, access_token_updated_at, username, password, email, access_token FROM users WHERE email = $1
`

func (q *Queries) GetUserByEmail(ctx context.Context, email string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserByEmail, email)
	var i User
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.AccessTokenUpdatedAt,
		&i.Username,
		&i.Password,
		&i.Email,
		&i.AccessToken,
	)
	return i, err
}

const getUserByEmailAndPassword = `-- name: GetUserByEmailAndPassword :one
SELECT id, created_at, updated_at, access_token_updated_at, username, password, email, access_token FROM users WHERE (email = $1 and password = $2)
`

type GetUserByEmailAndPasswordParams struct {
	Email    string
	Password string
}

func (q *Queries) GetUserByEmailAndPassword(ctx context.Context, arg GetUserByEmailAndPasswordParams) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserByEmailAndPassword, arg.Email, arg.Password)
	var i User
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.AccessTokenUpdatedAt,
		&i.Username,
		&i.Password,
		&i.Email,
		&i.AccessToken,
	)
	return i, err
}

const updateAccessTokenAndGetUser = `-- name: UpdateAccessTokenAndGetUser :one
UPDATE users SET 
access_token = encode(sha256(random()::text::bytea), 'hex'), 
access_token_updated_at = $3 
WHERE (email = $1 and password = $2)
RETURNING id, created_at, updated_at, access_token_updated_at, username, password, email, access_token
`

type UpdateAccessTokenAndGetUserParams struct {
	Email                string
	Password             string
	AccessTokenUpdatedAt time.Time
}

func (q *Queries) UpdateAccessTokenAndGetUser(ctx context.Context, arg UpdateAccessTokenAndGetUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, updateAccessTokenAndGetUser, arg.Email, arg.Password, arg.AccessTokenUpdatedAt)
	var i User
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.AccessTokenUpdatedAt,
		&i.Username,
		&i.Password,
		&i.Email,
		&i.AccessToken,
	)
	return i, err
}

const updateAccessTokenExpiryTimeAndGetUser = `-- name: UpdateAccessTokenExpiryTimeAndGetUser :one
UPDATE users SET 
access_token_updated_at = $2 
WHERE email = $1
RETURNING id, created_at, updated_at, access_token_updated_at, username, password, email, access_token
`

type UpdateAccessTokenExpiryTimeAndGetUserParams struct {
	Email                string
	AccessTokenUpdatedAt time.Time
}

func (q *Queries) UpdateAccessTokenExpiryTimeAndGetUser(ctx context.Context, arg UpdateAccessTokenExpiryTimeAndGetUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, updateAccessTokenExpiryTimeAndGetUser, arg.Email, arg.AccessTokenUpdatedAt)
	var i User
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.AccessTokenUpdatedAt,
		&i.Username,
		&i.Password,
		&i.Email,
		&i.AccessToken,
	)
	return i, err
}
