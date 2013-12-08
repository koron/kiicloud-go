package kii

type Token struct {
	Id          string
	AccessToken string
	ExpiresIn   int64
	TokenType   string
}

type ErrorResponse struct {
	StatusCode int
	Code       string `json:"errorCode"`
	Message    string `json:"message"`
}

func (er *ErrorResponse) Error() string {
	return er.Code + ":" + string(er.StatusCode) + ":" + er.Message
}

type oauthTokenRequest struct {
	Username     string `json:"username,omitempty"`
	Password     string `json:"password,omitempty"`
	ExpiresAt    int64  `json:"expiresAt,omitempty"`
	ClientId     string `json:"client_id,omitempty"`
	ClientSecret string `json:"client_secret,omitempty"`
}

type oauthTokenResponse struct {
	Id          string `json:"id"`
	AccessToken string `json:"access_token"`
	ExpiresIn   int64  `json:"expires_in"`
	TokenType   string `json:"token_type"`
}
