package oscrud

// Service :
type Service interface {
	Find(*Context) *Context
	Create(*Context) *Context
	Get(string, *Context) *Context
	Update(string, *Context) *Context
	Patch(string, *Context) *Context
	Delete(string, *Context) *Context
}

// ServiceOptions :
type ServiceOptions struct {
	DisableFind   bool
	DisableCreate bool
	DisableGet    bool
	DisableUpdate bool
	DisablePatch  bool
	DisableDelete bool
}

// ServiceAction :
type ServiceAction string

// ServiceActions :
var (
	ServiceActionCreate ServiceAction = "CREATE"
	ServiceActionFind   ServiceAction = "FIND"
	ServiceActionGet    ServiceAction = "GET"
	ServiceActionUpdate ServiceAction = "UPDATE"
	ServiceActionPatch  ServiceAction = "PATCH"
	ServiceActionDelete ServiceAction = "DELETE"
)

// ServiceModel :
type ServiceModel interface {
	ToResult(*Context, ServiceAction) (interface{}, error)
	ToQuery(*Context, ServiceAction) (interface{}, error)
	ToCreate(*Context) error
	ToDelete(*Context) error
	ToPatch(*Context, ServiceModel) error
	ToUpdate(*Context, ServiceModel) error
}

// transforms id endpoint to proper oscrud handler
func serviceHandler(handler func(string, *Context) *Context) Handler {
	return func(ctx *Context) *Context {
		var i struct {
			ID string `param:"id"`
		}

		ctx.Bind(&i)
		return handler(i.ID, ctx)
	}
}
