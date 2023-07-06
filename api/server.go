package api

import (
	db "autoGarage/db/sqlc"
	"autoGarage/token"
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

var tokenMakerSymmetricKey = "12378437289437289283721827374758"
var accessTokenDuration time.Duration = 15 * time.Minute

type Server struct {
	store      db.Store
	router     *gin.Engine
	tokenMaker token.Maker
}

func NewServer(store db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(tokenMakerSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}
	server := &Server{
		store:      store,
		tokenMaker: tokenMaker,
	}
	router := gin.Default()
	router.POST("/users", server.CreateUser)
	router.POST("/users/login", server.loginUser)

	authRoutes := router.Group("/").Use(authMiddleware(server.tokenMaker))

	authRoutes.POST("/client", server.CreateClient)
	authRoutes.GET("/client/:id", server.GetClient)
	authRoutes.GET("/client/list", server.GetListClients)
	authRoutes.POST("/magazine", server.CreateMagazine)
	authRoutes.GET("/magazine/list", server.ListMagazine)

	server.router = router
	return server, nil
}

func (s *Server) Run(address string) error {
	return s.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{
		"error": err.Error(),
	}
}
