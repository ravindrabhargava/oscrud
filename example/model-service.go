package main

// User :
type User struct {
	Key  int64  `json:"id"`
	Name string `json:"name"`
}

// ToQuery :
func (user User) ToQuery() map[string]interface{} {
	query := make(map[string]interface{})

	if user.Key != 0 {
		query["Key"] = user.Key
	}

	if user.Name != "" {
		query["Name"] = user.Name
	}

	return query
}

// ToUpdate :
func (user User) ToUpdate() interface{} {
	return user
}

// ToCreate :
func (user User) ToCreate() interface{} {
	return user
}

// ToResult :
func (user User) ToResult() interface{} {
	return user
}
