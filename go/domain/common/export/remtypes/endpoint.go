package remtypes

// TODO(teawithsand): export these endpoints to separate package

type Endpoint struct {
	Path    string
	Handler interface{}
}

func (ep *Endpoint) GetPattern() string {
	return ep.Path
}

var UserConfirmRegistrationEndpoint = Endpoint{
	Path: "/register/confirm",
}
