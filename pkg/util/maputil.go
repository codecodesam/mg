package util

// CopyMap
// Low Copy
func CopyMap(dst, src map[string]any) {

	if src == nil || dst == nil {
		return
	}

	for k, v := range src {
		dst[k] = v
	}

}
