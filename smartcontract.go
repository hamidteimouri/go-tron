package gotron

// RemoveMethodSignature the method signature (first 4 bytes) from the input data
func RemoveMethodSignature(data string) string {
	return data[8:]
}
