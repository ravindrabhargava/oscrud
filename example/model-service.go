package main

// User :
type User struct {
	Key  int64  `json:"id"`
	Name string `json:"name"`
}

// ToQuery :
func (user User) ToQuery(pk string) map[string]interface{} {
	query := make(map[string]interface{})

	if pk != "" {
		query["Key"] = pk
	}

	if user.Name != "" {
		query["Name"] = user.Name
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
