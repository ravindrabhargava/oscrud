package main

import "github.com/si3nloong/sqlike/sql/expr"

// User :
type User struct {
	Key  int64  `json:"id"`
	Name string `json:"name"`
}

// ToQuery :
func (user User) ToQuery(pk string) interface{} {
	var query interface{}

	if pk != "" {
		query = expr.Equal("Key", pk)
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
