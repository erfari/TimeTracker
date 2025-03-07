package http

import (
	"TimeTracker/internal/entity"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
	"strconv"
)

type UserService interface {
	AddUserApi(passport *entity.PassportDocument) (*entity.Users, *entity.ResponseError)
	AddUser(user *entity.Users) (*entity.Users, *entity.ResponseError)
	UpdateUser(user *entity.Users) *entity.ResponseError
	DeleteUser(userId string) *entity.ResponseError
	GetUser(userId string) (*entity.Users, *entity.ResponseError)
	GetUsersBach(limit int, offset int, name, surname, patronimyc, address, passportSerialNumber, passportNumber string) ([]*entity.Users, *entity.ResponseError)
	Info(passportSerial string, passportNumber string) (*entity.Users, *entity.ResponseError)
}

type UserController struct {
	userService UserService
}

func NewUserController(userService UserService) *UserController {
	return &UserController{
		userService: userService,
	}
}

// @Summary Add new user
// @Schemes
// @Description get Labor Costs by User ID asd
// @Tags users
// @Accept json
// @Produce json
// @Param        request body types.Users true "add new user json"
// @Success      200  {object}  types.Users
// @Failure      400
// @Failure      404
// @Failure      500
// @Success 200
// @Router /add_user_api [post]
func (uc UserController) AddUserApi(ctx *gin.Context) {
	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		log.Println("Error while reading create user request body", err)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}
	var userDocument entity.PassportDocument
	err = json.Unmarshal(body, &userDocument)
	if err != nil {
		log.Println("Error while unmarshaling create user request body", err)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}
	response, responseErr := uc.userService.AddUserApi(&userDocument)
	if responseErr != nil {
		ctx.AbortWithStatusJSON(responseErr.Status, responseErr)
		return
	}
	ctx.JSON(http.StatusOK, response)
}

// @Summary Add new user by passport
// @Schemes
// @Description add new user by passport if exist
// @Tags users
// @Accept json
// @Produce json
// @Param        request body types.PassportDocument true "add new user by passport json"
// @Success      200  {object}  types.Users
// @Failure      400
// @Failure      404
// @Failure      500
// @Success 200
// @Router /add_user [post]
func (uc UserController) AddUser(ctx *gin.Context) {
	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		log.Println("Error while reading create user request body", err)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	var user entity.Users
	err = json.Unmarshal(body, &user)
	if err != nil {
		log.Println("Error while unmarshaling create user request body", err)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	response, responseErr := uc.userService.AddUser(&user)
	if responseErr != nil {
		ctx.AbortWithStatusJSON(responseErr.Status, responseErr)
		return
	}
	ctx.JSON(http.StatusOK, response)
}

// @Summary Update users
// @Schemes
// @Description update user info
// @Tags users
// @Accept json
// @Produce json
// @Param        request body types.Users true "update user json"
// @Success      200  {object}  types.Users
// @Failure      400
// @Failure      404
// @Failure      500
// @Success 200
// @Router /update_user [put]
func (uc UserController) UpdateUser(ctx *gin.Context) {
	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		log.Println("Error while reading update user request body", err)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	var user entity.Users
	err = json.Unmarshal(body, &user)
	if err != nil {
		log.Println("Error while unmarshaling update user request body", err)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	responseErr := uc.userService.UpdateUser(&user)
	if responseErr != nil {
		ctx.AbortWithStatusJSON(responseErr.Status, responseErr)
		return
	}
	ctx.Status(http.StatusOK)
}

// @Summary Delete users
// @Schemes
// @Description delete user
// @Tags users
// @Accept json
// @Produce json
// @Param        user_id   path      int  true  "User ID"
// @Success      204  {object}  types.Users
// @Failure      400
// @Failure      404
// @Failure      500
// @Success 200
// @Router /delete_user/{user_id} [delete]
func (uc UserController) DeleteUser(ctx *gin.Context) {
	userId := ctx.Param("id")
	responseErr := uc.userService.DeleteUser(userId)
	if responseErr != nil {
		ctx.AbortWithStatusJSON(responseErr.Status, responseErr)
		return
	}
	ctx.Status(http.StatusNoContent)
}

// @Summary Get user
// @Schemes
// @Description Get user by user id
// @Tags users
// @Accept json
// @Produce json
// @Param        user_id   query      int  true  "User ID" optional
// @Success      200  {object}  types.Users
// @Failure      400
// @Failure      404
// @Failure      500
// @Success 200
// @Router /get_user [get]
func (uc UserController) GetUser(ctx *gin.Context) {
	userId := ctx.Query("id")
	response, responseErr := uc.userService.GetUser(userId)
	if responseErr != nil {
		ctx.AbortWithStatusJSON(responseErr.Status, responseErr)
		return
	}
	ctx.JSON(http.StatusOK, response)
}

// @Summary Get list users
// @Schemes
// @Description Get list users with filters by fields and pagination by limit and offset
// @Tags users
// @Accept json
// @Produce json
// @Param        limit   query      int  true  "Limit"
// @Param        offset   query      int  true  "Offset"
// @Param        name   query      string  false  "Name" *
// @Param        surname   query      string  false  "Surname" *
// @Param        patronimyc   query      string  false  "Patronimyc" *
// @Param        address   query      string  false  "Address" *
// @Param        passportSerialNumber   query      string  false  "PassportSerialNumber" *
// @Param        passportNumber   query      string  false  "PassportNumber" *
// @Success      200  {array}  types.Users
// @Failure      400
// @Failure      404
// @Failure      500
// @Success 200
// @Router /get_all_users [get]
func (uc UserController) GetUsers(ctx *gin.Context) {
	limitStr := ctx.Query("limit")
	limit, err := strconv.ParseUint(limitStr, 10, 64)
	if err != nil {
		limit = 5
	}
	offsetStr := ctx.Query("offset")
	offset, err := strconv.ParseUint(offsetStr, 10, 64)
	if err != nil {
		offset = 0
	}
	name := ctx.Query("name")
	surname := ctx.Query("surname")
	patronimyc := ctx.Query("patronimyc")
	address := ctx.Query("address")
	passportSerialNumber := ctx.Query("passportSerialNumber")
	passportNumber := ctx.Query("passportNumber")
	response, responseErr := uc.userService.GetUsersBach(int(limit), int(offset), name, surname, patronimyc, address, passportSerialNumber, passportNumber)
	if responseErr != nil {
		ctx.AbortWithStatusJSON(responseErr.Status, responseErr)
		return
	}
	ctx.JSON(http.StatusOK, response)
}

func (uc UserController) Info(ctx *gin.Context) {
	params := ctx.Request.URL.Query()
	passportSerial := params.Get("passportSerial")
	passportNumber := params.Get("passportNumber")
	response, responseErr := uc.userService.Info(passportSerial, passportNumber)
	if responseErr != nil {
		ctx.AbortWithStatusJSON(responseErr.Status, responseErr)
		return
	}
	ctx.JSON(http.StatusOK, response)
}
