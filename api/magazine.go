package api

import (
	db "autoGarage/db/sqlc"
	"database/sql"
	"github.com/gin-gonic/gin"
	"net/http"
)

type createMagazine struct {
	Name string
}

func (s *Server) CreateMagazine(ctx *gin.Context) {
	var req createMagazine
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	magazine, err := s.store.CreateMagazine(ctx, req.Name)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, magazine)
}

func (s *Server) ListMagazine(ctx *gin.Context) {
	arg := db.ListMagazinesParams{
		Limit: 10,
	}
	list, err := s.store.ListMagazines(ctx, arg)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, list)
}
