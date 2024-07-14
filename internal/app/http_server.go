package app

import (
	"TimeTracker/docs"
	"TimeTracker/internal/controller/http"
	"TimeTracker/internal/usecase/tasks"
	services "TimeTracker/internal/usecase/user"
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
)

type HttpServer struct {
	config         *viper.Viper
	router         *gin.Engine
	userController *http.UserController
}

func InitHttpServer(config *viper.Viper, dbHandler *sql.DB) (HttpServer, error) {
	//user
	userRepository := services.NewUserRepository(dbHandler)
	userService := services.NewUserService(userRepository)
	userController := http.NewUserController(userService)

	//task
	taskRepository := tasks.NewTaskRepository(dbHandler)
	taskService := tasks.NewTaskService(taskRepository)
	taskController := http.NewTaskController(taskService)

	router := gin.Default()
	// user
	router.GET("/info", userController.Info)
	router.POST("/add_user_api", userController.AddUser)
	router.POST("/add_user", userController.AddUserApi)
	router.PUT("/update_user", userController.UpdateUser)
	router.DELETE("/delete_user/:id", userController.DeleteUser)
	router.GET("/get_user", userController.GetUser)
	router.GET("/get_all_users", userController.GetUsers)

	//tasks
	router.PUT("/start_task", taskController.StartTask)
	router.PUT("/end_task", taskController.EndTask)
	router.POST("/add_task", taskController.AddTask)
	router.POST("/get_task", taskController.GetTask)
	router.POST("/update_task", taskController.UpdateTask)
	router.GET("/get_labor_costs", taskController.LaborsCost)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	docs.SwaggerInfo.Description = "This is a sample time tracker server."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "time tracker.swagger.io"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	return HttpServer{
		config:         config,
		router:         router,
		userController: userController,
	}, nil
}

func (hs HttpServer) Start() {
	err := hs.router.Run(hs.config.GetString(
		"SERVER_PORT"))
	if err != nil {
		log.Fatalf("Error while starting HTTP server: %v", err)
	}
}
