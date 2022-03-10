package handlers

import (
	"RU-Smart-Workspace/ru-smart-api/repositories"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type studentHandlers struct {
	studentRepo repositories.StudentRepoInterface
}

func NewStudentHandlers(studentRepo repositories.StudentRepoInterface) studentHandlers {
	return studentHandlers{studentRepo: studentRepo}
}

func (h studentHandlers) Authentication(c *gin.Context) {

	var requestBody repositories.AuthenPlayload
	err := c.ShouldBindJSON(&requestBody)
	if err != nil {
		log.Println("ตายที่ Authentication",err.Error())
		c.IndentedJSON(http.StatusUnauthorized, "ตายที่ Authentication")
		c.Abort()
	}

	token, err := h.studentRepo.GetAuthentication(requestBody.Std_code)
	if err != nil {
		log.Println("ตายที่ Authentication",err.Error())
		c.IndentedJSON(http.StatusUnauthorized, "ตายที่ generate token Authentication")
		c.Abort()
	}

	log.Println("stdCode ==> ", token)
}
