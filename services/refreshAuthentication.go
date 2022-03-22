package services

import (
	"RU-Smart-Workspace/ru-smart-api/middlewares"
	"log"
)

func (s studentServices) RefreshAuthentication(refreshToken, stdCode string) (*TokenResponse, error) {

	log.Println("----- RefreshAuthentication -----")
	log.Println("--------- RefreshToken -----------")
	log.Println("RefreshToken :: ", refreshToken)
	log.Println("--------- Student code ----------")
	log.Println("Student code :: ", stdCode)
	log.Println("---------------------------------")

	// ส่ง Token ไปตรวจสอบว่าได้รับสิทธิ์เข้าใช้งานหรือไม่
	isToken, err := middlewares.VerifyToken("refreshToken", refreshToken, s.redis_cache)
	if err != nil {
		return nil, err
	}
	log.Println(" \t\t     <========== isToken ==========> \n ", isToken, " \n ")
	log.Println(" ------------------------------------------------------------------- \n ")

	prepareToken, err := s.studentRepo.Authentication(stdCode)
	if err != nil || prepareToken.STATUS != 1 {
		return nil, err
	}

	generateToken, err := middlewares.GenerateToken(prepareToken.STD_CODE, s.redis_cache)
	if err != nil {
		return nil, err
	}

	studentTokenResponse := TokenResponse{
		AccessToken:  generateToken.AccessToken,
		RefreshToken: generateToken.RefreshToken,
		IsAuth:       generateToken.IsAuth,
		Message:      "Generate Token success...",
		StatusCode:   200,
	}
	return &studentTokenResponse, nil

}
