package usacase

import (
	"rumah-sakit/model/dto"
	"rumah-sakit/repository"
)

type UserUseCase interface {
	GetAllUser() ([]dto.User_response, error)
	GetUser(id int) ([]dto.User_response, error)
	Login(user_request dto.User_request) (dto.User_response, error)
	CreateUser(user_request dto.User_request) (dto.User_response, error)
	UpdateUser(user_request dto.User_request) (dto.User_response, error)
	DeleteUser(id int) error
}

func CreateUserUseCase(repository repository.UserRepository) UserUseCase {
	return &userUseCase{
		userRepository: repository,
	}
}

type userUseCase struct {
	userRepository repository.UserRepository
}

func (user *userUseCase) GetAllUser() ([]dto.User_response, error) {

	var response []dto.User_response
	var resp dto.User_response
	user_list, err := user.userRepository.GetAllUser()

	if err != nil {
		return response, err
	}

	for _, data := range user_list {
		resp = dto.User_response{
			Email: data.Email,
			Name:  data.Name,
		}

		response = append(response, resp)
	}

	return response, nil

}

func (user *userUseCase) GetUser(id int) ([]dto.User_response, error) {

	var response []dto.User_response
	var resp dto.User_response
	user_list, err := user.userRepository.GetUser(id)

	if err != nil {
		return response, err
	}

	for _, data := range user_list {
		resp = dto.User_response{
			Email: data.Email,
			Name:  data.Name,
		}

		response = append(response, resp)
	}

	return response, nil

}

func (user *userUseCase) Login(user_request dto.User_request) (dto.User_response, error) {

	var resp dto.User_response
	user_login, err := user.userRepository.Login(user_request)

	if err != nil {
		return resp, err
	}

	resp = dto.User_response{
		Email: user_login.Email,
		Name:  user_login.Name,
	}

	return resp, nil

}

func (user *userUseCase) CreateUser(user_request dto.User_request) (dto.User_response, error) {

	var resp dto.User_response
	user_create, err := user.userRepository.CreateUser(user_request)

	if err != nil {
		return resp, err
	}

	resp = dto.User_response{
		Id:    user_request.Id,
		Email: user_create.Email,
		Name:  user_create.Name,
	}

	return resp, nil

}

func (user *userUseCase) UpdateUser(user_request dto.User_request) (dto.User_response, error) {

	var resp dto.User_response
	user_create, err := user.userRepository.UpdateUser(user_request)

	if err != nil {
		return resp, err
	}

	resp = dto.User_response{
		Email: user_create.Email,
		Name:  user_create.Name,
	}

	return resp, nil

}

func (user *userUseCase) DeleteUser(id int) error {

	err := user.userRepository.DeleteUser(id)

	if err != nil {
		return err
	}

	return nil

}
