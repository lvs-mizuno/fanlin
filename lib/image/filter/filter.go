package filter

import (
	"crypto/sha512"
	"fmt"
	"net/http"
)

var filters = map[string]func([]byte, *http.Request) []byte{}

func andArray(r []byte, b [64]byte) []byte {
	for i := range r {
		r[i] &= b[i]
	}
	return r
}

func convertHash(options []string) string {
	result := make([]byte, 64)
	for i := range result {
		result[i] = 0xff
	}
	for _, o := range options {
		b := sha512.Sum512([]byte(o))
		result = andArray(result, b)
	}
	return fmt.Sprintf("%x", result)
}

func RegistFilter(options []string, f func([]byte, *http.Request) []byte) {
	filters[convertHash(options)] = f
}
