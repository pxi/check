// Package is provides helpers for testing.
package is

import (
	"reflect"
	"testing"
)

// Ok fails the test if the given error is not nil.
func Ok(tb testing.TB, err error) {
	tb.Helper()
	if err != nil {
		tb.Fatal(err)
	}
}

// Equal fails the test if got is not equal to want.
func Equal(tb testing.TB, got, want interface{}) {
	tb.Helper()
	if !reflect.DeepEqual(got, want) {
		tb.Fatalf("\n got: %#v\nwant: %#v", got, want)
	}
}

// Content fails the test if condition is false.
func Content(tb testing.TB, condition bool, format string, args ...interface{}) {
	tb.Helper()
	if !condition {
		tb.Fatalf(format, args...)
	}
}
