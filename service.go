package oscrud

// ServiceContext :
type ServiceContext struct {
	Service string
	Action  string
	ID      string
	Body    map[string]interface{}
	Query   map[string]interface{}
	Result  *ServiceResult
}

// ServiceResult :
type ServiceResult struct {
	Status      int
	ContentType string
	Result      interface{}
}

// GetContext :
func (c ServiceContext) GetContext() interface{} {
	return nil
}

// GetTransport :
func (c ServiceContext) GetTransport() string {
	return "INTERNAL"
}

// GetType :
func (c ServiceContext) GetType() string {
	return c.Action
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

// String :
func (c ServiceContext) String(status int, text string) error {
	c.Result = &ServiceResult{
		Status:      status,
		Result:      text,
		ContentType: "text/plain",
	}
	return nil
}

// HTML :
func (c ServiceContext) HTML(status int, html string) error {
	c.Result = &ServiceResult{
		Status:      status,
		Result:      html,
		ContentType: "text/html",
	}
	return nil
}

// JSON :
func (c ServiceContext) JSON(status int, i interface{}) error {
	c.Result = &ServiceResult{
		Status:      status,
		Result:      i,
		ContentType: "application/json",
	}
	return nil
}

// XML :
func (c ServiceContext) XML(status int, i interface{}) error {
	c.Result = &ServiceResult{
		Status:      status,
		Result:      i,
		ContentType: "application/xml",
	}
	return nil
}
