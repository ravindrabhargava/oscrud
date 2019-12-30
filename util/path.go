package util

import (
	"strings"
)

// RadixPath :
func RadixPath(method, path string) string {
	return strings.ToLower(method) + FixPath(path)
}

// FixPath :
func FixPath(path string) string {
	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}

	if strings.HasSuffix(path, "/") {
		path = strings.TrimSuffix(path, "/")
	}

	return path
}
