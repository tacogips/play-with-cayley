package main

import (
	"crypto/sha1"
	"fmt"
	"reflect"
)

var h = sha1.New()

func main() {
	fmt.Println("vim-go")

	p := make([]byte, 20)
	//p := make([]byte, 1)
	///r := HashTo("12346", p)
	//fmt.Printf("%#v\n", p)
	fmt.Printf("%s\n", reflect.TypeOf(p))
	//fmt.Printf("%#v\n", r)

}

// quad/value.go
// HashTo calculates a hash of value v, storing it in a slice p.
func HashTo(s string, p []byte) []byte {
	h.Reset()
	if len(s) > 0 {
		h.Write([]byte(s))
	}
	return h.Sum(p[:0])
	//h.Sum(p)
}
