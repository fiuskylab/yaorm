package helpers

// InSlice checks if value exists in a given Slice of T.
func InSlice[T comparable](value T, slice []T) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}
