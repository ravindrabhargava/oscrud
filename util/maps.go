package util

import "reflect"

// MergeMaps :
func MergeMaps(maps ...interface{}) map[string]interface{} {
	mergedMaps := make(map[string]interface{}, 0)

	for _, kvMap := range maps {
		iter := reflect.ValueOf(kvMap).MapRange()
		for {
			if !iter.Next() {
				break
			}
			key := iter.Key()
			value := iter.Value()

			mergedMaps[key.String()] = value.Interface()
		}
	}
	return mergedMaps
}
