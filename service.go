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
}

// QueryOne :
type QueryOne struct {
	Select string `query:"$select"`
}

// ServiceMeta :
type ServiceMeta struct {
	Cursor     string `json:"cursor,omitempty" xml:"cursor"`
	Limit      string `json:"perPage,omitempty" xml:"perPage"`
	Page       string `json:"currentPage,omitempty" xml:"currentPage"`
	Total      string `json:"total,omitempty" xml:"total"`
	TotalPages string `json:"totalPages,omitempty" xml:"totalPages"`
}

// ServiceResult :
type ServiceResult struct {
	Meta   ServiceMeta `json:"meta" xml:"meta"`
	Result interface{} `json:"result" xml:"result"`
}
