package handlers

import (
	"RU-Smart-Workspace/ru-smart-api/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type studentHandlers struct {
	studentService services.StudentServicesInterface
}

func NewStudentHandlers(studentService services.StudentServicesInterface) studentHandlers {
	return studentHandlers{studentService: studentService}
}

func (h studentHandlers) Authentication(c *gin.Context) {

	var requestBody services.AuthenPlayload

	err := c.ShouldBindJSON(&requestBody)
	if err != nil {
		c.IndentedJSON(http.StatusUnauthorized, err)
		c.Abort()
	}

	token, err := h.studentService.Authentication(requestBody.Std_code)
	
	if err != nil {
		c.IndentedJSON(http.StatusUnauthorized, err)
		c.Abort()
	}

	c.IndentedJSON(http.StatusOK, token)

}

func (h studentHandlers) GetStudentProfile(c *gin.Context) {

	STD_CODE := c.Param("std_code")
	studentProfileResponse, err := h.studentService.GetStudentProfile(STD_CODE)
	if err != nil {
		c.IndentedJSON(http.StatusUnauthorized, err)
		c.Abort()
	}

	c.IndentedJSON(http.StatusOK, studentProfileResponse)

}
