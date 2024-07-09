package controllers

import (
	"TimeTracker/models"
	"TimeTracker/services"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
	"strconv"
)

type TaskController struct {
	taskService *services.TaskService
}

func NewTaskController(taskService *services.TaskService) *TaskController {
	return &TaskController{
		taskService: taskService,
	}
}

//func (rh TaskController) AddTask(ctx *gin.Context) {
//	body, err := io.ReadAll(ctx.Request.Body)
//	if err != nil {
//		log.Println("Error while reading create user request body", err)
//		ctx.AbortWithError(http.StatusInternalServerError, err)
//		return
//	}
//	var user models.Users
//	err = json.Unmarshal(body, &user)
//	if err != nil {
//		log.Println("Error while unmarshaling create user request body", err)
//		ctx.AbortWithError(http.StatusInternalServerError, err)
//		return
//	}
//	response, responseErr := rh.taskService.AddTask(&runner)
//	if responseErr != nil {
//		ctx.AbortWithStatusJSON(responseErr.Status, responseErr)
//		return
//	}
//	ctx.JSON(http.StatusOK, response)
//}

//func (rh TaskController) UpdateTask(ctx *gin.Context) {
//	body, err := io.ReadAll(ctx.Request.Body)
//	if err != nil {
//		log.Println("Error while reading update user request body", err)
//		ctx.AbortWithError(http.StatusInternalServerError, err)
//		return
//	}
//	var task models.Task
//	err = json.Unmarshal(body, &task)
//	if err != nil {
//		log.Println("Error while unmarshaling update user request body", err)
//		ctx.AbortWithError(http.StatusInternalServerError, err)
//		return
//	}
//	responseErr := rh.taskService.UpdateTask(&task)
//	if responseErr != nil {
//		ctx.AbortWithStatusJSON(responseErr.Status, responseErr)
//		return
//	}
//	ctx.Status(http.StatusNoContent)
//}

//func (rh TaskController) DeleteTask(ctx *gin.Context) {
//	taskId := ctx.Param("id")
//	responseErr := rh.taskService.DeleteTask(taskId)
//	if responseErr != nil {
//		ctx.AbortWithError(responseErr.Status, responseErr)
//		return
//	}
//	ctx.Status(http.StatusNoContent)
//}

//func (rh TaskController) Task(ctx *gin.Context) {
//	params := ctx.Request.URL.Query()
//	taskID := params.Get("id")
//	response, responseErr := rh.taskService.Task(taskID)
//	if responseErr != nil {
//		ctx.JSON(responseErr.Status, responseErr)
//		return
//	}
//	ctx.JSON(http.StatusOK, response)
//}
//
//func (rh TaskController) Tasks(ctx *gin.Context) {
//	response, responseErr := rh.taskService.Tasks()
//	if responseErr != nil {
//		ctx.JSON(responseErr.Status, responseErr)
//		return
//	}
//	ctx.JSON(http.StatusOK, response)
//}

//func (rh TaskController) Labors(ctx *gin.Context) {
//	params := ctx.Request.URL.Query()
//	userID := params.Get("user_id")
//	response, responseErr := rh.taskService.Info(userID)
//	if responseErr != nil {
//		ctx.JSON(responseErr.Status, responseErr)
//		return
//	}
//	ctx.JSON(http.StatusOK, response)
//}

// @Summary get Labor Costs by User ID
// @Schemes
// @Description get Labor Costs by User ID asd
// @Tags labor
// @Accept json
// @Produce json
// @Param        user_id   query      int  true  "User ID"
// @Success      200  {object}  models.Labor
// @Failure      400
// @Failure      404
// @Failure      500
// @Success 200
// @Router /get_labor_costs [get]
func (tc TaskController) LaborsCost(ctx *gin.Context) {
	params := ctx.Request.URL.Query()
	userId := params.Get("user_id")
	userIdInt, err := strconv.Atoi(userId)
	if err != nil {
		userIdInt = 0
	}
	response, responseErr := tc.taskService.GetLaborCostsByUserId(userIdInt)
	if responseErr != nil {
		ctx.JSON(responseErr.Status, responseErr)
		return
	}
	ctx.JSON(http.StatusOK, response)
}

// @Summary StartTask
// @Schemes
// @Description Start users task by username
// @Tags tasks
// @Accept json
// @Produce json
// @Param        request body models.StartTask true "start task json"
// @Success      204
// @Failure      400
// @Failure      404
// @Failure      500
// @Success 200
// @Router /start_task [put]
func (tc TaskController) StartTask(ctx *gin.Context) {
	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		log.Println("Error while reading update user request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	var task models.StartTask
	err = json.Unmarshal(body, &task)
	if err != nil {
		log.Println("Error while unmarshaling update user request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	responseErr := tc.taskService.StartTask(&task)
	if responseErr != nil {
		ctx.AbortWithStatusJSON(responseErr.Status, responseErr)
		return
	}
	ctx.Status(http.StatusNoContent)
}

// @Summary EndTask
// @Schemes
// @Description End users task by username
// @Tags tasks
// @Accept json
// @Produce json
// @Param        request body models.EndTask true "end task json"
// @Success      204
// @Failure      400
// @Failure      404
// @Failure      500
// @Success 200
// @Router /end_task [put]
func (tc TaskController) EndTask(ctx *gin.Context) {
	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		log.Println("Error while reading update user request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	var task models.EndTask
	err = json.Unmarshal(body, &task)
	if err != nil {
		log.Println("Error while unmarshaling update user request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	responseErr := tc.taskService.EndTask(&task)
	if responseErr != nil {
		ctx.AbortWithStatusJSON(responseErr.Status, responseErr)
		return
	}
	ctx.Status(http.StatusNoContent)
}
