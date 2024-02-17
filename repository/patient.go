package repository

import (
	"log"
	"rumah-sakit/model/dto"
	"rumah-sakit/model/entity"

	"gorm.io/gorm"
)

type PatientRepository interface {
	GetAllPatient() ([]entity.Patient, error)
	GetPatient(id int) ([]entity.Patient, error)
	CreatePatient(patient_request dto.Patient_request) (entity.Patient, error)
	UpdatePatient(patient_request dto.Patient_request) (entity.Patient, error)
	DeletePatient(id int) error
}

func CreatePatientRepository(MySQLConn *gorm.DB) PatientRepository {
	return &patientRepository{
		MySQL: MySQLConn,
	}
}

type patientRepository struct {
	MySQL *gorm.DB
}

func (db *patientRepository) GetAllPatient() ([]entity.Patient, error) {
	patient_list := []entity.Patient{}

	if err := db.MySQL.Model(entity.Patient{}).
		Order("id DESC").
		Scan(&patient_list).Error; err != nil {
		log.Println("Get all patient failed")

		return []entity.Patient{}, err
	}

	return patient_list, nil

}

func (db *patientRepository) GetPatient(id int) ([]entity.Patient, error) {
	patient_list := []entity.Patient{}

	queryBuilder := db.MySQL.Model(entity.Patient{})

	queryBuilder = queryBuilder.Where("id = ?", id)

	if err := queryBuilder.Order("id DESC").Scan(&patient_list).Error; err != nil {
		return patient_list, err
	}

	return patient_list, nil

}

func (db *patientRepository) CreatePatient(patient_request dto.Patient_request) (entity.Patient, error) {
	patient := entity.Patient{
		Patient_name: patient_request.Patient_name,
		Umur:         patient_request.Umur,
		Alamat:       patient_request.Alamat,
	}
	if err := db.MySQL.Model(entity.Patient{}).Create(&patient_request).Error; err != nil {
		log.Println("Store patient failed")

		return entity.Patient{}, err
	}

	return patient, nil
}

func (db *patientRepository) UpdatePatient(patient_request dto.Patient_request) (entity.Patient, error) {

	update_patient := entity.Patient{
		Patient_name: patient_request.Patient_name,
		Umur:         patient_request.Umur,
		Alamat:       patient_request.Alamat,
	}

	if err := db.MySQL.Model(entity.Patient{}).Where("id = ?", patient_request.Id).Updates(&update_patient).Error; err != nil {
		log.Println("Update patient failed")

		return entity.Patient{}, err
	}

	return update_patient, nil

}

func (db *patientRepository) DeletePatient(id int) error {

	if err := db.MySQL.Model(entity.Patient{}).Where("id = ?", id).Delete(&entity.Patient{}).Error; err != nil {
		log.Println("Delete patient failed")

		return err
	}

	return nil

}
