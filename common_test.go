// Like for github.com/stretchr/testify/assert/assertions.go:58 v1.7.0

package crypto_test

import (
	"bytes"
	"reflect"
)

type TestingAny interface {
	Errorf(format string, args ...interface{})
}

func noError(t TestingAny, err error) {
	if err != nil {
		t.Errorf("Received unexpected error:\n%+v", err)
	}
}

func equalLength(t TestingAny, a, b []byte) {
	al, bl := len(a), len(b)
	if al != bl {
		t.Errorf("Not equal length: \n"+
			"\texpected: %s\n"+
			"\tactual  : %s", al, bl)
	}
}

func equal(t TestingAny, a, b interface{}) {
	if !areEqual(a, b) {
		t.Errorf("Not equal: \n"+
			"\texpected: %s\n"+
			"\tactual  : %s", a, b)
	}
}

func areEqual(expected, actual interface{}) bool {
	if expected == actual {
		return true
	}
	if expected == nil || actual == nil {
		return false
	}

	exp, ok := expected.([]byte)
	if !ok {
		return reflect.DeepEqual(expected, actual)
	}

	act, ok := actual.([]byte)
	if !ok {
		return false
	}
	if exp == nil || act == nil {
		return exp == nil && act == nil
	}
	return bytes.Equal(exp, act)
}
