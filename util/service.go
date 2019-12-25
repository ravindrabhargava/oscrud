package util

import "strings"

// TransformPath :
func TransformPath(base, action string) string {
	switch strings.ToLower(action) {
	case "find":
		return base
	case "create":
		return base
	case "get":
		return base + "/:id"
	case "update":
		return base + "/:id"
	case "patch":
		return base + "/:id"
	case "remove":
		return base + "/:id"
	}
	return ""
}

// GetActionByRoute :
func GetActionByRoute(route string) string {
	if strings.HasPrefix(route, "POST") {
		return "create"
	}

	if strings.HasPrefix(route, "PUT") {
		return "update"
	}

	if strings.HasPrefix(route, "PATCH") {
		return "patch"
	}

	if strings.HasPrefix(route, "DELETE") {
		return "remove"
	}

	if strings.HasSuffix(route, "/:id") {
		return "get"
	}
	return "find"
}

// GetMethodByAction :
func GetMethodByAction(action string) string {
	switch strings.ToLower(action) {
	case "find", "get":
		return "get"
	case "create":
		return "post"
	case "update":
		return "put"
	case "patch":
		return "patch"
	case "remove":
		return "delete"
	}
	return ""
}
