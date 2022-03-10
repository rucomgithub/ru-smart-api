package middlewares

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"google.golang.org/api/oauth2/v1"
	"google.golang.org/api/option"
)

var ctx = context.Background()

type (
	Payload struct {
		Std_code string `json:"std_code"`
	}

	TokenResponse struct {
		AccessToken     string `json:"access_token_id"`
		RefreshToken    string `json:"refresh_token_id"`
		IsAuth      bool   `json:"is_auth"`
		Message     string `json:"message"`
		StatusCode  int    `json:"status_code"`
	}

	// claims คือข้อมูลที่อยู่ในส่วน Payload ของ Token
	// -iss (issuer) : เว็บหรือบริษัทเจ้าของ token
	// -sub (subject) : subject ของ token
	// -aud (audience) : ผู้รับ token
	// -exp (expiration time) : เวลาหมดอายุของ token
	// -nbf (not before) : เป็นเวลาที่บอกว่า token จะเริ่มใช้งานได้เมื่อไหร่
	// -iat (issued at) : ใช้เก็บเวลาที่ token นี้เกิดปัญหา
	// -jti (JWT id) : เอาไว้เก็บไอดีของ JWT แต่ละตัวนะครับ
	// -name (Full name) : เอาไว้เก็บชื่อ
	ClaimsToken struct {
		Issuer              string `json:"issuer"`
		Subject             string `json:"subject"`
		Role                string `json:"role"`
		ExpiresAccessToken  string `json:"expires_access_token"`
		ExpiresRefreshToken string `json:"expiration_refresh_token"`
	}

	generateToken struct {
		oracle_db *sqlx.DB
	}

	GoogleAuthInterface interface {
		Authentication(stdCode string) (*TokenResponse, error)
	}

)

func NewGenerateToken(oracle_db *sqlx.DB) GoogleAuthInterface {
	return generateToken{oracle_db: oracle_db}
}

func GoogleAuth(c *gin.Context) {

	const BEARER_SCHEMA = "Bearer "
	AUTH_HEADER := c.GetHeader("Authorization")
	ID_TOKEN := AUTH_HEADER[len(BEARER_SCHEMA):]

	if len(AUTH_HEADER) == 0 {
		c.IndentedJSON(http.StatusUnauthorized, "authorization key in header not found")
		c.Abort()
		return
	}

	var student Payload
	err := c.ShouldBindJSON(&student)
	if err != nil {
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"accessToken": "", "isAuth": false, "message": "Google is not authorized"})
		c.Abort()
		return
	}
	_, err = verifyGoogleAuth(ID_TOKEN)
	if err != nil {
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"accessToken": "", "isAuth": false, "message": "Google is not authorized"})
		c.Abort()
		return
	}

	log.Println("generate token here...")
	/// generate token
	// generateToken()

	c.IndentedJSON(http.StatusOK, gin.H{"accessToken": "1234_TOKEN", "isAuth": true, "message": "Google is authorized"})
	c.Next()

}

func verifyGoogleAuth(id_token string) (*oauth2.Tokeninfo, error) {

	// ตั้งเวลาการรอคอย หรือกำหนดการหยุดเชื่อมต่อกรณีที่ไม่ได้ Respose จาก Google
	timeout := time.Duration(5 * time.Second)
	httpClient := &http.Client{Timeout: timeout}

	oauth2Service, err := oauth2.NewService(ctx, option.WithHTTPClient(httpClient))
	if err != nil {
		return nil, err
	}

	tokenInfoCall := oauth2Service.Tokeninfo()
	tokenInfoCall.IdToken(id_token)
	tokenInfo, err := tokenInfoCall.Do()
	if err != nil {
		return nil, err
	}
	return tokenInfo, nil
}

func (r generateToken) Authentication(stdCode string) (*TokenResponse, error) {

	// newStudentRepo := repositories.NewStudentProfileRepo(r.oracle_db)

	log.Println("generateToken")

	return nil,nil

}
