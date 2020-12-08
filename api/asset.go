package api

import (
	"database/sql"

	"net/http"
	"github.com/gin-gonic/gin"
	db "github.com/stanlee321/assets_service/db/sqlc"
)

type createAssetRequest struct {
	InternalID     int64  `json:"internal_id"`
	AssetName      string `json:"asset_name"`
	AssetCreatedAt string `json:"asset_created_at"`
}

func (server *Server) createAsset(ctx *gin.Context) {
	var req createAssetRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateAssetParams{
		InternalID: req.InternalID,
		AssetName: req.AssetName,
		AssetCreatedAt : req.AssetCreatedAt,
	}

	account, err := server.store.CreateAsset(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, account)
}



type getAssetRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) getAsset(ctx *gin.Context) {
	var req getAssetRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	account, err := server.store.GetAsset(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, account)
}

type listAssetsRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}

func (server *Server) listAssets(ctx *gin.Context) {
	var req listAssetsRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.ListAssetsParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	accounts, err := server.store.ListAssets(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, accounts)
}


