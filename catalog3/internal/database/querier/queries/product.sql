-- name: InsertProduct :one
INSERT INTO product (
  name,
  price,
  created_at,
  updated_at
) VALUES (
  $1, 
  $2,
  NOW(),
  NOW()
)
RETURNING *;

-- name: FindProduct :one
SELECT * FROM product
WHERE name = $1;
