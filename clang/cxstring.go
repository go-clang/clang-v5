package clang

// #include "go-clang.h"
import "C"

// cxstring a character string.
//
// The C.CXString type is used to return strings from the interface when the ownership of that string might different from one call to the next.
// Use String() to retrieve the string data and, once finished with the string data, call Dispose() to free the string.
type cxstring struct {
	c C.CXString
}

// Retrieve the character data associated with the given string.
func (c cxstring) String() string {
	cstr := C.clang_getCString(c.c)
	return C.GoString(cstr)
}

// Free the given string.
func (c cxstring) Dispose() {
	C.clang_disposeString(c.c)
}
