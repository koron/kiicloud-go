package kiicloud

// AdminClient packs info to access the app as administrator.
type AdminClient struct {
	// All inherit from Client.
	Client

	// ClientId of the app.
	ClientId string

	// ClientSecret of the app.
	ClientSecret string
}

func NewAdminClient(entryPoint, appId, appKey, clientId, clientSecret string) (*AdminClient, error) {
	return &AdminClient{
		Client{entryPoint, appId, appKey},
		clientId, clientSecret,
	}, nil
}
