package middlewares

import (
	"strings"

	"github.com/gin-gonic/gin"
)

// ทำการแกะ header HTTP request
// Authorization: Bearer TOKEN
func getHeaderAuthorization(c *gin.Context) (token string, err error) {

	const BEARER_SCHEMA = "Bearer "
	AUTH_HEADER := c.GetHeader("Authorization")

	if len(AUTH_HEADER) == 0 {
		return "", err
	}

	if strings.HasPrefix(AUTH_HEADER, BEARER_SCHEMA) {
		token = AUTH_HEADER[len(BEARER_SCHEMA):]
		return token, nil
	} else {
		return "", err
	}

}
