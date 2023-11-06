package utils

// this function returns a pointer to the value it accepted
// useful for easily populating aws input structs, as they present *string fields
func Ptr[T any](v T) *T {
	return &v
}
