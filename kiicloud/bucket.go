package kiicloud

// Bucket packs info to access the app.
type Bucket struct {
	// All inherit from Client.
	*Client

	// PathRoot holds path of bucket root.
	PathRoot string

	// Parent object.
	Parent interface{}
}
