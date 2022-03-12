package middlewares

type (
	TokenResponse struct {
		AccessToken     string `json:"access_token"`
		RefreshToken    string `json:"refresh_token"`
		IsAuth          bool   `json:"is_auth"`
		AccessTokenKey  string `json:"access_token_key"`
		RefreshTokenKey string `json:"refresh_token_key"`
	}

	ClaimsToken struct {
		Issuer          string `json:"issuer"`
		Subject         string `json:"subject"`
		Role            string `json:"role"`
		ExpiresToken    string `json:"expires_token"`
		AccessTokenKey  string `json:"access_token_key"`
		RefreshTokenKey string `json:"refresh_token_key"`
	}

	CacheStudent struct {
		StdCode string `json:"std_code"`
		Role    string `json:"std_role"`
	}
)
