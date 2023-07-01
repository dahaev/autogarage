package api

import (
	db "autoGarage/db/sqlc"
	"github.com/gin-gonic/gin"
)

type Server struct {
	store  db.Store
	router *gin.Engine
}

func NewServer(store db.Store) *Server {
	server := &Server{
		store: store,
	}
	router := gin.Default()
	router.POST("/client", server.CreateClient)
	router.GET("/client/:id", server.GetClient)
	router.GET("/client/list", server.GetListClients)
	router.POST("/magazine", server.CreateMagazine)
	router.GET("/magazine/list", server.ListMagazine)
	router.POST("/users", server.CreateUser)
	server.router = router
	return server
}

func (s *Server) Run(address string) error {
	return s.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{
		"error": err.Error(),
	}
}
