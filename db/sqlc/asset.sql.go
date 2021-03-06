// Code generated by sqlc. DO NOT EDIT.
// source: asset.sql

package db

import (
	"context"
)

const createAsset = `-- name: CreateAsset :one
INSERT INTO "Assets" (
  internal_id,
  asset_name,
  asset_created_at,
  status,
  asset_link
) VALUES (
  $1, $2, $3, $4, $5
)
RETURNING id, internal_id, asset_name, asset_created_at, status, asset_link, created_at
`

type CreateAssetParams struct {
	InternalID     int64  `json:"internal_id"`
	AssetName      string `json:"asset_name"`
	AssetCreatedAt string `json:"asset_created_at"`
	Status         bool   `json:"status"`
	AssetLink      string `json:"asset_link"`
}

func (q *Queries) CreateAsset(ctx context.Context, arg CreateAssetParams) (Asset, error) {
	row := q.db.QueryRowContext(ctx, createAsset,
		arg.InternalID,
		arg.AssetName,
		arg.AssetCreatedAt,
		arg.Status,
		arg.AssetLink,
	)
	var i Asset
	err := row.Scan(
		&i.ID,
		&i.InternalID,
		&i.AssetName,
		&i.AssetCreatedAt,
		&i.Status,
		&i.AssetLink,
		&i.CreatedAt,
	)
	return i, err
}

const deleteAsset = `-- name: DeleteAsset :exec
DELETE FROM "Assets"
WHERE id = $1
`

func (q *Queries) DeleteAsset(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteAsset, id)
	return err
}

const getAsset = `-- name: GetAsset :one
SELECT id, internal_id, asset_name, asset_created_at, status, asset_link, created_at FROM "Assets"
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetAsset(ctx context.Context, id int64) (Asset, error) {
	row := q.db.QueryRowContext(ctx, getAsset, id)
	var i Asset
	err := row.Scan(
		&i.ID,
		&i.InternalID,
		&i.AssetName,
		&i.AssetCreatedAt,
		&i.Status,
		&i.AssetLink,
		&i.CreatedAt,
	)
	return i, err
}

const getAssetByAssetName = `-- name: GetAssetByAssetName :one
SELECT id, internal_id, asset_name, asset_created_at, status, asset_link, created_at FROM "Assets"
WHERE asset_name = $1 LIMIT 1
`

func (q *Queries) GetAssetByAssetName(ctx context.Context, assetName string) (Asset, error) {
	row := q.db.QueryRowContext(ctx, getAssetByAssetName, assetName)
	var i Asset
	err := row.Scan(
		&i.ID,
		&i.InternalID,
		&i.AssetName,
		&i.AssetCreatedAt,
		&i.Status,
		&i.AssetLink,
		&i.CreatedAt,
	)
	return i, err
}

const getAssetByInternalId = `-- name: GetAssetByInternalId :one
SELECT id, internal_id, asset_name, asset_created_at, status, asset_link, created_at FROM "Assets"
WHERE internal_id = $1 LIMIT 1
`

func (q *Queries) GetAssetByInternalId(ctx context.Context, internalID int64) (Asset, error) {
	row := q.db.QueryRowContext(ctx, getAssetByInternalId, internalID)
	var i Asset
	err := row.Scan(
		&i.ID,
		&i.InternalID,
		&i.AssetName,
		&i.AssetCreatedAt,
		&i.Status,
		&i.AssetLink,
		&i.CreatedAt,
	)
	return i, err
}

const listAssets = `-- name: ListAssets :many
SELECT id, internal_id, asset_name, asset_created_at, status, asset_link, created_at FROM "Assets"
ORDER BY id
`

func (q *Queries) ListAssets(ctx context.Context) ([]Asset, error) {
	rows, err := q.db.QueryContext(ctx, listAssets)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Asset{}
	for rows.Next() {
		var i Asset
		if err := rows.Scan(
			&i.ID,
			&i.InternalID,
			&i.AssetName,
			&i.AssetCreatedAt,
			&i.Status,
			&i.AssetLink,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateAsset = `-- name: UpdateAsset :one
UPDATE "Assets" SET status = $2
WHERE id = $1 
RETURNING id, internal_id, asset_name, asset_created_at, status, asset_link, created_at
`

type UpdateAssetParams struct {
	ID     int64 `json:"id"`
	Status bool  `json:"status"`
}

func (q *Queries) UpdateAsset(ctx context.Context, arg UpdateAssetParams) (Asset, error) {
	row := q.db.QueryRowContext(ctx, updateAsset, arg.ID, arg.Status)
	var i Asset
	err := row.Scan(
		&i.ID,
		&i.InternalID,
		&i.AssetName,
		&i.AssetCreatedAt,
		&i.Status,
		&i.AssetLink,
		&i.CreatedAt,
	)
	return i, err
}
