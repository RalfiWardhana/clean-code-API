package usacase

import (
	"rumah-sakit/model/dto"
	"rumah-sakit/repository"
)

type DoctorUseCase interface {
	GetAllDoctor() ([]dto.Doctor_response, error)
	GetDoctor(id int) ([]dto.Doctor_response, error)
	CreateDoctor(doctor_request dto.Doctor_request) (dto.Doctor_response, error)
	UpdateDoctor(doctor_request dto.Doctor_request) (dto.Doctor_response, error)
	DeleteDoctor(id int) error
}

func CreateDoctorUseCase(repository repository.DoctorRepository) DoctorUseCase {
	return &doctorUseCase{
		doctorRepository: repository,
	}
}

type doctorUseCase struct {
	doctorRepository repository.DoctorRepository
}

func (doctor *doctorUseCase) GetAllDoctor() ([]dto.Doctor_response, error) {

	var response []dto.Doctor_response
	var resp dto.Doctor_response
	doctor_list, err := doctor.doctorRepository.GetAllDoctor()

	if err != nil {
		return response, err
	}

	for _, data := range doctor_list {
		resp = dto.Doctor_response{
			Doctor_name:      data.Doctor_name,
			Jadwal_praktek:   data.Jadwal_praktek,
			Harga_konsultasi: data.Harga_konsultasi,
		}

		response = append(response, resp)
	}

	return response, nil

}

func (doctor *doctorUseCase) GetDoctor(id int) ([]dto.Doctor_response, error) {

	var response []dto.Doctor_response
	var resp dto.Doctor_response
	doctor_list, err := doctor.doctorRepository.GetDoctor(id)

	if err != nil {
		return response, err
	}

	for _, data := range doctor_list {
		resp = dto.Doctor_response{
			Doctor_name:      data.Doctor_name,
			Jadwal_praktek:   data.Jadwal_praktek,
			Harga_konsultasi: data.Harga_konsultasi,
		}

		response = append(response, resp)
	}

	return response, nil

}

func (doctor *doctorUseCase) CreateDoctor(doctor_request dto.Doctor_request) (dto.Doctor_response, error) {

	var resp dto.Doctor_response
	doctor_create, err := doctor.doctorRepository.CreateDoctor(doctor_request)

	if err != nil {
		return resp, err
	}

	resp = dto.Doctor_response{
		Doctor_name:      doctor_create.Doctor_name,
		Jadwal_praktek:   doctor_create.Jadwal_praktek,
		Harga_konsultasi: doctor_create.Harga_konsultasi,
	}

	return resp, nil

}

func (doctor *doctorUseCase) UpdateDoctor(doctor_request dto.Doctor_request) (dto.Doctor_response, error) {

	var resp dto.Doctor_response
	doctor_create, err := doctor.doctorRepository.UpdateDoctor(doctor_request)

	if err != nil {
		return resp, err
	}

	resp = dto.Doctor_response{
		Doctor_name:      doctor_create.Doctor_name,
		Jadwal_praktek:   doctor_create.Jadwal_praktek,
		Harga_konsultasi: doctor_create.Harga_konsultasi,
	}

	return resp, nil

}

func (doctor *doctorUseCase) DeleteDoctor(id int) error {

	err := doctor.doctorRepository.DeleteDoctor(id)

	if err != nil {
		return err
	}

	return nil

}
