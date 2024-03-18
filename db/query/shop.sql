-- name: GetShopByID :one
SELECT * FROM shops
WHERE shop_id = $1 LIMIT 1;

-- name: ListShops :many
SELECT * FROM shops;

-- name: CreateShop :one
INSERT INTO shops (
  owner_id, shop_name, shop_description, shop_address
) VALUES (
  $1, $2, $3, $4
)
RETURNING *;

-- name: UpdateShop :one
UPDATE shops
SET 
  owner_id = $2,
  shop_name = $3,
  shop_description = $4,
  shop_address = $5
WHERE shop_id = $1
RETURNING *;

-- name: DeleteShop :exec
DELETE FROM shops
WHERE shop_id = $1;