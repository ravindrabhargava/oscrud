package oscrud

// Handler :
type Handler func(Context) Context

// Route :
type Route struct {
	Method  string
	Route   string
	Path    string
	Handler []Handler
}
