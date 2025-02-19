-- name: SearchUsers :many
SELECT *
FROM users
WHERE firstname ilike '%' || $1
   OR firstname ilike '%' || $1 || '%'
LIMIT $2;