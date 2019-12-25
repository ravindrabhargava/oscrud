package util

import "strings"

// GetMethodPathFromRoute :
func GetMethodPathFromRoute(route string) (string, string) {
	setting := strings.Split(route, " ")
	return setting[0], setting[1]
}
