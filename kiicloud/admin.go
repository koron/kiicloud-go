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
		Client{entryPoint, appId, appKey, "", false},
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

	err := c.Send("/oauth2/token", "POST", reqobj,
		"application/vnd.kii.OauthTokenRequest+json",
		func(resp *http.Response) error {
			switch resp.StatusCode {
			case 200:
				var respobj OauthTokenResponse
				err := parseJson(resp.Body, &respobj)
				if err != nil {
					return err
				}
				c.Authorization = fmt.Sprintf("%s %s", respobj.TokenType,
					respobj.AccessToken)
				return nil
			default:
				return ToError(resp)
			}
		})
	if err != nil {
		return "", err
	}

	return c.Authorization, nil
}

// UnregisterUser unregisters a user by login name.
func (c *AdminClient) UnregisterUser(loginName string) (bool, error) {
	_, err := c.Authorize()
	if err != nil {
		return false, err
	}

	err = c.Send(c.appPath("users/LOGIN_NAME:"+loginName), "DELETE", nil, "",
		func(resp *http.Response) error {
			switch resp.StatusCode {
			case 204:
				return nil
			default:
				return ToError(resp)
			}
		})
	if err != nil {
		return false, err
	}

	return true, err
}
