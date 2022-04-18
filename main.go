package main

import (
	"RU-Smart-Workspace/ru-smart-api/databases"
	"RU-Smart-Workspace/ru-smart-api/environments"

	"RU-Smart-Workspace/ru-smart-api/handlers"
	"RU-Smart-Workspace/ru-smart-api/middlewares"
	student_repositories "RU-Smart-Workspace/ru-smart-api/repositories/student"
	student_services "RU-Smart-Workspace/ru-smart-api/services/student"

	"github.com/gin-gonic/gin"
	_ "github.com/godror/godror"
	"github.com/spf13/viper"
)

func init() {
	environments.TimeZoneInit()
	environments.EnvironmentInit()
}

func main() {




	oracle_db, err := databases.NewDatabases().OracleConnection()
	if err != nil {
		panic(err)
	}
	defer oracle_db.Close()

	redis_cache := databases.NewDatabases().RedisConnection()
	defer redis_cache.Close()

	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()
	router.Use(middlewares.NewCorsAccessControl().CorsAccessControl())

	studentRepo := student_repositories.NewStudentRepo(oracle_db)
	studentService := student_services.NewStudentServices(studentRepo, redis_cache)
	studentHandler := handlers.NewStudentHandlers(studentService)

	googleAuth := router.Group("/google")
	{
		googleAuth.POST("/authorization", middlewares.GoogleAuth, studentHandler.Authentication)
	}

	student := router.Group("/student")
	{
		student.POST("/refresh-authentication", studentHandler.RefreshAuthentication)
		student.GET("/profile/:std_code", middlewares.Authorization(redis_cache), studentHandler.GetStudentProfile)
	}

	PORT := viper.GetString("ruSmart.port")
	router.Run(PORT)
}
