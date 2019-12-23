package action

// ServiceHandler :
type ServiceHandler func(ServiceContext) error

// ServiceContext :
type ServiceContext interface {
	GetTransport() string
	GetType() string
	GetID() string
	GetBody() string
	GetQuery() map[string]interface{}
	Bind(i interface{}) error
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
