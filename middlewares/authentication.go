package middlewares

import (
	"time"

	"github.com/dgrijalva/jwt-go/v4"
	"github.com/spf13/viper"
)

type (
	TokenResponse struct {
		AccessToken  string `json:"access_token_id"`
		RefreshToken string `json:"refresh_token_id"`
		IsAuth       bool   `json:"is_auth"`
	}

	ClaimsToken struct {
		Issuer         string `json:"issuer"`
		Subject        string `json:"subject"`
		Role           string `json:"role"`
		ExpiresToken   string `json:"expires_token"`
		AccessTokenID  string `json:"access_token_id"`
		RefreshTokenID string `json:"refresh_token_id"`
	}
)

//  time.Now().Add(time.Hour * 30).Unix()
func GenerateToken(stdCode string) (*TokenResponse, error) {
	
	generateToken := &TokenResponse{}
	expirationAccessToken := time.Now().AddDate(0, 0, 1).Unix()
	expirationRefreshToken := time.Now().AddDate(0, 1, 0).Unix()

	generateToken.IsAuth = true

	// ---------------------  Create Access Token  ----------------------------------------- //
	accessTokenClaims := jwt.MapClaims{}
	accessTokenClaims["issuer"] = "Ru-Smart"
	accessTokenClaims["subject"] = "Ru-Smart" + stdCode
	accessTokenClaims["role"] = ""
	accessTokenClaims["expires_token"] = expirationAccessToken
	accessTokenClaims["access_token_id"] = stdCode + "::access"
	accessTokenClaims["refresh_token_id"] = stdCode + "::refresh"

	accessTokenHeader := jwt.NewWithClaims(jwt.SigningMethodHS256, accessTokenClaims)
	NEW_ACCESS_TOKEN, err := accessTokenHeader.SignedString([]byte(viper.GetString("token.secretKey")))
	if err != nil {
		return nil, err
	}

	generateToken.AccessToken = NEW_ACCESS_TOKEN	

	// ---------------------  Create Refresh Token  ----------------------------------------- //
	refreshTokenClaims := jwt.MapClaims{}
	refreshTokenClaims["issuer"] = "Ru-Smart"
	refreshTokenClaims["subject"] = "Ru-Smart" + stdCode
	refreshTokenClaims["role"] = ""
	refreshTokenClaims["expires_token"] = expirationRefreshToken
	refreshTokenClaims["access_token_id"] = stdCode + "::access"
	refreshTokenClaims["refresh_token_id"] = stdCode + "::refresh"

	refreshTokenHeader := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshTokenClaims)
	NEW_REFRESH_TOKEN, err := refreshTokenHeader.SignedString([]byte(viper.GetString("token.secretKey")))
	if err != nil {
		return nil, err
	}

	generateToken.RefreshToken = NEW_REFRESH_TOKEN

	return generateToken, nil
}
