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
