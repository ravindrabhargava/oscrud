package main

// Service :
type Service interface {
	Find(Context) Context
	Get(Context) Context
	Create(Context) Context
	Update(Context) Context
	Patch(Context) Context
	Delete(Context) Context
}

// DataModel :
type DataModel interface {
	ToCreate() interface{}
	ToUpdate() interface{}
	ToResult() interface{}
	ToQuery() interface{}
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
