package main

import (
	"RU-Smart-Workspace/ru-smart-api/databases"
	"RU-Smart-Workspace/ru-smart-api/environments"
	"RU-Smart-Workspace/ru-smart-api/handlers"
	"RU-Smart-Workspace/ru-smart-api/middlewares"
	"RU-Smart-Workspace/ru-smart-api/repositories"
	"RU-Smart-Workspace/ru-smart-api/services"

	"github.com/gin-gonic/gin"
	_ "github.com/godror/godror"
	"github.com/spf13/viper"
)

func init() {
	environments.TimeZoneInit()
	environments.EnvironmentInit()
}

func main() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.Use(middlewares.NewCorsAccessControl().CorsAccessControl())

	oracle_db, err := databases.NewDatabases().OracleConnection()
	if err != nil {
		panic(err)
	}
	defer oracle_db.Close()

	studentRepo := repositories.NewStudentRepo(oracle_db)
	studentService := services.NewStudentServices(studentRepo)
	studentHandler := handlers.NewStudentHandlers(studentService)

	googleAuth := router.Group("/google")
	{
		googleAuth.POST("/authorization", middlewares.GoogleAuth, studentHandler.Authentication)
	}


	// student := router.Group("/student")
	// {
	// 	student.POST("/authentication", newStudentHandler.Authentication)
	// 	student.POST("/profile", middlewares.authorization, newStudentHandler.Authentication)
	// 	student.POST("/news", newStudentHandler.Authentication)
	// }


	
	// router.GET("/Authentication",func (c *gin.Context) {
	// 	studentProfile, err :=  newStudentRepo.GetAuthentication("6299999991")
	// 	if err != nil {
	// 		log.Fatal(err)
	// 		c.Abort()
	// 	}

	// 	c.IndentedJSON(http.StatusOK, studentProfile)
	// 	c.Next()
	// })

    PORT := viper.GetString("ruSmart.port")
	router.Run(PORT)
}