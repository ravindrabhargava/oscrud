package oscrud

// Service :
type Service interface {
	Find(Context) Context
	Get(Context) Context
	Create(Context) Context
	Update(Context) Context
	Patch(Context) Context
	Delete(Context) Context
}

// ServiceOptions :
type ServiceOptions struct {
	DisableFind   bool
	DisableGet    bool
	DisableCreate bool
	DisableUpdate bool
	DisablePatch  bool
	DisableDelete bool
}

// ServiceModel :
type ServiceModel interface {
	ToCreate() (interface{}, error)
	ToResult() (interface{}, error)

	ToQuery() (interface{}, error)
	ToPatch(ServiceModel) (interface{}, error)
	ToUpdate(ServiceModel) (interface{}, error)
	ToDelete() (interface{}, error)
}

// Query :
type Query struct {
	Cursor string `query:"$cursor"`
	Offset int    `query:"$offset"`
	Page   int    `query:"$page"`
	Limit  int    `query:"$limit"`
	Order  string `query:"$order"`
	Select string `query:"$select"`
	Query  map[string]interface{}
}

// QueryOne :
type QueryOne struct {
	Query  map[string]interface{}
	Select string `query:"$select"`
}
