package kiicloud

// Bucket packs info to access the app.
type Bucket struct {
	// Parent object.
	Parent interface{}

	// Client 
	Client *Client

	PathRoot string
}
