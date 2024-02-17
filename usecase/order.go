package usacase

import (
	"rumah-sakit/model/dto"
	"rumah-sakit/repository"
)

type OrderUseCase interface {
	GetAllOrder() ([]dto.Order_response, error)
	GetOrder(id int) ([]dto.Order_response, error)
	CreateOrder(order_request dto.Order_request) (dto.Order_response, error)
	UpdateOrder(order_request dto.Order_request) (dto.Order_response, error)
	DeleteOrder(id int) error
}

func CreateOrderUseCase(repository repository.OrderRepository) OrderUseCase {
	return &orderUseCase{
		orderRepository: repository,
	}
}

type orderUseCase struct {
	orderRepository repository.OrderRepository
}

func (order *orderUseCase) GetAllOrder() ([]dto.Order_response, error) {

	var response []dto.Order_response
	var resp dto.Order_response
	order_list, err := order.orderRepository.GetAllOrder()

	if err != nil {
		return response, err
	}

	for _, data := range order_list {
		resp = dto.Order_response{
			Patient_name:     data.Patient_name,
			Doctor_name:      data.Doctor_name,
			Waktu:            data.Waktu,
			Harga_konsultasi: data.Harga_konsultasi,
			Harga_obat:       data.Harga_obat,
			Total_harga:      data.Total_harga,
		}

		response = append(response, resp)
	}

	return response, nil

}

func (order *orderUseCase) GetOrder(id int) ([]dto.Order_response, error) {

	var response []dto.Order_response
	var resp dto.Order_response
	order_list, err := order.orderRepository.GetOrder(id)

	if err != nil {
		return response, err
	}

	for _, data := range order_list {
		resp = dto.Order_response{
			Patient_name:     data.Patient_name,
			Doctor_name:      data.Doctor_name,
			Waktu:            data.Waktu,
			Harga_konsultasi: data.Harga_konsultasi,
			Harga_obat:       data.Harga_obat,
			Total_harga:      data.Total_harga,
		}

		response = append(response, resp)
	}

	return response, nil

}

func (order *orderUseCase) CreateOrder(order_request dto.Order_request) (dto.Order_response, error) {

	var resp dto.Order_response
	order_create, err := order.orderRepository.CreateOrder(order_request)

	if err != nil {
		return resp, err
	}

	resp = dto.Order_response{
		Id_patient:       order_create.Id_patient,
		Id_doctor:        order_create.Id_doctor,
		Waktu:            order_create.Waktu,
		Harga_konsultasi: order_create.Harga_konsultasi,
		Harga_obat:       order_create.Harga_obat,
		Total_harga:      order_create.Total_harga,
	}

	return resp, nil

}

func (order *orderUseCase) UpdateOrder(order_request dto.Order_request) (dto.Order_response, error) {

	var resp dto.Order_response
	order_create, err := order.orderRepository.UpdateOrder(order_request)

	if err != nil {
		return resp, err
	}

	resp = dto.Order_response{
		Patient_name:     order_create.Patient_name,
		Doctor_name:      order_create.Doctor_name,
		Waktu:            order_create.Waktu,
		Harga_konsultasi: order_create.Harga_konsultasi,
		Harga_obat:       order_create.Harga_obat,
		Total_harga:      order_create.Total_harga,
	}

	return resp, nil

}

func (order *orderUseCase) DeleteOrder(id int) error {

	err := order.orderRepository.DeleteOrder(id)

	if err != nil {
		return err
	}

	return nil

}
