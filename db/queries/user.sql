-- name: CreateUser :one
INSERT INTO users (
    username,
    hashed_password,
    email
) VALUES (
    $1, $2, $3
) RETURNING *;

-- name: GetUser :one
SELECT * FROM users
WHERE username = $1
LIMIT 1;

-- name: CheckUserExists :one
SELECT EXISTS (
    SELECT 1 FROM users WHERE email = $1
);