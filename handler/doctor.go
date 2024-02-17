package handler

import (
	"log"
	"net/http"
	"rumah-sakit/model/dto"
	usacase "rumah-sakit/usecase"
	"strconv"

	"github.com/gin-gonic/gin"
)

type DoctorHandler interface {
	GetAllDoctor(c *gin.Context)
	GetDoctor(c *gin.Context)
	CreateDoctor(c *gin.Context)
	UpdateDoctor(c *gin.Context)
	DeleteDoctor(c *gin.Context)
}

func CreateDoctorHandler(usecase usacase.DoctorUseCase) DoctorHandler {
	return &doctorHandler{
		doctorUseCase: usecase,
	}
}

type doctorHandler struct {
	doctorUseCase usacase.DoctorUseCase
}

func (doctor *doctorHandler) GetAllDoctor(c *gin.Context) {

	// Memanggil use case untuk mendapatkan semua doctor
	doctors, err := doctor.doctorUseCase.GetAllDoctor()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Mengembalikan data doctor dalam format JSON
	c.JSON(200, map[string]any{
		"Data":        doctors,
		"Status Code": 200,
		"Message":     "Success get all doctor",
	})
}

func (doctor *doctorHandler) GetDoctor(c *gin.Context) {

	id := c.Param("id")

	doctorID, err := strconv.Atoi(id)

	if err != nil {
		log.Println("Invalid doctor ID:", err)
	}

	// Memanggil use case untuk mendapatkan doctor
	get_doctor, err := doctor.doctorUseCase.GetDoctor(doctorID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Mengembalikan data doctor dalam format JSON
	c.JSON(200, map[string]any{
		"Data":        get_doctor,
		"Status Code": 200,
		"Message":     "Success get doctor",
	})
}

func (doctor *doctorHandler) CreateDoctor(c *gin.Context) {

	var request_doctor dto.Doctor_request

	if err := c.ShouldBind(&request_doctor); err != nil {
		c.JSON(400, map[string]string{
			"message": "invalid input",
		})
		return
	}

	// Memanggil use case untuk mendapatkan doctor
	create_doctor, err := doctor.doctorUseCase.CreateDoctor(request_doctor)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Mengembalikan data doctor dalam format JSON
	c.JSON(200, map[string]any{
		"Data":        create_doctor,
		"Status Code": 201,
		"Message":     "Success create doctor",
	})
}

func (doctor *doctorHandler) UpdateDoctor(c *gin.Context) {

	var request_doctor dto.Doctor_request

	if err := c.ShouldBind(&request_doctor); err != nil {
		c.JSON(400, map[string]string{
			"message": "invalid input",
		})
		return
	}

	// Memanggil use case untuk mendapatkan doctor
	update_doctor, err := doctor.doctorUseCase.UpdateDoctor(request_doctor)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Mengembalikan data doctor dalam format JSON
	c.JSON(200, map[string]any{
		"Data":        update_doctor,
		"Status Code": 200,
		"Message":     "Success update doctor",
	})
}

func (doctor *doctorHandler) DeleteDoctor(c *gin.Context) {

	id := c.Param("id")

	doctorID, err := strconv.Atoi(id)

	if err != nil {
		log.Println("Invalid doctor ID:", err)
	}

	// Memanggil use case untuk mendapatkan doctor
	err = doctor.doctorUseCase.DeleteDoctor(doctorID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Mengembalikan data doctor dalam format JSON
	c.JSON(200, map[string]any{
		"Status Code": 200,
		"Message":     "Success delete doctor",
	})
}
