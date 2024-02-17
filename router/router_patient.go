package routes

import (
	config "rumah-sakit/config"
	handler "rumah-sakit/handler"
	usecase "rumah-sakit/usecase"
	gate "rumah-sakit/middleware"

	repo "rumah-sakit/repository"

	"github.com/gin-gonic/gin"
)

func PatientRoute(router *gin.Engine) {

	db := config.ConnectDB()
	repoPatient := repo.CreatePatientRepository(db)
	useCasePatient := usecase.CreatePatientUseCase(repoPatient)
	hanlderPatient := handler.CreatePatientHandler(useCasePatient)

	router.GET("/patients", gate.WithAuthentication(), hanlderPatient.GetAllPatient)
	router.GET("/patient/:id", gate.WithAuthentication(), hanlderPatient.GetPatient)
	router.POST("/patient", gate.WithAuthentication(), hanlderPatient.CreatePatient)
	router.PUT("/patient", gate.WithAuthentication(), hanlderPatient.UpdatePatient)
	router.DELETE("/patient/:id", gate.WithAuthentication(), hanlderPatient.DeletePatient)
}
