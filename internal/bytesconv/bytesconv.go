package bytesconv

import "unsafe"

// StringToBytes converts string to  byte slice without a memory allocation.
func StringToBytes(s string) []byte {
	p := unsafe.Pointer(&struct {
		string
		Cap int
	}{s, len(s)})
	return *(*[]byte)(p)
}

//BytesToString converts byte slice to  string without a memory allocation.
func BytesToString(bts []byte) string {
	return *(*string)(unsafe.Pointer(&bts))
}
