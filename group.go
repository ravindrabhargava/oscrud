package oscrud

import (
	"path"
)

// Group :
type Group struct {
	path   string
	opts   []Options
	server *Oscrud
}

// Group :
func (server *Oscrud) Group(path string, opts ...Options) Group {
	return Group{path, opts, server}
}

// RegisterEndpoint :
func (g Group) RegisterEndpoint(method, endpoint string, handler Handler, opts ...Options) Group {
	cleanPath := path.Clean(g.path + endpoint)
	opts = append(g.opts)
	g.server = g.server.RegisterEndpoint(method, cleanPath, handler, opts...)
	return g
}

// RegisterService :
func (g Group) RegisterService(basePath string, service Service, serviceOptions *ServiceOptions, opts ...Options) Group {
	cleanPath := path.Clean(g.path + basePath)
	opts = append(g.opts)
	g.server = g.server.RegisterService(cleanPath, service, serviceOptions, opts...)
	return g
}
