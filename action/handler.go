package action

// Handler :
type Handler func(EndpointContext, ServiceContext)

// RequestHandler :
type RequestHandler func(interface{}) Handler
