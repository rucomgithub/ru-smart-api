package middlewares

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

func RefreshAuthentication(redis_cache *redis.Client) gin.HandlerFunc {

	return func(c *gin.Context) {

		refreshToken, err := getHeaderAuthorization(c)
		if err != nil {
			c.IndentedJSON(http.StatusUnauthorized, gin.H{"message": "กรุณา Sign-in เพื่อเข้าสู่ระบบใหม่"})
			c.Abort()
			return
		}

		log.Println(" ---------------------- Refresh Authentication ---------------------- \n ")
		log.Println(" \t\t     <========== RefreshToken ==========> \n ", refreshToken+"\n")
		log.Println(" ------------------------------------------------------------------- \n ")

		// ส่ง Token ไปตรวจสอบว่าได้รับสิทธิ์เข้าใช้งานหรือไม่
		isToken, err := verifyToken("refreshToken", refreshToken, redis_cache)
		if err != nil {
			c.IndentedJSON(http.StatusUnauthorized, err)
			c.Abort()
			return
		}

		if isToken {
			c.Next()
		}
	}
}
