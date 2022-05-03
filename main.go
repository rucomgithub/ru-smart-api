package main

import (
	"RU-Smart-Workspace/ru-smart-api/databases"
	"RU-Smart-Workspace/ru-smart-api/environments"

	student_handlers "RU-Smart-Workspace/ru-smart-api/handlers/student"
	"RU-Smart-Workspace/ru-smart-api/middlewares"
	student_repositories "RU-Smart-Workspace/ru-smart-api/repositories/student"
	student_services "RU-Smart-Workspace/ru-smart-api/services/student"

	mr30_handlers "RU-Smart-Workspace/ru-smart-api/handlers/public/mr30"
	mr30_repositories "RU-Smart-Workspace/ru-smart-api/repositories/public/mr30"
	mr30_services "RU-Smart-Workspace/ru-smart-api/services/public/mr30"

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

	googleAuth := router.Group("/google")
	{
		studentRepo := student_repositories.NewStudentRepo(oracle_db)
		studentService := student_services.NewStudentServices(studentRepo, redis_cache)
		studentHandler := student_handlers.NewStudentHandlers(studentService)

		googleAuth.POST("/authorization", middlewares.GoogleAuth, studentHandler.Authentication)
	}

	student := router.Group("/student")
	{
		studentRepo := student_repositories.NewStudentRepo(oracle_db)
		studentService := student_services.NewStudentServices(studentRepo, redis_cache)
		studentHandler := student_handlers.NewStudentHandlers(studentService)

		student.POST("/refresh-authentication", studentHandler.RefreshAuthentication)
		student.GET("/profile/:std_code", middlewares.Authorization(redis_cache), studentHandler.GetStudentProfile)
	}


	mr30 := router.Group("/mr30")
	{
		
		mr30Repo := mr30_repositories.NewMr30Repo(oracle_db)
		mr30Service := mr30_services.NewMr30Services(mr30Repo, redis_cache)
		mr30Handler := mr30_handlers.NewMr30Handlers(mr30Service)
		mr30.GET("/data", mr30Handler.GetMr30)
	}

	PORT := viper.GetString("ruSmart.port")
	router.Run(PORT)
}
