package packages

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func injectValueToStruct(field interface{}, value interface{}) (ok bool) {
	t1 := reflect.TypeOf(field)
	t2 := reflect.TypeOf(value)
	if t1.Kind() != reflect.Ptr {
		return false
	}
	if t2.Kind() == reflect.Ptr {
		t2 = t2.Elem()
	}

	if t1.Elem().Kind() != t2.Kind() {
		return false
	}

	reflect.ValueOf(field).Elem().Set(reflect.ValueOf(value))

	return true
}

func TestInjectValueToStruct(t *testing.T) {
	m := struct{ Name string }{
		Name: "inits",
	}

	assert.False(t, injectValueToStruct(m.Name, nil), "not pointer field")
	assert.True(t, injectValueToStruct(&m.Name, "str2"), "into string without pointer value")
	assert.False(t, injectValueToStruct(&m.Name, 1234), "not same type")
}
