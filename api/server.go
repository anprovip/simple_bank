package api

import (
	"simplebank/db"

	"github.com/gin-gonic/gin"
)

// Server serves HTTP requests for our banking service

type Server struct {
	store  db.Store
	router *gin.Engine
}

func NewServer(store db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	router.GET("/accounts/:id", server.getAccount)
	router.POST("/accounts", server.createAccount)
	router.GET("/accounts", server.listAccounts)
	router.POST("/accounts/:id", server.updateAccount)
	router.DELETE("/accounts/:id", server.deleteAccount)

	server.router = router
	return server
}

// Start runs the HTTP server on a specific address
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
