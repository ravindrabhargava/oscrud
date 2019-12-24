package oscrud

// ServiceContext :
type ServiceContext struct {
	Action string
	Path   string
	Type   string
	ID     string
	Body   map[string]interface{}
	Query  map[string]interface{}
}

// GetTransport :
func (c ServiceContext) GetTransport() string {
	return "INTERNAL"
}

// GetType :
func (c ServiceContext) GetType() string {
	return c.Type
}

// GetID :
func (c ServiceContext) GetID() string {
	return c.ID
}

// GetBody :
func (c ServiceContext) GetBody() map[string]interface{} {
	return c.Body
}

// GetQuery :
func (c ServiceContext) GetQuery() map[string]interface{} {
	return c.Query
}

// Bind :
func (c ServiceContext) Bind(i interface{}) error {
	return BindService(c.ID, c.Body, c.Query, i)
}
