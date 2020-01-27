package util

// MergeMaps :
func MergeMaps(maps ...interface{}) map[string]interface{} {
	mergedMaps := make(map[string]interface{}, 0)

	for _, kvMap := range maps {
		for k, v := range kvMap.(map[string]interface{}) {
			mergedMaps[k] = v
		}
	}

	return mergedMaps
}
