package oscrud

// Service :
type Service interface {
	Find(Context) Context
	Create(Context) Context
	Get(string, Context) Context
	Update(string, Context) Context
	Patch(string, Context) Context
	Delete(string, Context) Context
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

type serviceAction string

// ServiceActions :
var (
	ServiceActionCreate serviceAction = "CREATE"
	ServiceActionFind   serviceAction = "FIND"
	ServiceActionGet    serviceAction = "GET"
	ServiceActionUpdate serviceAction = "UPDATE"
	ServiceActionPatch  serviceAction = "PATCH"
	ServiceActionDelete serviceAction = "DELETE"
)

// ServiceModel :
type ServiceModel interface {
	ToResult(serviceAction) (interface{}, error)
	ToQuery(serviceAction) (interface{}, error)
	ToCreate() error
	ToDelete() error
	ToPatch(ServiceModel) error
	ToUpdate(ServiceModel) error
}

// transforms $id endpoint to proper oscrud handler
func serviceHandler(handler func(string, Context) Context) Handler {
	return func(ctx Context) Context {
		var i struct {
			ID string `param:"id"`
		}

		ctx.Bind(&i)
		return handler(i.ID, ctx)
	}
}
