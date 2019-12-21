package action

// ServiceHandler :
type ServiceHandler func(ServiceContext) error

// ServiceContext :
type ServiceContext interface {
	GetType() string
	GetID() string
	GetBody() string
	GetTransport() string
	// ParseID(assign interface{}) error
	// ParseBody(body interface{}) error
	// GetQuery() map[string]interface{}
	// ParseQuery(query interface{}) error
	// Get(assign interface{}) error
}

// Service :
type Service interface {
	Find(ServiceContext) error
	Get(ServiceContext) error
	Create(ServiceContext) error
	Update(ServiceContext) error
	Patch(ServiceContext) error
	Remove(ServiceContext) error
}
