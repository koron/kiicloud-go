package kiicloud

import (
	"fmt"
	"net/http"
)

// UserClient packs info to access the app as a user.
type UserClient struct {
	// All inherit from Client.
	Client
}

// NewUserClient create a new user client.
func NewUserClient(entryPoint, appId, appKey, loginName, password string) (*UserClient, error) {
	c := &UserClient{Client {entryPoint, appId, appKey, "", false}}

	var err error
	c.Authorization, err = c.authorize(loginName, password)
	if err != nil {
		return nil, err
	}

	return c, nil
}

// authorize assure user login.
func (c *UserClient) authorize(loginName, password string) (string, error) {
	// Create a request object.
	reqobj := newMap(nil)
	reqobj["username"] = loginName
	reqobj["password"] = password

	authorization := ""
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
				authorization = fmt.Sprintf("%s %s", respobj.TokenType,
					respobj.AccessToken)
				return nil
			default:
				return ToError(resp)
			}
		})
	if err != nil {
		return "", err
	}

	return authorization, nil
}

func (c *UserClient) Bucket(name string) (*Bucket, error) {
	// TODO: compose correct path root.
	pathRoot := "/"
	return &Bucket{c, &c.Client, pathRoot}, nil
}
