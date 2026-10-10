package reflt

import (
	"reflect"
)

func Walk(x interface{}, fn func(input string)) {
	val := getValue(x)

	walkvalue := func(value reflect.Value) {
		Walk(value.Interface(), fn)
	}

	switch val.Kind() {
	case reflect.String:
		fn(val.String())
	case reflect.Struct:
		for i := 0; i < val.NumField(); i++ {
			walkvalue(val.Field(i))
		}
	case reflect.Slice, reflect.Array:
		for i := 0; i < val.Len(); i++ {
			walkvalue(val.Index(i))
		}
	case reflect.Map:
		for _, key := range val.MapKeys() {
			walkvalue(val.MapIndex(key))
		}
	case reflect.Chan:
		for {
			if v, ok := val.Recv(); ok {
				walkvalue(v)
			} else {
				break
			}
		}
	case reflect.Func:
		vals := val.Call(nil)
		for _, res := range vals {
			walkvalue(res)
		}
	}
}

func getValue(x interface{}) reflect.Value {
	val := reflect.ValueOf(x)
	if reflect.Pointer == val.Kind() {
		val = val.Elem()
	}
	return val
}
