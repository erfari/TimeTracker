package app

import (
	"TimeTracker/docs"
	controllers "TimeTracker/internal/controllers"
	repository "TimeTracker/internal/repository"
	services "TimeTracker/internal/services"
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
	userController *controllers.UserController
}

func InitHttpServer(config *viper.Viper, dbHandler *sql.DB) (HttpServer, error) {
	//user
	userRepository := repository.NewUserRepository(dbHandler)
	userService := services.NewUserService(userRepository)
	userController := controllers.NewUserController(userService)

	//task
	taskRepository := repository.NewTaskRepository(dbHandler)
	taskService := services.NewTaskService(taskRepository)
	taskController := controllers.NewTaskController(taskService)

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
