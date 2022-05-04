package assertions

import (
	"reflect"
)

func isNil(object interface{}) bool {
	if object == nil {
		return true
	}
	value := reflect.ValueOf(object)
	kind := value.Kind()
	if kind >= reflect.Chan && kind <= reflect.Slice && value.IsNil() {
		return true
	}
	return false
}

func isNotNil(object interface{}) bool {
	return !isNil(object)
}

func areSame(objectA, objectB interface{}) bool {
	if areNotEqual(objectA, objectB) {
		return false
	}
	if objectA != objectB {
		return false
	}
	return true
}

func areNotSame(objectA, objectB interface{}) bool {
	return !areSame(objectA, objectB)
}

func areEqual(objectA, objectB interface{}) bool {
	if isNil(objectA) && isNil(objectB) {
		return true
	}
	if isNil(objectA) || isNil(objectB) {
		return false
	}
	if reflect.DeepEqual(objectA, objectB) {
		return true
	}
	aValue := reflect.ValueOf(objectA)
	bValue := reflect.ValueOf(objectB)
	return aValue == bValue
}

func areNotEqual(objectA, objectB interface{}) bool {
	return !areEqual(objectA, objectB)
}
