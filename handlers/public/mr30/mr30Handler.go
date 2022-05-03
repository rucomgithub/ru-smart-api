package mr30_handlers

import (
	mr30_services "RU-Smart-Workspace/ru-smart-api/services/public/mr30"
	"net/http"

	"github.com/gin-gonic/gin"
)

type mr30Handlers struct {
	mr30Services mr30_services.Mr30ServiceInterface
}

func NewMr30Handlers(mr30Services mr30_services.Mr30ServiceInterface) mr30Handlers {
	return mr30Handlers{mr30Services: mr30Services}
}

func (h *mr30Handlers) GetMr30(c *gin.Context) {

	var requestBody mr30_services.Mr30Request

	err := c.ShouldBindJSON(&requestBody)
	if err != nil {
		c.IndentedJSON(http.StatusUnprocessableEntity, err)
		c.Abort()
		return
	}

	mr30Response, err := h.mr30Services.GetMr30(requestBody.Course_year, requestBody.Course_semester)
	if err != nil {
		c.IndentedJSON(http.StatusUnprocessableEntity, err)
		c.Abort()
		return
	}

	c.IndentedJSON(http.StatusOK, mr30Response)

}