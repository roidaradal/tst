package tst

import "reflect"

// isNil checks if given item is nil
func isNil(x any) bool {
	if x == nil {
		return true
	}
	switch reflect.TypeOf(x).Kind() {
	case reflect.Pointer, reflect.Map, reflect.Slice, reflect.Chan, reflect.Func, reflect.Interface:
		return reflect.ValueOf(x).IsNil()
	default:
		return false
	}
}
