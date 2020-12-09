package api


import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	db "github.com/stanlee321/assets_service/db/sqlc"

)

// Server serves HTTP requests for our banking service.
type Server struct {
	store db.Store
	router *gin.Engine
}

// NewServer creates a new HTTP server and set up routing.
func NewServer(store db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()
	router.Use(cors.Default())

	router.POST("/create_asset", server.createAsset)
	router.GET("/read_asset/:id", server.getAsset)
	router.GET("/read_asset_by_internal_id/:id", server.getAssetByInternalID)
	router.GET("/read_asset_by_asset_name/:name", server.getAssetByAssetName)
	router.GET("/list_assets", server.listAssets)
	router.GET("/update_asset", server.updateAssetStatusRequest)


	// router.GET("/accounts/:id", server.getAccount)
	// router.GET("/accounts", server.listAccounts)

	// router.POST("/transfers", server.createTransfer)

	server.router = router
	return server
}

// Start runs the HTTP server on a specific address.
func (server *Server) Start(address string) error {


	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
