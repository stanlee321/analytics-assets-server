-- name: CreateAsset :one
INSERT INTO "Assets" (
  internal_id,
  asset_name,
  asset_created_at,
  status,
  asset_link
) VALUES (
  $1, $2, $3, $4, $5
)
RETURNING *;

-- name: GetAsset :one
SELECT * FROM "Assets"
WHERE id = $1 LIMIT 1;

-- name: GetAssetByInternalId :one
SELECT * FROM "Assets"
WHERE internal_id = $1 LIMIT 1;

-- name: GetAssetByAssetName :one
SELECT * FROM "Assets"
WHERE asset_name = $1 LIMIT 1;

-- name: ListAssets :many
SELECT * FROM "Assets"
ORDER BY id;

-- name: DeleteAsset :exec
DELETE FROM "Assets"
WHERE id = $1;

-- name: UpdateAsset :one
UPDATE "Assets" SET status = $2
WHERE id = $1 
RETURNING *;