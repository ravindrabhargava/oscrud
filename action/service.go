package action

// ServiceHandler :
type ServiceHandler func(ServiceContext) error

// ServiceContext :
type ServiceContext interface {
	GetTransport() string
	GetType() string
	GetID() string
	GetBody() map[string]interface{}
	GetQuery() map[string]interface{}
	GetContext() interface{}
	Bind(i interface{}) error

	String(status int, text string) error
	HTML(status int, html string) error
	JSON(status int, i interface{}) error
	XML(status int, i interface{}) error
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
