-- name: ListItems :many
SELECT *
FROM items;

-- name: AddItem :one
INSERT INTO items (name)
VALUES ($1)
RETURNING *;

-- name: DeleteItem :exec
DELETE FROM items
WHERE id = $1;