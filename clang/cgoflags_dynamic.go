//go:build !static
// +build !static

package clang

// #cgo LDFLAGS: -lclang
import "C"
