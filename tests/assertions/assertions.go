package assertions

import (
	"testing"
)

const (
	defaultEqualsMessage  = "\nExpected: '%#v' (type %T)\nBut got:  '%#v' (type %T)"
	defaultNotSameMessage = "the objects\n(%p) %T %v and\n(%p) %T %v\nare not same object."
)

func AssertEquals(t *testing.T, expected, obtained interface{}) {
	t.Helper()
	if areNotEqual(expected, obtained) {
		t.Errorf(defaultEqualsMessage, expected, expected, obtained, obtained)
	}
}

func AssertSame(t *testing.T, objectA, objectB interface{}) {
	t.Helper()
	if areNotSame(objectA, objectB) {
		t.Errorf(
			defaultNotSameMessage,
			objectA, objectA, objectA, objectB, objectB, objectB,
		)
	}
}

func AssertUnexpectedError(t *testing.T, err error) {
	t.Helper()
	if isNotNil(err) {
		t.Errorf("An unexpected error occurred: %s", err)
	}
}

func AssertError(t *testing.T, expected, obtained error) {
	t.Helper()
	if isNil(obtained) {
		t.Fatal("An expected error didn't occur.")
	}
	if expected != obtained {
		t.Errorf("Expected: '%s', but was: '%s'", expected, obtained)
	}
}
