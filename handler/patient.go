package handler

import (
	"log"
	"net/http"
	"rumah-sakit/model/dto"
	usacase "rumah-sakit/usecase"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PatientHandler interface {
	GetAllPatient(c *gin.Context)
	GetPatient(c *gin.Context)
	CreatePatient(c *gin.Context)
	UpdatePatient(c *gin.Context)
	DeletePatient(c *gin.Context)
}

func CreatePatientHandler(usecase usacase.PatientUseCase) PatientHandler {
	return &patientHandler{
		patientUseCase: usecase,
	}
}

type patientHandler struct {
	patientUseCase usacase.PatientUseCase
}

func (patient *patientHandler) GetAllPatient(c *gin.Context) {

	// Memanggil use case untuk mendapatkan semua patient
	patients, err := patient.patientUseCase.GetAllPatient()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Mengembalikan data patient dalam format JSON
	c.JSON(200, map[string]any{
		"Data":        patients,
		"Status Code": 200,
		"Message":     "Success get all patient",
	})
}

func (patient *patientHandler) GetPatient(c *gin.Context) {

	id := c.Param("id")

	patientID, err := strconv.Atoi(id)

	if err != nil {
		log.Println("Invalid patient ID:", err)
	}

	// Memanggil use case untuk mendapatkan patient
	get_patient, err := patient.patientUseCase.GetPatient(patientID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Mengembalikan data patient dalam format JSON
	c.JSON(200, map[string]any{
		"Data":        get_patient,
		"Status Code": 200,
		"Message":     "Success get patient",
	})
}

func (patient *patientHandler) CreatePatient(c *gin.Context) {

	var request_patient dto.Patient_request

	if err := c.ShouldBind(&request_patient); err != nil {
		c.JSON(400, map[string]string{
			"message": "invalid input",
		})
		return
	}

	// Memanggil use case untuk mendapatkan patient
	create_patient, err := patient.patientUseCase.CreatePatient(request_patient)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Mengembalikan data patient dalam format JSON
	c.JSON(200, map[string]any{
		"Data":        create_patient,
		"Status Code": 201,
		"Message":     "Success create patient",
	})
}

func (patient *patientHandler) UpdatePatient(c *gin.Context) {

	var request_patient dto.Patient_request

	if err := c.ShouldBind(&request_patient); err != nil {
		c.JSON(400, map[string]string{
			"message": "invalid input",
		})
		return
	}

	// Memanggil use case untuk mendapatkan patient
	update_patient, err := patient.patientUseCase.UpdatePatient(request_patient)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Mengembalikan data patient dalam format JSON
	c.JSON(200, map[string]any{
		"Data":        update_patient,
		"Status Code": 200,
		"Message":     "Success update patient",
	})
}

func (patient *patientHandler) DeletePatient(c *gin.Context) {

	id := c.Param("id")

	patientID, err := strconv.Atoi(id)

	if err != nil {
		log.Println("Invalid patient ID:", err)
	}

	// Memanggil use case untuk mendapatkan patient
	err = patient.patientUseCase.DeletePatient(patientID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Mengembalikan data patient dalam format JSON
	c.JSON(200, map[string]any{
		"Status Code": 200,
		"Message":     "Success delete patient",
	})
}
