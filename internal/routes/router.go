package routes

import (
	"github.com/gin-gonic/gin"
	"go.mod/internal/handler"
	"go.mod/internal/middleware"
	"go.mod/internal/repository"
	"go.mod/internal/service"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	repo := &repository.WalletRepository{}
	svc := service.NewWalletService(repo)
	handler := handler.NewWalletHandler(svc)

	r.Use(middleware.AuthMiddleware())

	r.POST("/wallet/check", handler.CheckWallet)
	r.POST("/wallet/topup", handler.TopUp)
	r.POST("/wallet/stats", handler.GetStats)
	r.POST("/wallet/balance", handler.GetBalance)

	return r
}
