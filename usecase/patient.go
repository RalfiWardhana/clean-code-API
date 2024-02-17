package usacase

import (
	"rumah-sakit/model/dto"
	"rumah-sakit/repository"
)

type PatientUseCase interface {
	GetAllPatient() ([]dto.Patient_response, error)
	GetPatient(id int) ([]dto.Patient_response, error)
	CreatePatient(patient_request dto.Patient_request) (dto.Patient_response, error)
	UpdatePatient(patient_request dto.Patient_request) (dto.Patient_response, error)
	DeletePatient(id int) error
}

func CreatePatientUseCase(repository repository.PatientRepository) PatientUseCase {
	return &patientUseCase{
		patientRepository: repository,
	}
}

type patientUseCase struct {
	patientRepository repository.PatientRepository
}

func (patient *patientUseCase) GetAllPatient() ([]dto.Patient_response, error) {

	var response []dto.Patient_response
	var resp dto.Patient_response
	patient_list, err := patient.patientRepository.GetAllPatient()

	if err != nil {
		return response, err
	}

	for _, data := range patient_list {
		resp = dto.Patient_response{
			Patient_name: data.Patient_name,
			Umur:         data.Umur,
			Alamat:       data.Alamat,
		}

		response = append(response, resp)
	}

	return response, nil

}

func (patient *patientUseCase) GetPatient(id int) ([]dto.Patient_response, error) {

	var response []dto.Patient_response
	var resp dto.Patient_response
	patient_list, err := patient.patientRepository.GetPatient(id)

	if err != nil {
		return response, err
	}

	for _, data := range patient_list {
		resp = dto.Patient_response{
			Patient_name: data.Patient_name,
			Umur:         data.Umur,
			Alamat:       data.Alamat,
		}

		response = append(response, resp)
	}

	return response, nil

}

func (patient *patientUseCase) CreatePatient(patient_request dto.Patient_request) (dto.Patient_response, error) {

	var resp dto.Patient_response
	patient_create, err := patient.patientRepository.CreatePatient(patient_request)

	if err != nil {
		return resp, err
	}

	resp = dto.Patient_response{
		Patient_name: patient_create.Patient_name,
		Umur:         patient_create.Umur,
		Alamat:       patient_create.Alamat,
	}

	return resp, nil

}

func (patient *patientUseCase) UpdatePatient(patient_request dto.Patient_request) (dto.Patient_response, error) {

	var resp dto.Patient_response
	patient_create, err := patient.patientRepository.UpdatePatient(patient_request)

	if err != nil {
		return resp, err
	}

	resp = dto.Patient_response{
		Patient_name: patient_create.Patient_name,
		Umur:         patient_create.Umur,
		Alamat:       patient_create.Alamat,
	}

	return resp, nil

}

func (patient *patientUseCase) DeletePatient(id int) error {

	err := patient.patientRepository.DeletePatient(id)

	if err != nil {
		return err
	}

	return nil

}
