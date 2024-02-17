package routes

import (
	config "rumah-sakit/config"
	handler "rumah-sakit/handler"
	usecase "rumah-sakit/usecase"
	gate "rumah-sakit/middleware"

	repo "rumah-sakit/repository"

	"github.com/gin-gonic/gin"
)

func DoctorRoute(router *gin.Engine) {

	db := config.ConnectDB()
	repoDoctor := repo.CreateDoctorRepository(db)
	useCaseDoctor := usecase.CreateDoctorUseCase(repoDoctor)
	hanlderDoctor := handler.CreateDoctorHandler(useCaseDoctor)

	router.GET("/doctors", gate.WithAuthentication(), hanlderDoctor.GetAllDoctor)
	router.GET("/doctor/:id", gate.WithAuthentication(), hanlderDoctor.GetDoctor)
	router.POST("/doctor", gate.WithAuthentication(), hanlderDoctor.CreateDoctor)
	router.PUT("/doctor", gate.WithAuthentication(), hanlderDoctor.UpdateDoctor)
	router.DELETE("/doctor/:id", gate.WithAuthentication(), hanlderDoctor.DeleteDoctor)
}
