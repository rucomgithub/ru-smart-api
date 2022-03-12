package middlewares

import (
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go/v4"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
)

func Authorization(redis_cache *redis.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := getHeaderAuthorization(c)
		if err != nil {
			c.IndentedJSON(http.StatusUnauthorized, gin.H{"accessToken": "", "isAuth": false, "message": "authorization key in header not found"})
			c.Abort()
			return
		}

		// ส่ง Token ไปตรวจสอบว่าได้รับสิทธิ์เข้าใช้งานหรือไม่
		isToken, err := verifyAccessToken(token, redis_cache)
		if err != nil {
			fmt.Println(err)
			c.IndentedJSON(http.StatusUnauthorized, err)
			c.Abort()
			return
		}

		if isToken {
			c.Next()
		}
	}

}

func verifyAccessToken(accessToken string, redis_cache *redis.Client) (bool, error) {

	claims, err := getClaims(accessToken)
	if err != nil {
		return false, err
	}

	_, err = redis_cache.Get(ctx, claims.AccessTokenKey).Result()
	if err != nil {
		return false, err
	}

	return true, nil
}

func getClaims(encodedToken string) (*ClaimsToken, error) {

	parseToken, err := jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(viper.GetString("token.secretKey")), nil
	})
	if err != nil {
		return nil, err
	}

	claimsToken := &ClaimsToken{}
	parseClaims := parseToken.Claims.(jwt.MapClaims)

	if parseClaims["issuer"] != nil {
		claimsToken.Issuer = parseClaims["issuer"].(string)
	}

	if parseClaims["subject"] != nil {
		claimsToken.Subject = parseClaims["subject"].(string)
	}

	if parseClaims["role"] != "" {
		claimsToken.Role = parseClaims["role"].(string)
	} else {
		claimsToken.Role = ""
	}
	
	if parseClaims["access_token_key"] != nil {
		claimsToken.AccessTokenKey = parseClaims["access_token_key"].(string)
	}
	
	if parseClaims["refresh_token_key"] != nil {
		claimsToken.RefreshTokenKey = parseClaims["refresh_token_key"].(string)	}
	
	if parseClaims["expires_token"] != nil {
		claimsToken.ExpiresToken = fmt.Sprintf("%v", parseClaims["expires_token"])
	}

	return claimsToken, nil
}
