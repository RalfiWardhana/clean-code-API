package repository

import (
	"log"
	"rumah-sakit/model/dto"
	"rumah-sakit/model/entity"

	"gorm.io/gorm"
)

type UserRepository interface {
	GetAllUser() ([]entity.User, error)
	GetUser(id int) ([]entity.User, error)
	Login(user_request dto.User_request) (entity.User, error)
	CreateUser(user_request dto.User_request) (entity.User, error)
	UpdateUser(user_request dto.User_request) (entity.User, error)
	DeleteUser(id int) error
}

func CreateUserRepository(MySQLConn *gorm.DB) UserRepository {
	return &userRepository{
		MySQL: MySQLConn,
	}
}

type userRepository struct {
	MySQL *gorm.DB
}

func (db *userRepository) GetAllUser() ([]entity.User, error) {
	user_list := []entity.User{}

	if err := db.MySQL.Model(entity.User{}).
		Order("id DESC").
		Scan(&user_list).Error; err != nil {
		log.Println("Get all User failed")

		return []entity.User{}, err
	}

	return user_list, nil

}

func (db *userRepository) GetUser(id int) ([]entity.User, error) {
	user_list := []entity.User{}

	queryBuilder := db.MySQL.Model(entity.User{})

	queryBuilder = queryBuilder.Where("id = ?", id)

	if err := queryBuilder.Order("id DESC").Scan(&user_list).Error; err != nil {
		return user_list, err
	}

	return user_list, nil

}

func (db *userRepository) Login(user_request dto.User_request) (entity.User, error) {

	user_list := entity.User{}

	queryBuilder := db.MySQL.Model(entity.User{})

	queryBuilder = queryBuilder.Where("email = ?", user_request.Email)
	queryBuilder = queryBuilder.Where("password = ?", user_request.Password)

	if err := queryBuilder.Order("id DESC").Scan(&user_list).Error; err != nil {
		return user_list, err
	}

	return user_list, nil

}

func (db *userRepository) CreateUser(user_request dto.User_request) (entity.User, error) {
	user := entity.User{
		Email:    user_request.Email,
		Password: user_request.Password,
		Name:     user_request.Name,
	}
	if err := db.MySQL.Model(entity.User{}).Create(&user_request).Error; err != nil {
		log.Println("Store User failed")

		return entity.User{}, err
	}

	return user, nil
}

func (db *userRepository) UpdateUser(user_request dto.User_request) (entity.User, error) {

	update_User := entity.User{
		Email:    user_request.Email,
		Password: user_request.Password,
		Name:     user_request.Name,
	}

	if err := db.MySQL.Model(entity.User{}).Where("id = ?", user_request.Id).Updates(&update_User).Error; err != nil {
		log.Println("Update User failed")

		return entity.User{}, err
	}

	return update_User, nil

}

func (db *userRepository) DeleteUser(id int) error {

	if err := db.MySQL.Model(entity.User{}).Where("id = ?", id).Delete(&entity.User{}).Error; err != nil {
		log.Println("Delete User failed")

		return err
	}

	return nil

}
