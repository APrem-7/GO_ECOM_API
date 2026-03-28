-- name: ListProducts :many
SELECT * FROM products;

-- name: GetProductByID :one
SELECT * FROM products WHERE id=$1;

-- name: createOrderItem :one
INSERT INTO order_items(order_id, quantity,price_in_centers,product_id) VALUES($1, $2, $3, $4) RETURNING *;