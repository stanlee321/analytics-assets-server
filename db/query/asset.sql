-- name: CreateAsset :one
INSERT INTO "Assets" (
  internal_id,
  asset_name,
  asset_created_at
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: GetAsset :one
SELECT * FROM "Assets"
WHERE id = $1 LIMIT 1;

-- name: ListAssets :many
SELECT * FROM "Assets"
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: DeleteAsset :exec
DELETE FROM "Assets"
WHERE id = $1;
