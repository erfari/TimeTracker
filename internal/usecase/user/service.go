package user

import (
	"TimeTracker/internal/entity"
	"encoding/json"
	"io"
	"log"
	"net/http"
	url "net/url"
)

type Service struct {
	UserRepository *UserRepository
}

func NewUserService(
	UserRepository *UserRepository) *Service {
	return &Service{
		UserRepository: UserRepository,
	}
}

// можно вынести в микросервис ============================
func getInfoUser(PassportSerial string, passportNumber string) (*entity.Users, *entity.ResponseError) {
	params := url.Values{}
	params.Add("PassportSerial", PassportSerial)
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
	var result entity.Users
	if err = json.Unmarshal(body, &result); err != nil {
		log.Fatalln(err)
	}
	return &entity.Users{
		Name:           result.Name,
		Surname:        result.Surname,
		Patronymic:     result.Patronymic,
		Address:        result.Address,
		PassportSerial: result.PassportSerial,
		PassportNumber: result.PassportNumber,
	}, nil
}

/**
1. Идем за обогашщением в метод getInfoUser
2.
*/

func (us *Service) AddUserApi(passport *entity.PassportDocument) (*entity.Users, *entity.ResponseError) {
	passportSerialNumber := passport.PassportNumber[0:4]
	passportNumber := passport.PassportNumber[5:11]
	user, responseErr := getInfoUser(passportSerialNumber, passportNumber)
	if responseErr != nil {
		return nil, responseErr
	}
	return us.UserRepository.AddUserApi(user)
}

func (us *Service) AddUser(user *entity.Users) (*entity.Users, *entity.ResponseError) {
	responseErr := validateUser(user)
	if responseErr != nil {
		return nil, responseErr
	}
	return us.UserRepository.AddUser(user)
}

func validateUser(user *entity.Users) *entity.ResponseError {
	if user.Name == "" {
		return &entity.ResponseError{
			Message: "Invalid Name",
			Status:  http.StatusBadRequest,
		}
	}
	if user.Surname == "" {
		return &entity.ResponseError{
			Message: "Invalid Surname",
			Status:  http.StatusBadRequest,
		}
	}
	if user.PassportSerial == "" || len([]rune(user.PassportSerial)) > 4 {
		return &entity.ResponseError{
			Message: "Invalid Passport Serial Number",
			Status:  http.StatusBadRequest,
		}
	}
	if user.PassportNumber == "" || len([]rune(user.PassportNumber)) > 6 {
		return &entity.ResponseError{
			Message: "Invalid Passport Passport Number",
			Status:  http.StatusBadRequest,
		}
	}
	return nil
}

func (us *Service) UpdateUser(user *entity.Users) *entity.ResponseError {
	responseErr := validateUser(user)
	if responseErr != nil {
		return responseErr
	}
	return us.UserRepository.UpdateUser(user)
}

func (us *Service) DeleteUser(userId string) *entity.ResponseError {
	if userId == "" {
		return &entity.ResponseError{
			Message: "Invalid result ID",
			Status:  http.StatusBadRequest,
		}
	}
	err := BeginTransaction(us.UserRepository)
	if err != nil {
		return &entity.ResponseError{
			Message: "Failed to start transaction",
			Status:  http.StatusBadRequest,
		}
	}
	//Добавить роллбек
	responseErr := us.UserRepository.DeleteUser(userId)
	if responseErr != nil {
		return responseErr
	}
	CommitTransaction(us.UserRepository)
	return nil
}

func (us *Service) GetUser(userId string) (*entity.Users, *entity.ResponseError) {
	user, responseErr := us.UserRepository.GetUser(userId)
	if responseErr != nil {
		return nil, responseErr
	}
	return user, nil
}

// TODO тут фильтрацию сделать по-человечески
func (us *Service) GetUsersBach(limit int, offset int, name, surname, patronimyc, address, passportSerialNumber, passportNumber string) ([]*entity.Users, *entity.ResponseError) {
	if name != "" {
		users, responseErr := us.UserRepository.GetUserByParams(limit, offset, name)
		if responseErr != nil {
			return nil, &entity.ResponseError{
				Message: "Invalid Passport Passport Number",
				Status:  http.StatusNotFound,
			}
		}
		return users, nil
	}
	if surname != "" {
		users, responseErr := us.UserRepository.GetUserByParams(limit, offset, surname)
		if responseErr != nil {
			return nil, responseErr
		}
		return users, nil
	}
	if patronimyc != "" {
		users, responseErr := us.UserRepository.GetUserByParams(limit, offset, patronimyc)
		if responseErr != nil {
			return nil, responseErr
		}
		return users, nil
	}
	if address != "" {
		users, responseErr := us.UserRepository.GetUserByParams(limit, offset, address)
		if responseErr != nil {
			return nil, responseErr
		}
		return users, nil
	}
	if passportSerialNumber != "" {
		users, responseErr := us.UserRepository.GetUserByParams(limit, offset, passportSerialNumber)
		if responseErr != nil {
			return nil, responseErr
		}
		return users, nil
	}
	if passportNumber != "" {
		users, responseErr := us.UserRepository.GetUserByParams(limit, offset, passportNumber)
		if responseErr != nil {
			return nil, responseErr
		}
		return users, nil
	}
	users, responseErr := us.UserRepository.GetAllUsers(limit, offset)
	if responseErr != nil {
		return nil, responseErr
	}
	return users, nil
}

func (us *Service) Info(passportSerial string, passportNumber string) (*entity.Users, *entity.ResponseError) {
	user, responseErr := us.UserRepository.Info(passportSerial, passportNumber)
	if responseErr != nil {
		return nil, responseErr
	}
	return user, nil
}
