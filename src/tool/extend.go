package tool

func MapCopy[K comparable, V any](from map[K]V) map[K]V {
	rtn := make(map[K]V, len(from))
	for k, v := range from {
		rtn[k] = v
	}
	return rtn
}
func MapAppend[K comparable, V any](from map[K]V, to map[K]V) map[K]V {
	if to == nil {
		return from
	}
	for k, v := range from {
		to[k] = v
	}
	return to
}
