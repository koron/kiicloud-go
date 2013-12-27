package kii

// Administrator packs administration ID and secret of the application.
type Administrator struct {

	// Client ID of the application.
	Id     string

	// Client secret of the application.
	Secret string
}

// Application packs application's information: site ID, key and
// administration.
type Application struct {

	// Site of the application.  Be determined when it created.
	Site  Site

	// ID of the application.
	Id    string

	// Key of the application.
	Key   string

	// Administrator information.  Omittable.
	Admin *Administrator
}

// Create an application scope.
func (app *Application) NewAppScope() (Scope, err error) {
	// TODO:
	return nil, nil
}

// Create an user scope.
func (app *Application) NewUserScope(name, pass string) (Scope, err error) {
	// TODO:
	return nil, nil
}

// Create an administration scope.
func (app *Application) NewAdminScope() (Scope, err error) {
	// TODO:
	return nil, nil
}
