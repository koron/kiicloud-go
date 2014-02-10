package kiicloud

// Client packs info to access the app of Kii Cloud.
type Client struct {
	// EndPoint of the app.
	EndPoint string

	// AppId of the app.
	AppId    string

	// AppKey of the app.
	AppKey   string
}

// NewClient creates new Client instance.
func NewClient(endPoint, id, key string) (*Client, error) {
	return &Client{endPoint, id, key}, nil
}

// RegisterUser register a new user.
func (c *Client) RegisterUser(loginName, password string) (bool, error) {
	// TODO: implement me.
	return false, nil
}
