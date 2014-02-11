package kiicloud

// Site represent region of server for the app.
type Site int

const (
	// US is United States region.
	US Site = iota
	// JP is Japan region.
	JP Site = iota
	// CN is China region.
	CN Site = iota
)

// EndPoint return end point for the site.
func (s Site) EndPoint() (endPoint string) {
	switch s {
	case US:
		endPoint = "https://api.kii.com/api"
	case JP:
		endPoint = "https://api-jp.kii.com/api"
	case CN:
		endPoint = "https://api-cn2.kii.com/api"
	}
	return
}

// expandSite expand a Site as an end point.
func expandSite(src string) (dst string, err error) {
	switch src {
	case "US":
		dst = US.EndPoint()
	case "JP":
		dst = JP.EndPoint()
	case "CN":
		dst = CN.EndPoint()
	default:
		dst = src
	}
	return
}
