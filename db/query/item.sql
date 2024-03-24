
-- name: GetItemByID :one
SELECT * FROM Items
WHERE item_id = $1 LIMIT 1;

-- name: ListItems :many
SELECT * FROM Items;

-- name: CreateItem :one
INSERT INTO Items (
  owner_id, shop_id, item_code, item_name, description, quantity
) VALUES (
  $1, $2, $3, $4, $5, $6
)
RETURNING *;

-- name: UpdateItem :one
UPDATE Items
SET 
  owner_id = $2,
  shop_id = $3,
  item_code = $4,
  item_name = $5,
  description = $6,
  quantity = $7
WHERE item_id = $1
RETURNING *;

-- name: DeleteItem :exec
DELETE FROM Items
WHERE item_id = $1;
