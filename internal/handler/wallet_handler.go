package handler

import (
	"github.com/gin-gonic/gin"
	"go.mod/internal/service"
	"net/http"
)

type WalletHandler struct {
	service *service.WalletService
}

func NewWalletHandler(s *service.WalletService) *WalletHandler {
	return &WalletHandler{service: s}
}

func (h *WalletHandler) CheckWallet(c *gin.Context) {
	var req struct {
		WalletID uint `json:"wallet_id"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}
	exists := h.service.CheckWalletExist(req.WalletID)
	c.JSON(http.StatusOK, gin.H{"exists": exists})
}

func (h *WalletHandler) TopUp(ctx *gin.Context) {
	var req struct {
		WalletID uint    `json:"wallet_id"`
		Amount   float64 `json:"amount"`
	}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	if err := h.service.TopUp(req.WalletID, req.Amount); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "wallet topped up"})
}

func (h *WalletHandler) GetStats(c *gin.Context) {
	var req struct {
		WalletID uint `json:"wallet_id"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}
	count, total, err := h.service.GetMonthlyStats(req.WalletID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"count": count, "total_amount": total})
}

func (h *WalletHandler) GetBalance(c *gin.Context) {
	var req struct {
		WalletID uint `json:"wallet_id"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}
	balance, err := h.service.GetBalance(req.WalletID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"balance": balance})
}
