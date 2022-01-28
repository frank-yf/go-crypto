// Like for github.com/stretchr/testify/assert/assertions.go:58 v1.7.0

package crypto_test

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

func equalBytes(t TestingAny, a, b []byte) {
	equalString(t, string(a), string(b))
}

func equalString(t TestingAny, a, b string) {
	if a != b {
		t.Errorf("Not equal: \n"+
			"\texpected: %s\n"+
			"\tactual  : %s", a, b)
	}
}
