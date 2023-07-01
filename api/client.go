package api

import (
	db "autoGarage/db/sqlc"
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
	"net/http"
)

type createClientRequest struct {
	Name        string `json:"name" binding:"required"`
	Country     string `json:"country" binding:"required"`
	PhoneNumber string `json:"phone_number" binding:"required"`
}

func (s *Server) CreateClient(ctx *gin.Context) {
	var request createClientRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	arg := db.CreateClientParams{
		Name:        request.Name,
		Country:     request.Country,
		PhoneNumber: request.PhoneNumber,
	}
	client, err := s.store.CreateClient(ctx, arg)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "foreign_key_violation", "unique_violation":
				ctx.JSON(http.StatusForbidden, errorResponse(err))
				return
			}
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, client)
}

type getClient struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (s *Server) GetClient(ctx *gin.Context) {
	var request getClient
	if err := ctx.ShouldBindUri(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	client, err := s.store.GetClient(ctx, request.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, client)
}

func (s *Server) GetListClients(ctx *gin.Context) {
	arg := db.ListClientsParams{
		Limit: 10,
	}
	clients, err := s.store.ListClients(ctx, arg)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, clients)
}
