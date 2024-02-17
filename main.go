package main

import (
	"log"
	"rumah-sakit/config"
	routing "rumah-sakit/router"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	cfg := config.ConnectDB()
	log.Println(cfg)

	routing.PatientRoute(router)
	routing.DoctorRoute(router)
	routing.OrderRoute(router)
	routing.UserRoute(router)

	router.Run("localhost:9000")

}
