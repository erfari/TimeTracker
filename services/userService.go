package services

import (
	"TimeTracker/models"
	"TimeTracker/repository"
	"encoding/json"
	"io"
	"log"
	"net/http"
	url "net/url"
)

type UserService struct {
	userRepository *repository.UserRepository
}

func NewUserService(
	userRepository *repository.UserRepository) *UserService {
	return &UserService{
		userRepository: userRepository,
	}
}

// можно вынести в микросервис
func getInfoUser(passportSerie string, passportNumber string) (*models.Users, *models.ResponseError) {
	params := url.Values{}
	params.Add("passportSerie", passportSerie)
	params.Add("passportNumber", passportNumber)
	urlInfo := "http://localhost:8080/info?" + params.Encode()

	response, err := http.Get(urlInfo)
	if err != nil {
		log.Fatalln(err)
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatalln(err)
	}
	var result models.Users
	if err = json.Unmarshal(body, &result); err != nil {
		log.Fatalln(err)
	}
	return &models.Users{
		Name:           result.Name,
		Surname:        result.Surname,
		Patronymic:     result.Patronymic,
		Address:        result.Address,
		PassportSerie:  result.PassportSerie,
		PassportNumber: result.PassportNumber,
	}, nil
}

/**
1. Идем за обогашщением в метод getInfoUser
2.
*/

func (us *UserService) AddUserApi(passport *models.PassportDocument) (*models.Users, *models.ResponseError) {
	passportSerialNumber := passport.PassportNumber[0:4]
	passportNumber := passport.PassportNumber[5:11]
	user, responseErr := getInfoUser(passportSerialNumber, passportNumber)
	if responseErr != nil {
		return nil, responseErr
	}
	return us.userRepository.AddUserApi(user)
}

func (us *UserService) AddUser(user *models.Users) (*models.Users, *models.ResponseError) {
	responseErr := validateUser(user)
	if responseErr != nil {
		return nil, responseErr
	}
	return us.userRepository.AddUser(user)
}

func validateUser(user *models.Users) *models.ResponseError {
	if user.Name == "" {
		return &models.ResponseError{
			Message: "Invalid Name",
			Status:  http.StatusBadRequest,
		}
	}
	if user.Surname == "" {
		return &models.ResponseError{
			Message: "Invalid Surname",
			Status:  http.StatusBadRequest,
		}
	}
	if user.PassportSerie == "" {
		return &models.ResponseError{
			Message: "Invalid Passport Serial Number",
			Status:  http.StatusBadRequest,
		}
	}
	if user.PassportNumber == "" {
		return &models.ResponseError{
			Message: "Invalid Passport Passport Number",
			Status:  http.StatusBadRequest,
		}
	}
	return nil
}

func (us *UserService) UpdateUser(user *models.Users) *models.ResponseError {
	responseErr := validateUser(user)
	if responseErr != nil {
		return responseErr
	}
	return us.userRepository.UpdateUser(user)
}

func (us *UserService) DeleteUser(userId string) *models.ResponseError {
	if userId == "" {
		return &models.ResponseError{
			Message: "Invalid result ID",
			Status:  http.StatusBadRequest,
		}
	}
	err := repository.BeginTransaction(us.userRepository)
	if err != nil {
		return &models.ResponseError{
			Message: "Failed to start transaction",
			Status:  http.StatusBadRequest,
		}
	}
	//Добавить роллбек
	responseErr := us.userRepository.DeleteUser(userId)
	if responseErr != nil {
		return responseErr
	}
	repository.CommitTransaction(us.userRepository)
	return nil
}

func (us *UserService) GetUser(userId string) (*models.Users, *models.ResponseError) {
	user, responseErr := us.userRepository.GetUser(userId)
	if responseErr != nil {
		return nil, responseErr
	}
	return user, nil
}

// TODO тут фильтрацию сделать по-человечески
func (us *UserService) GetUsersBach(limit int, offset int, name, surname, patronimyc, address, passportSerialNumber, passportNumber string) ([]*models.Users, *models.ResponseError) {
	if name != "" {
		users, responseErr := us.userRepository.GetUserByParams(limit, offset, name)
		if responseErr != nil {
			return nil, responseErr
		}
		return users, nil
	}
	if surname != "" {
		users, responseErr := us.userRepository.GetUserByParams(limit, offset, surname)
		if responseErr != nil {
			return nil, responseErr
		}
		return users, nil
	}
	if patronimyc != "" {
		users, responseErr := us.userRepository.GetUserByParams(limit, offset, patronimyc)
		if responseErr != nil {
			return nil, responseErr
		}
		return users, nil
	}
	if address != "" {
		users, responseErr := us.userRepository.GetUserByParams(limit, offset, address)
		if responseErr != nil {
			return nil, responseErr
		}
		return users, nil
	}
	if passportSerialNumber != "" {
		users, responseErr := us.userRepository.GetUserByParams(limit, offset, passportSerialNumber)
		if responseErr != nil {
			return nil, responseErr
		}
		return users, nil
	}
	if passportNumber != "" {
		users, responseErr := us.userRepository.GetUserByParams(limit, offset, passportNumber)
		if responseErr != nil {
			return nil, responseErr
		}
		return users, nil
	}
	users, responseErr := us.userRepository.GetAllUsers(limit, offset)
	if responseErr != nil {
		return nil, responseErr
	}
	return users, nil
}

func (us *UserService) Info(passportSerial string, passportNumber string) (*models.Users, *models.ResponseError) {
	user, responseErr := us.userRepository.Info(passportSerial, passportNumber)
	if responseErr != nil {
		return nil, responseErr
	}
	return user, nil
}
