package http

import (
	"TimeTracker/internal/entity"
	"TimeTracker/internal/usecase/tasks"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
	"strconv"
)

type TaskController struct {
	taskService *tasks.TaskService
}

func NewTaskController(taskService *tasks.TaskService) *TaskController {
	return &TaskController{
		taskService: taskService,
	}
}

// @Summary get Labor Costs by User ID
// @Schemes
// @Description get Labor Costs by User ID asd
// @Tags labor
// @Accept json
// @Produce json
// @Param        user_id   query      int  true  "User ID"
// @Success      200  {object}  types.Labor
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
		ctx.AbortWithStatusJSON(responseErr.Status, responseErr)
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
// @Param        request body types.StartTask true "start task json"
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
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}
	var task entity.StartTask
	err = json.Unmarshal(body, &task)
	if err != nil {
		log.Println("Error while unmarshaling update user request body", err)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err)
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
// @Param        request body types.EndTask true "end task json"
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
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}
	var task entity.EndTask
	err = json.Unmarshal(body, &task)
	if err != nil {
		log.Println("Error while unmarshaling update user request body", err)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}
	responseErr := tc.taskService.EndTask(&task)
	if responseErr != nil {
		ctx.AbortWithStatusJSON(responseErr.Status, responseErr)
		return
	}
	ctx.Status(http.StatusNoContent)
}

func (tc TaskController) AddTask(ctx *gin.Context) {
	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		log.Println("Error while reading create task request body", err)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}
	var task entity.Task
	err = json.Unmarshal(body, &task)
	if err != nil {
		log.Println("Error while unmarshaling create user request body", err)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}
	response, responseErr := tc.taskService.AddTask(&task)
	if responseErr != nil {
		ctx.AbortWithStatusJSON(responseErr.Status, responseErr)
		return
	}
	ctx.JSON(http.StatusOK, response)
}

func (tc TaskController) UpdateTask(ctx *gin.Context) {
	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		log.Println("Error while reading update user request body", err)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}
	var task entity.Task
	err = json.Unmarshal(body, &task)
	if err != nil {
		log.Println("Error while unmarshaling update user request body", err)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}
	responseErr := tc.taskService.UpdateTask(&task)
	if responseErr != nil {
		ctx.AbortWithStatusJSON(responseErr.Status, responseErr)
		return
	}
	ctx.Status(http.StatusNoContent)
}

func (tc TaskController) DeleteTask(ctx *gin.Context) {
	taskID := ctx.Param("id")
	newTaskID, err := strconv.Atoi(taskID)
	if err != nil {
		newTaskID = 0
	}
	responseErr := tc.taskService.DeleteTask(newTaskID)
	if responseErr != nil {
		ctx.AbortWithStatusJSON(responseErr.Status, responseErr)
		return
	}
	ctx.Status(http.StatusNoContent)
}

func (tc TaskController) GetTask(ctx *gin.Context) {
	taskID := ctx.Param("id")
	newTaskID, err := strconv.Atoi(taskID)
	if err != nil {
		newTaskID = 0
	}
	response, responseErr := tc.taskService.GetTask(newTaskID)
	if responseErr != nil {
		ctx.JSON(responseErr.Status, responseErr)
		return
	}
	ctx.JSON(http.StatusOK, response)
}
