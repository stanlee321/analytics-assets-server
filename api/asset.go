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
	Status     		bool  `json:"status"`
	AssetLink	   string `json:"asset_link"`
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
		Status : req.Status,
		AssetLink : req.AssetLink,
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


func (server *Server) getAssetByInternalID(ctx *gin.Context) {
	var req getAssetRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	
	account, err := server.store.GetAssetByInternalId(ctx, req.ID)
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



type getAssetRequestByName struct {
	AssetName string `uri:"name" binding:"required"`
}

func (server *Server) getAssetByAssetName(ctx *gin.Context) {
	var req getAssetRequestByName
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	
	account, err := server.store.GetAssetByAssetName(ctx, req.AssetName)
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




func (server *Server) listAssets(ctx *gin.Context) {

	accounts, err := server.store.ListAssets(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, accounts)
}

type updateAssetStatusRequest struct {
	ID 			int64 `json:"id" binding:"required,min=1"`
	Status     	bool  `json:"status"`
}

func (server *Server) updateAssetStatusRequest(ctx *gin.Context) {
	var req updateAssetStatusRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.UpdateAssetParams{
		ID: req.ID,
		Status: req.Status,
	}

	asset, err := server.store.UpdateAsset(ctx, arg)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, asset)
}

