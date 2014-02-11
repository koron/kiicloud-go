package kiicloud

import (
	"fmt"
	"net/http"
)

// AdminClient packs info to access the app as administrator.
type AdminClient struct {
	// All inherit from Client.
	Client

	// ClientId of the app.
	ClientId string

	// ClientSecret of the app.
	ClientSecret string
}

// NewAdminClient create a new AdminClient.
func NewAdminClient(entryPoint, appId, appKey, clientId, clientSecret string) (*AdminClient, error) {
	return &AdminClient{
		Client{entryPoint, appId, appKey, ""},
		clientId, clientSecret,
	}, nil
}

// Authorize assure admin authorization.
func (c *AdminClient) Authorize() (string, error) {
	if c.Authorization != "" {
		return c.Authorization, nil
	}

	// Create a request object.
	reqobj := newMap(nil)
	reqobj["client_id"] = c.ClientId
	reqobj["client_secret"] = c.ClientSecret

	// Create a HTTP request.
	req, err := c.NewRequest("POST", "/oauth2/token", reqobj,
		"application/vnd.kii.OauthTokenRequest+json")
	if err != nil {
		return "", err
	}

	// Do the request and read its response's body.
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	switch resp.StatusCode {
	case 200:
		var respobj OauthTokenResponse
		err = parseJson(resp.Body, &respobj)
		if err != nil {
			return "", err
		}
		c.Authorization = fmt.Sprintf("%s %s", respobj.TokenType,
			respobj.AccessToken)
		return c.Authorization, nil
	default:
		return "", ToError(resp)
	}
}

// UnregisterUser unregisters a user by login name.
func (c *AdminClient) UnregisterUser(loginName string) (bool, error) {
	_, err := c.Authorize()
	if err != nil {
		return false, err
	}

	// Create a HTTP request.
	req, err := c.NewRequest("DELETE",
		c.appPath("users/LOGIN_NAME:"+loginName), nil, "")
	if err != nil {
		return false, err
	}

	// Do the request and read its response's body.
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	switch resp.StatusCode {
	case 204:
		return true, nil
	default:
		return false, ToError(resp)
	}
}
