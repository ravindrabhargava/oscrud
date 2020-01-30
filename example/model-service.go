package main

import "github.com/si3nloong/sqlike/sql/expr"

// User :
type User struct {
	Key  int64  `json:"id" qm:"$id"`
	Name string `json:"name"`
}

// ToQuery :
func (user User) ToQuery() interface{} {
	var query interface{}

	if user.Key != 0 {
		query = expr.Equal("Key", user.Key)
	}

	return query
}

// ToUpdate :
func (user *User) ToUpdate() interface{} {
	return user
}

// ToCreate :
func (user *User) ToCreate() interface{} {
	user.Name += "-NEW"
	return user
}

// ToResult :
func (user *User) ToResult() interface{} {
	return user
}
