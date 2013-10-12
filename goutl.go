package goutl

// Convert all byte slices in a slice of bytes to string slice
func BytesToStrings(b [][]byte) []string {
	s := make([]string, len(b))

	for i, val := range b {
		s[i] = string(val)
	}

	return s
}
