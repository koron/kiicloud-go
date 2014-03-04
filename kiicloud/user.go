package kiicloud

type UserClient struct {
	// All inherit from Client.
	Client
}

func NewUserClient(entryPoint, appId, appKey, userName, password string) (*UserClient, error) {
	// TODO: implement me.
	return nil, nil
}
