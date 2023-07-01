package api

import (
	db "autoGarage/db/sqlc"
	"github.com/gin-gonic/gin"
	"net/http"
)

type createOrderTxRequest struct {
	CarID         int64 `json:"car_id"`
	ClientID      int64 `json:"client_id"`
	ManagerID     int64 `json:"manager_id"`
	MagazineID    int64 `json:"magazine_id"`
	DeliveryTime  int32 `json:"delivery_time"`
	CarPrice      int64 `json:"car_price"`
	DeliveryPrice int64 `json:"delivery_price"`
	Tax           int64 `json:"tax"`
	TotalPrice    int64 `json:"total_price"`
}

func (s *Server) CreateOrder(ctx *gin.Context) {
	var req createOrderTxRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	arg := db.CreateOrderTxParams{
		CarID:         req.CarID,
		ClientID:      req.ClientID,
		ManagerID:     req.ManagerID,
		MagazineID:    req.MagazineID,
		DeliveryTime:  req.DeliveryTime,
		CarPrice:      req.CarPrice,
		DeliveryPrice: req.DeliveryPrice,
		Tax:           req.Tax,
		TotalPrice:    req.TotalPrice,
	}
	result, err := s.store.CreateOrderTx(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, result)
}
