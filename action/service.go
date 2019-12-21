package action

// ServiceContext :
type ServiceContext interface {
	// GetType() string
	// GetID() interface{}
	// GetBody() []byte
	// ParseBody(body interface{}) error
	// GetQuery() map[string]interface{}
	// ParseQuery(query interface{}) error
	// Get(assign interface{}) error
}

// Service :
type Service interface {
	Find(ServiceContext, EndpointContext) error
	Get(ServiceContext, EndpointContext) error
	Create(ServiceContext, EndpointContext) error
	Update(ServiceContext, EndpointContext) error
	Patch(ServiceContext, EndpointContext) error
	Remove(ServiceContext, EndpointContext) error
}
