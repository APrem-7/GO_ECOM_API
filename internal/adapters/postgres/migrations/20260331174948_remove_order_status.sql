-- +goose Up
ALTER TABLE orders DROP COLUMN IF EXISTS order_status;

-- +goose Down
ALTER TABLE orders ADD COLUMN IF NOT EXISTS order_status TEXT NOT NULL DEFAULT '';