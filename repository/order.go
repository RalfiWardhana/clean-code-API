package repository

import (
	"log"
	"rumah-sakit/model/dto"
	"rumah-sakit/model/entity"

	"gorm.io/gorm"
)

type OrderRepository interface {
	GetAllOrder() ([]entity.Order, error)
	GetOrder(id int) ([]entity.Order, error)
	CreateOrder(order_request dto.Order_request) (entity.Order, error)
	UpdateOrder(order_request dto.Order_request) (entity.Order, error)
	DeleteOrder(id int) error
}

func CreateOrderRepository(MySQLConn *gorm.DB) OrderRepository {
	return &orderRepository{
		MySQL: MySQLConn,
	}
}

type orderRepository struct {
	MySQL *gorm.DB
}

func (db *orderRepository) GetAllOrder() ([]entity.Order, error) {
	order_list := []entity.Order{}

	// Melakukan join dengan tabel Dokter dan patient
	if err := db.MySQL.Joins("JOIN doctors ON orders.id_doctor = doctors.id").
		Joins("JOIN patients ON orders.id_patient = patients.id").
		Select("orders.*, doctors.Doctor_name,patients.Patient_name").
		Order("orders.id DESC").
		Find(&order_list).Error; err != nil {
		log.Println("Get all order failed")
		return []entity.Order{}, err
	}

	return order_list, nil

}

func (db *orderRepository) GetOrder(id int) ([]entity.Order, error) {
	order_list := []entity.Order{}

	queryBuilder := db.MySQL.Model(entity.Order{})

	queryBuilder = queryBuilder.Where("id = ?", id)

	if err := queryBuilder.Order("id DESC").Scan(&order_list).Error; err != nil {
		log.Println("Get order failed")

		return order_list, err
	}

	return order_list, nil

}

func (db *orderRepository) CreateOrder(order_request dto.Order_request) (entity.Order, error) {
	order := entity.Order{
		Id_patient:       order_request.Id_patient,
		Id_doctor:        order_request.Id_doctor,
		Waktu:            order_request.Waktu,
		Harga_konsultasi: order_request.Harga_konsultasi,
		Harga_obat:       order_request.Harga_obat,
		Total_harga:      order_request.Total_harga,
	}
	if err := db.MySQL.Model(entity.Order{}).Create(&order_request).Error; err != nil {
		log.Println("Store order failed")

		return entity.Order{}, err
	}

	return order, nil
}

func (db *orderRepository) UpdateOrder(order_request dto.Order_request) (entity.Order, error) {

	update_order := entity.Order{
		Id_patient:       order_request.Id_patient,
		Id_doctor:        order_request.Id_doctor,
		Waktu:            order_request.Waktu,
		Harga_konsultasi: order_request.Harga_konsultasi,
		Harga_obat:       order_request.Harga_obat,
		Total_harga:      order_request.Total_harga,
	}

	if err := db.MySQL.Model(entity.Order{}).Where("id = ?", order_request.Id).Updates(&update_order).Error; err != nil {
		log.Println("Update order failed")

		return entity.Order{}, err
	}

	return update_order, nil

}

func (db *orderRepository) DeleteOrder(id int) error {

	if err := db.MySQL.Model(entity.Order{}).Where("id = ?", id).Delete(&entity.Order{}).Error; err != nil {
		log.Println("Delete order failed")

		return err
	}

	return nil

}
