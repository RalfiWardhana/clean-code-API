package routes

import (
	config "rumah-sakit/config"
	handler "rumah-sakit/handler"
	usecase "rumah-sakit/usecase"
	gate "rumah-sakit/middleware"

	repo "rumah-sakit/repository"

	"github.com/gin-gonic/gin"
)

func OrderRoute(router *gin.Engine) {

	db := config.ConnectDB()
	repoOrder := repo.CreateOrderRepository(db)
	useCaseOrder := usecase.CreateOrderUseCase(repoOrder)
	hanlderOrder := handler.CreateOrderHandler(useCaseOrder)

	router.GET("/orders", gate.WithAuthentication(), hanlderOrder.GetAllOrder)
	router.GET("/order/:id", gate.WithAuthentication(), hanlderOrder.GetOrder)
	router.POST("/order", gate.WithAuthentication(), hanlderOrder.CreateOrder)
	router.PUT("/order", gate.WithAuthentication(), hanlderOrder.UpdateOrder)
	router.DELETE("/order/:id", gate.WithAuthentication(), hanlderOrder.DeleteOrder)
}
