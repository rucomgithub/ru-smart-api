package services

import "RU-Smart-Workspace/ru-smart-api/repositories"

type (

	studentServices struct {
		studentRepo repositories.StudentRepoInterface
	}

	AuthenPlayload struct {
		Std_code string `json:"std_code"`
	}

	TokenResponse struct {
		AccessToken  string `json:"access_token_id"`
		RefreshToken string `json:"refresh_token_id"`
		IsAuth       bool   `json:"is_auth"`
		Message      string `json:"message"`
		StatusCode   int    `json:"status_code"`
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

	StudentServicesInterface interface {
		Authentication(stdCode string) (*TokenResponse, error)
	}
)

func NewStudentServices(studentRepo repositories.StudentRepoInterface) StudentServicesInterface {
	return studentServices{studentRepo: studentRepo}
}

