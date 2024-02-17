package routes

import (
	config "rumah-sakit/config"
	handler "rumah-sakit/handler"
	gate "rumah-sakit/middleware"
	usecase "rumah-sakit/usecase"

	repo "rumah-sakit/repository"

	"github.com/gin-gonic/gin"
)

func UserRoute(router *gin.Engine) {

	db := config.ConnectDB()
	repoUser := repo.CreateUserRepository(db)
	useCaseUser := usecase.CreateUserUseCase(repoUser)
	hanlderUser := handler.CreateUserHandler(useCaseUser)

	router.GET("/users", gate.WithAuthentication(), hanlderUser.GetAllUser)
	router.GET("/user/:id", gate.WithAuthentication(), hanlderUser.GetUser)
	router.POST("/login", hanlderUser.Login)
	router.POST("/user", hanlderUser.CreateUser)
	router.PUT("/user", gate.WithAuthentication(), hanlderUser.UpdateUser)
	router.DELETE("/user/:id", gate.WithAuthentication(), hanlderUser.DeleteUser)
}
