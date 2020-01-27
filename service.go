package oscrud

// Service :
type Service interface {
	Find(Context) Context
	Get(Context) Context
	Create(Context) Context
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
	Pk     string `param:"$id"`
	Select string `query:"$select"`
}

// ServiceModel :
type ServiceModel interface {
	ToCreate() interface{}
	ToUpdate() interface{}
	ToResult() interface{}
	ToQuery(string) interface{}
}
