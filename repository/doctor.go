package repository

import (
	"log"
	"rumah-sakit/model/dto"
	"rumah-sakit/model/entity"

	"gorm.io/gorm"
)

type DoctorRepository interface {
	GetAllDoctor() ([]entity.Doctor, error)
	GetDoctor(id int) ([]entity.Doctor, error)
	CreateDoctor(doctor_request dto.Doctor_request) (entity.Doctor, error)
	UpdateDoctor(doctor_request dto.Doctor_request) (entity.Doctor, error)
	DeleteDoctor(id int) error
}

func CreateDoctorRepository(MySQLConn *gorm.DB) DoctorRepository {
	return &doctorRepository{
		MySQL: MySQLConn,
	}
}

type doctorRepository struct {
	MySQL *gorm.DB
}

func (db *doctorRepository) GetAllDoctor() ([]entity.Doctor, error) {
	doctor_list := []entity.Doctor{}

	if err := db.MySQL.Model(entity.Doctor{}).
		Order("id DESC").
		Scan(&doctor_list).Error; err != nil {
		log.Println("Get all doctor failed")

		return []entity.Doctor{}, err
	}

	return doctor_list, nil

}

func (db *doctorRepository) GetDoctor(id int) ([]entity.Doctor, error) {
	doctor_list := []entity.Doctor{}

	queryBuilder := db.MySQL.Model(entity.Doctor{})

	queryBuilder = queryBuilder.Where("id = ?", id)

	if err := queryBuilder.Order("id DESC").Scan(&doctor_list).Error; err != nil {
		log.Println("Get doctor failed")
		return doctor_list, err
	}

	return doctor_list, nil

}

func (db *doctorRepository) CreateDoctor(doctor_request dto.Doctor_request) (entity.Doctor, error) {
	doctor := entity.Doctor{
		Doctor_name:      doctor_request.Doctor_name,
		Jadwal_praktek:   doctor_request.Jadwal_praktek,
		Harga_konsultasi: doctor_request.Harga_konsultasi,
	}
	if err := db.MySQL.Model(entity.Doctor{}).Create(&doctor_request).Error; err != nil {
		log.Println("Store doctor failed")

		return entity.Doctor{}, err
	}

	return doctor, nil
}

func (db *doctorRepository) UpdateDoctor(doctor_request dto.Doctor_request) (entity.Doctor, error) {

	update_doctor := entity.Doctor{
		Doctor_name:      doctor_request.Doctor_name,
		Jadwal_praktek:   doctor_request.Jadwal_praktek,
		Harga_konsultasi: doctor_request.Harga_konsultasi,
	}

	if err := db.MySQL.Model(entity.Doctor{}).Where("id = ?", doctor_request.Id).Updates(&update_doctor).Error; err != nil {
		log.Println("Update doctor failed")

		return entity.Doctor{}, err
	}

	return update_doctor, nil

}

func (db *doctorRepository) DeleteDoctor(id int) error {

	if err := db.MySQL.Model(entity.Doctor{}).Where("id = ?", id).Delete(&entity.Doctor{}).Error; err != nil {
		log.Println("Delete doctor failed")

		return err
	}

	return nil

}
