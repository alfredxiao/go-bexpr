package bexpr

import (
	"fmt"
	"reflect"
)

func primitiveLessThanFn(kind reflect.Kind) func(first interface{}, second reflect.Value) bool {
	println(fmt.Sprintf("kind: %v", kind))
	switch kind {
	case reflect.Bool:
		return doEqualBool
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return doLessThanInt64
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return doLessThanUint64
	case reflect.Float32:
		return doEqualFloat32
	case reflect.Float64:
		return doEqualFloat64
	case reflect.String:
		println("yada less than string")
		return doEqualString
	default:
		return nil
	}
}

func doLessThanInt64(first interface{}, second reflect.Value) bool {
	return second.Int() < first.(int64)
}

func doLessThanUint64(first interface{}, second reflect.Value) bool {
	return second.Uint() < first.(uint64)
}
