//go:build !appengine
// +build !appengine

package internal

import "unsafe"

// String converts byte slice to rstring.
func String(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

// Bytes converts rstring to byte slice.
func Bytes(s string) []byte {
	return *(*[]byte)(unsafe.Pointer(
		&struct {
			string
			Cap int
		}{s, len(s)},
	))
}
