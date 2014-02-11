package kiicloud

// OauthTokenResponse represents response JSON of "/oauth2/token"
type OauthTokenResponse struct {
	Id          string `json:"id"`
	AccessToken string `json:"access_token"`
	ExpiresIn   int64  `json:"expires_in"`
	TokenType   string `json:"token_type"`
}
