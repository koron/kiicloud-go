package kiicloud

import (
	"net/http"
)

// Client packs info to access the app of Kii Cloud.
type Client struct {
	// EndPoint of the app.
	EndPoint string

	// AppId of the app.
	AppId string

	// AppKey of the app.
	AppKey string

	// Authorization of this client.
	Authorization string
}

// NewClient creates new Client instance.
func NewClient(endPoint, id, key string) (*Client, error) {
	return &Client{endPoint, id, key, ""}, nil
}

// pathToEndPoint convert path to URL.
func (c *Client) pathToEndPoint(path string) string {
	return c.EndPoint + path
}

// NewRequest create a http.Request for the app.
func (c *Client) NewRequest(method, path string, v interface{}, ctype string) (req *http.Request, err error) {
	// Build request JSON.
	body, err := writeJson(v)
	if err != nil {
		return nil, err
	}
	req, err = http.NewRequest(method, c.pathToEndPoint(path), body)
	if err != nil {
		return nil, err
	}
	req.Header.Add("X-Kii-AppID", c.AppId)
	req.Header.Add("X-Kii-AppKey", c.AppKey)
	if ctype != "" {
		req.Header.Add("Content-Type", ctype)
	}
	if c.Authorization != "" {
		req.Header.Add("Authorization", c.Authorization)
	}
	return
}

// appPath generates path for the app.
func (c *Client) appPath(path string) string {
	return "/apps/" + c.AppId + "/" + path
}

// RegisterUser register a new user to the app.
func (c *Client) RegisterUser(loginName, password string, attrs *map[string]interface{}) (bool, error) {
	// Create a request object.
	reqobj := newMap(attrs)
	reqobj["loginName"] = loginName
	reqobj["password"] = password

	// Create a HTTP request.
	req, err := c.NewRequest("POST", c.appPath("users"), reqobj,
		"application/vnd.kii.RegistrationRequest+json")
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
	case 201:
		return true, nil
	default:
		return false, ToError(resp)
	}
}
