-- name: CreateTodo :one
INSERT INTO todos (
    user_id,
    title,
    description,
    is_done
) VALUES (
    $1, $2, $3, $4
) RETURNING *;

-- name: GetTodoByID :one 
SELECT * FROM todos 
WHERE id = $1;

-- name: UpdateTodo :one
UPDATE todos
SET 
    title = $3,
    description = $4,
    is_done = $5
WHERE id = $1 AND user_id = $2
RETURNING *;

-- name: DeleteTodo :exec
DELETE FROM todos
WHERE id = $1 AND user_id = $2;

-- name: GetTodosByID :many
SELECT * FROM todos
WHERE user_id = $1
ORDER BY create_at DESC
LIMIT $2 OFFSET $3;
