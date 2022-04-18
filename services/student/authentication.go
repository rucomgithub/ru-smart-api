package student_services

import (
	"RU-Smart-Workspace/ru-smart-api/middlewares"
)

func (s studentServices) Authentication(stdCode string) (*TokenResponse, error) {

	prepareToken, err := s.studentRepo.Authentication(stdCode)
	if err != nil || prepareToken.STATUS != 1 {  
		return nil, err
	}

	generateToken, err := middlewares.GenerateToken(prepareToken.STD_CODE, s.redis_cache)
	if err != nil {
		return nil, err
	}

	studentTokenResponse := TokenResponse{
		AccessToken  : generateToken.AccessToken,
		RefreshToken : generateToken.RefreshToken,
		IsAuth       : generateToken.IsAuth,
		Message      : "Generate Token success...",
		StatusCode   : 200,
	}
	return &studentTokenResponse, nil
}